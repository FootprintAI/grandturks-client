package transport_test

// grandturks#1092: an expired token made every CLI command fail with
//
//	Error: &{0 [] } (*models.RPCStatus) is not supported by the TextConsumer,
//	       can be resolved by supporting TextUnmarshaler interface
//
// rather than with "Token Expired. Require Login first." - a message the code
// has had all along and could not reach.
//
// # THE MECHANISM
//
// The generated readers are all shaped like this:
//
//	default:
//	    result := NewXDefault(response.Code())
//	    if err := result.readResponse(response, consumer, o.formats); err != nil {
//	        return nil, err            // <- the status code is DROPPED here
//	    }
//	    return nil, result
//
// So when the body cannot be deserialized into models.RPCStatus, what comes
// back is the CONSUMER's error - a bare error with no Code() - and
// openapierrors.Parse returns any non-openapiError verbatim. Every message in
// its switch is unreachable on exactly the responses that do not carry a
// well-formed RPCStatus body.
//
// That is not a 401-only problem, which is the half #1092 asked about and this
// file answers: the same is true of 403, 500 and anything else whose body an
// ingress, a proxy or a mux wrote. In the reported case the 401 never reached
// appkafeido at all - it was rejected upstream, by something that does not
// speak this API's error shape and never will.
//
// # WHY THE FIX IS A ROUND TRIPPER
//
// The status is present and correct in the HTTP response; it is lost one layer
// above, inside generated code this repo must not hand-edit. A RoundTripper
// sits BELOW that layer, sees the real *http.Response, and can hand the
// generated reader a body it can actually deserialize - after which the
// existing switch works, unchanged, for every status in it.
//
// Keying on the consumer's error string instead would be the other option, and
// is what production/ubm01/scripts/lifecycle_test.py currently does in
// FootprintAI/manifests. #1092 calls that a workaround for a reason: it breaks
// the moment go-openapi rewords the message, and it can only ever recognise
// the failure, never tell 401 from 500.

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"path/filepath"
	"testing"

	httptransport "github.com/go-openapi/runtime/client"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	kafeidoopenapi "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client"
	kafeidoclient "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
	openapierrors "github.com/footprintai/grandturks-client/v2/pkg/http/openapi/errors"
	openapitransport "github.com/footprintai/grandturks-client/v2/pkg/http/openapi/transport"
)

// listProjectsAgainst points a real generated client at a server that answers
// however the test says, and returns whatever the CLI would have printed.
//
// The whole path is exercised: the transport under test, go-openapi's runtime,
// the generated reader, and Parse. Asserting on any narrower slice would prove
// a layer works while the thing a user sees stayed broken - which is precisely
// how this defect survived: openapierrors' own unit tests pass today.
func listProjectsAgainst(t *testing.T, handler http.HandlerFunc) error {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	_, err := stubAgainst(t, srv.URL).KafeidoService.KafeidoServiceListProjects(
		kafeidoclient.NewKafeidoServiceListProjectsParams(), nil)
	require.Error(t, err, "the server answered an error status; the client reported success")

	// hasDetail=true, the CLI's --debug. The details half is where the
	// server's own words survive, and dropping them would trade one opaque
	// message for another.
	return openapierrors.Parse(err, true)
}

// TestExpiredTokenSaysTokenExpired is the defect, exactly as reported: a 401
// whose body is not RPCStatus JSON, because whatever rejected the request
// upstream does not speak this API's error shape.
func TestExpiredTokenSaysTokenExpired(t *testing.T) {
	err := listProjectsAgainst(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write([]byte("token is expired\n"))
	})

	assert.Contains(t, err.Error(), "Token Expired",
		"a user whose token aged out has to be told to log in again. Got: %v", err)
	assert.NotContains(t, err.Error(), "TextConsumer",
		"the go-openapi consumer error reached the user; it reads as a client/server contract "+
			"bug and sends them to debug the deployment (grandturks#1092)")
}

