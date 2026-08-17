package cmd

import (
	"net/http"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// captureRequest is a runtime.ClientRequest that records header params and
// does nothing else. Only SetHeaderParam/GetHeaderParams are exercised; the
// rest exist because the interface is wide.
type captureRequest struct {
	headers http.Header
}

func newCaptureRequest() *captureRequest {
	return &captureRequest{headers: http.Header{}}
}

func (c *captureRequest) SetHeaderParam(name string, values ...string) error {
	for _, v := range values {
		c.headers.Add(name, v)
	}
	return nil
}

func (c *captureRequest) GetHeaderParams() http.Header                          { return c.headers }
func (c *captureRequest) SetQueryParam(string, ...string) error                 { return nil }
func (c *captureRequest) SetFormParam(string, ...string) error                  { return nil }
func (c *captureRequest) SetPathParam(string, string) error                     { return nil }
func (c *captureRequest) GetQueryParams() url.Values                            { return nil }
func (c *captureRequest) SetFileParam(string, ...runtime.NamedReadCloser) error { return nil }
func (c *captureRequest) SetBodyParam(interface{}) error                        { return nil }
func (c *captureRequest) SetTimeout(time.Duration) error                        { return nil }
func (c *captureRequest) GetMethod() string                                     { return http.MethodGet }
func (c *captureRequest) GetPath() string                                       { return "/v1/projects" }
func (c *captureRequest) GetBody() []byte                                       { return nil }
func (c *captureRequest) GetBodyParam() interface{}                             { return nil }
func (c *captureRequest) GetFileParam() map[string][]runtime.NamedReadCloser    { return nil }

var _ runtime.ClientRequest = (*captureRequest)(nil)

const testAPIKey = "gtk_0123456789abcdef_c2VjcmV0dmFsdWU"

// TestNewAuthInformerHeaders is the heart of the api key work (#21).
//
// An api key MUST travel in X-Api-Key and MUST NOT travel in Authorization.
// Istio's RequestAuthentication parses any Authorization bearer token as a JWT
// and rejects what will not parse, so a gtk_ key in that header dies at the
// ingress with
//
//	401 Jwt is not in the form of Header.Payload.Signature with two dots and 3 sections
//
// before reaching any service (grandturks#1175). The server accepts both -
// common/apikey promotes X-Api-Key into the bearer credential before authn -
// so Authorization works in-cluster and fails through a gateway. That split is
// exactly what this test exists to prevent regressing.
//
// The "both" row is the other half, and it is not obvious: the server's
// PromoteToAuthorization gives Authorization PRECEDENCE when both headers are
// present. Sending a saved JWT alongside an explicitly-supplied key would mean
// the key is silently ignored and the request runs as the human - the opposite
// of what someone passing --api_key asked for. So the informer sends one
// credential, never two.
func TestNewAuthInformerHeaders(t *testing.T) {
	for _, tc := range []struct {
		name              string
		apiKey            string
		authToken         string
		wantAPIKeyHeader  string
		wantAuthorization string
	}{
		{
			name:             "api key only",
			apiKey:           testAPIKey,
			wantAPIKeyHeader: testAPIKey,
		},
		{
			name:              "auth token only",
			authToken:         "ya29.jwt-shaped-token",
			wantAuthorization: "Bearer ya29.jwt-shaped-token",
		},
		{
			name:             "both - the key wins and no Authorization is sent",
			apiKey:           testAPIKey,
			authToken:        "ya29.jwt-shaped-token",
			wantAPIKeyHeader: testAPIKey,
		},
		{
			name: "neither",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			req := newCaptureRequest()

			if err := newAuthInformer(tc.apiKey, tc.authToken).AuthenticateRequest(req, strfmt.Default); err != nil {
				t.Fatalf("AuthenticateRequest: %v", err)
			}

			if got := req.headers.Get(apiKeyHeader); got != tc.wantAPIKeyHeader {
				t.Errorf("%s = %q, want %q", apiKeyHeader, got, tc.wantAPIKeyHeader)
			}
			if got := req.headers.Get("Authorization"); got != tc.wantAuthorization {
				t.Errorf("Authorization = %q, want %q", got, tc.wantAuthorization)
			}
		})
	}
}

// TestNewAuthInformerNeverPutsAnApiKeyInAuthorization states the rule above as
// its own assertion rather than leaving it implicit in a table row: whatever
// else changes about credential precedence, the gtk_ token must never appear
// in the header the ingress inspects.
func TestNewAuthInformerNeverPutsAnApiKeyInAuthorization(t *testing.T) {
	req := newCaptureRequest()

	if err := newAuthInformer(testAPIKey, "some-jwt").AuthenticateRequest(req, strfmt.Default); err != nil {
		t.Fatalf("AuthenticateRequest: %v", err)
	}

	for _, value := range req.headers.Values("Authorization") {
		if strings.Contains(value, "gtk_") {
			t.Errorf("Authorization = %q carries an api key - it will be rejected as a malformed JWT at the ingress", value)
		}
	}
}

// TestNewAuthInformerRejectsAMalformedApiKey: a mistyped key must fail with a
// message that says what the CLI expected, before a request goes out. The
// alternative is a 401 from the server that looks identical to an expired one.
//
// The check lives in the informer rather than in flag parsing because the
// commands construct their client through mustNewRunCmd, which panics on
// error; returning it here surfaces as an ordinary command error instead.
func TestNewAuthInformerRejectsAMalformedApiKey(t *testing.T) {
	req := newCaptureRequest()

	err := newAuthInformer("not-a-key", "").AuthenticateRequest(req, strfmt.Default)
	if err == nil {
		t.Fatal("a malformed api key produced no error")
	}
	if !strings.Contains(err.Error(), "gtk_") {
		t.Errorf("error %q does not tell the user the expected shape", err)
	}
	if len(req.headers) != 0 {
		t.Errorf("headers = %v, want none set when the key is rejected", req.headers)
	}
}