// The broader half of #1092's question - "worth checking whether this affects
// only text/plain error bodies, or every non-2xx whose body is not valid
// RPCStatus JSON". It is every one of them, so every message in Parse's switch
// is affected, not just the 401.
func TestEveryOpaqueErrorStatusReachesItsMessage(t *testing.T) {
	for _, tc := range []struct {
		name        string
		status      int
		contentType string
		body        string
		want        string
	}{
		{
			name:        "403 from a proxy, as html",
			status:      http.StatusForbidden,
			contentType: "text/html",
			body:        "<html><body>403 Forbidden</body></html>",
			want:        "Permission denied",
		},
		{
			name:        "500 with an empty body",
			status:      http.StatusInternalServerError,
			contentType: "",
			body:        "",
			want:        "Internal error",
		},
		{
			name:        "404 from a mux that answers in plain text",
			status:      http.StatusNotFound,
			contentType: "text/plain; charset=utf-8",
			body:        "404 page not found\n",
			want:        "Resource not found",
		},
		{
			name:        "502 whose body is JSON, but not an RPCStatus",
			status:      http.StatusBadGateway,
			contentType: "application/json",
			body:        `["upstream connect error"]`,
			want:        "unknown error",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := listProjectsAgainst(t, func(w http.ResponseWriter, _ *http.Request) {
				if tc.contentType != "" {
					w.Header().Set("Content-Type", tc.contentType)
				}
				w.WriteHeader(tc.status)
				_, _ = w.Write([]byte(tc.body))
			})

			assert.Contains(t, err.Error(), tc.want, "got: %v", err)
			assert.NotContains(t, err.Error(), "TextConsumer", "got: %v", err)
		})
	}
}

// The server's own words must survive. Replacing an opaque error with a
// friendly one that says nothing about what happened is a different failure,
// not a fix - an operator debugging a 500 needs the body, and --debug is how
// they ask for it.
func TestTheServersOwnWordsSurvive(t *testing.T) {
	err := listProjectsAgainst(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("upstream database is down"))
	})

	assert.Contains(t, err.Error(), "upstream database is down",
		"the body was discarded; the user is told something failed and nothing about what")
}

// A well-formed RPCStatus body must be left completely alone. This is the
// regression guard on the fix rather than on the defect: the normalizer must
// not touch a response the generated client can already read, or it would
// rewrite the API's real error messages with its own.
func TestWellFormedErrorBodiesAreUntouched(t *testing.T) {
	err := listProjectsAgainst(t, func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write([]byte(`{"code":7,"message":"project 42 is not visible to you","details":[]}`))
	})

	assert.Contains(t, err.Error(), "Permission denied")
	assert.Contains(t, err.Error(), "project 42 is not visible to you",
		"the API's own message was replaced by the transport's reconstruction")
}

// And a success must be untouched too - the normalizer only ever looks at
// non-2xx responses, and a bug here would corrupt every response the product
// makes rather than only the broken ones.
func TestSuccessfulResponsesAreUntouched(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"projectIds":["p-1","p-2"]}`))
	}))
	t.Cleanup(srv.Close)

	ok, err := stubAgainst(t, srv.URL).KafeidoService.KafeidoServiceListProjects(
		kafeidoclient.NewKafeidoServiceListProjectsParams(), nil)
	require.NoError(t, err)
	require.NotNil(t, ok.Payload)
	assert.Equal(t, []string{"p-1", "p-2"}, ok.Payload.ProjectIds)
}

// stubAgainst builds the client EXACTLY as app/kafeido/cli/cmd's newRunCmd
// does - same base path, same transport wrapper - so what these tests exercise
// is the CLI's own construction and not a convenient approximation of it.
func stubAgainst(t *testing.T, rawURL string) *kafeidoopenapi.GrandturkKafeidoAPIDocumentations {
	t.Helper()

	hostURL, err := url.Parse(rawURL)
	require.NoError(t, err)

	httpClient := &http.Client{Transport: openapitransport.New(nil)}
	return kafeidoopenapi.New(httptransport.NewWithClient(
		hostURL.Host,
		filepath.Join("api", kafeidoopenapi.DefaultBasePath),
		[]string{hostURL.Scheme},
		httpClient,
	), nil)
}
