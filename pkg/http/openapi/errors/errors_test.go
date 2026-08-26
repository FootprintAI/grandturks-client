package openapierrors

import (
	"errors"
	"strings"
	"testing"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// fakeOpenapiError satisfies the unexported openapiError interface, which is
// the only thing Parse acts on - a plain error is returned untouched.
type fakeOpenapiError struct {
	code int
	msg  string
}

func (f *fakeOpenapiError) Code() int     { return f.code }
func (f *fakeOpenapiError) Error() string { return f.msg }

// TestParseMessages pins the user-facing text for every status Parse handles.
//
// These strings are the whole output of this package, and they are what a
// kafeido CLI user reads and quotes into a bug report. Two of them shipped
// misspelled - "adminstrator", and "<detained>" where "<redacted>" was meant
// (grandturks-client#18). Nothing referenced them, so nothing caught it.
//
// The same two typos were fixed in FootprintAI/grandturks#1010, in a
// BYTE-IDENTICAL copy of this file at common/http/openapi/errors/errors.go.
// That fix did not reach any user: app/kafeido/cli renders through THIS
// package, so the strings a person actually sees came from here and were
// unchanged. That is the failure this test exists to stop repeating - not the
// typo, but a fix that lands in the copy nobody runs.
func TestParseMessages(t *testing.T) {
	for _, tc := range []struct {
		name string
		code int
		want string
	}{
		{"unauthorized", 401, "Token Expired. Require Login first."},
		{"bad request", 400, "Bad Parameter."},
		{"forbidden", 403, "Permission denied. You either don't have enough permission or haven't login first."},
		{"not found", 404, "Resource not found."},
		{"conflict", 409, "Status Conflicted."},
		{"internal", 500, "Internal error. Please contact your system administrator."},
		{"gateway", 504, "Bad Gateway. Please try later."},
		{"unmapped", 418, "unknown error"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := Parse(&fakeOpenapiError{code: tc.code, msg: "underlying"}, true)
			var got *Error
			if !errors.As(err, &got) {
				t.Fatalf("Parse returned %T, want *Error", err)
			}
			if got.PlainText != tc.want {
				t.Errorf("PlainText = %q, want %q", got.PlainText, tc.want)
			}
		})
	}
}

// TestParseSpelling is deliberately separate from the table above.
//
// The table would catch a regression by exact match, but it states the correct
// spelling only once - so a future edit that reintroduces the typo in BOTH the
// code and the table would pass. This asserts the property directly.
func TestParseSpelling(t *testing.T) {
	for _, code := range []int{400, 401, 403, 404, 409, 500, 504, 418} {
		err := Parse(&fakeOpenapiError{code: code, msg: "underlying"}, false)
		text := err.Error()
		for _, typo := range []string{"adminstrator", "detained"} {
			if strings.Contains(text, typo) {
				t.Errorf("status %d message contains %q: %s", code, typo, text)
			}
		}
	}
}

// TestParseRedactsDetail covers the hasDetail=false path, which is the one that
// surfaces the sentinel to the user - and it is not confined to 500. It appears
// in the "(details:...)" suffix of EVERY status, which is what made the old
// "<detained>" the more widely seen of the two typos.
func TestParseRedactsDetail(t *testing.T) {
	in := &fakeOpenapiError{code: 500, msg: "connection refused to seaweedfs-eval-s3:8333"}

	redacted := Parse(in, false)
	if strings.Contains(redacted.Error(), "seaweedfs") {
		t.Errorf("hasDetail=false leaked the underlying error: %s", redacted.Error())
	}
	if !strings.Contains(redacted.Error(), "<redacted>") {
		t.Errorf("hasDetail=false should show <redacted>, got: %s", redacted.Error())
	}

	kept := Parse(in, true)
	if !strings.Contains(kept.Error(), "seaweedfs") {
		t.Errorf("hasDetail=true should keep the underlying error, got: %s", kept.Error())
	}
}

// The two branches that do not produce an *Error at all.
func TestParsePassesThroughNonOpenapiErrors(t *testing.T) {
	plain := errors.New("not an openapi error")
	if got := Parse(plain, true); got != plain {
		t.Errorf("Parse(%v) = %v, want the error unchanged", plain, got)
	}
}

func TestParseTreatsOKAsNoError(t *testing.T) {
	if got := Parse(&fakeOpenapiError{code: 200, msg: "ok"}, true); got != nil {
		t.Errorf("Parse(200) = %v, want nil", got)
	}
}

// fakePayloadError also carries an RPCStatus, as every generated *XDefault
// does. The canned strings above are what this package knows when the server
// said nothing specific; a server that wrote an instruction for a human must
// have it reach the human.
type fakePayloadError struct {
	code    int
	payload *models.RPCStatus
}

func (f fakePayloadError) Code() int     { return f.code }
func (f fakePayloadError) Error() string { return "generated-client error text" }
func (f fakePayloadError) GetPayload() *models.RPCStatus {
	return f.payload
}

const (
	codeInvalidArgument    = 3
	codeFailedPrecondition = 9
)

// The case this exists for: a deployment refusing an out-of-date CLI. Note
// hasDetail=false - the message must survive WITHOUT --debug, or the person
// who needs it never sees it.
func TestFailedPreconditionInstructionReachesTheUser(t *testing.T) {
	const upgrade = "kafeido CLI 2.6.0 is too old for this deployment, which requires 2.7.0 or newer."

	err := Parse(fakePayloadError{
		code:    400,
		payload: &models.RPCStatus{Code: codeFailedPrecondition, Message: upgrade},
	}, false)

	if !strings.Contains(err.Error(), upgrade) {
		t.Errorf("the server's instruction did not reach the user, got %q", err.Error())
	}
	if strings.Contains(err.Error(), "Bad Parameter.") {
		t.Errorf("the canned string replaced a message written for a human: %q", err.Error())
	}
}

// The mistake this guards against, and it is one I made: surfacing every 4xx
// payload. pkg/http/openapi/transport FABRICATES an RPCStatus for bodies that
// were never RPCStatus - filling Message with the raw upstream text, an HTML
// page or "token is expired\n" - and marks them INVALID_ARGUMENT, never 9.
// Relaying those would undo grandturks#1092.
func TestASynthesisedPayloadDoesNotDisplaceTheCannedText(t *testing.T) {
	err := Parse(fakePayloadError{
		code: 400,
		payload: &models.RPCStatus{
			Code:    codeInvalidArgument,
			Message: "<html><body>403 Forbidden</body></html>",
			// The label the transport stamps on anything it invented. Keyed on
			// provenance, not on code: a genuine INVALID_ARGUMENT from the
			// server carries the same 3.
			Details: []*models.ProtobufAny{{AtType: SynthesizedBodyMarker}},
		},
	}, false)

	if !strings.Contains(err.Error(), "Bad Parameter.") {
		t.Errorf("a fabricated payload displaced the canned text: %q", err.Error())
	}
	if strings.Contains(err.Error(), "<html>") {
		t.Errorf("raw upstream body reached the user: %q", err.Error())
	}
}

// A 5xx is the server's problem, and its internals are not the caller's
// business even when it says something.
func TestServerInternalsAreNotRelayed(t *testing.T) {
	err := Parse(fakePayloadError{
		code:    500,
		payload: &models.RPCStatus{Code: codeFailedPrecondition, Message: `pq: relation "projects" does not exist`},
	}, false)

	if !strings.Contains(err.Error(), "Internal error") {
		t.Errorf("expected the canned 500 text, got %q", err.Error())
	}
	if strings.Contains(err.Error(), "does not exist") {
		t.Errorf("a database error reached the user: %q", err.Error())
	}
}

// And with no payload at all, nothing changes.
func TestNoPayloadKeepsTheExistingBehaviour(t *testing.T) {
	err := Parse(fakePayloadError{code: 403, payload: nil}, false)
	if !strings.Contains(err.Error(), "Permission denied.") {
		t.Errorf("expected the canned 403 text, got %q", err.Error())
	}
}

// The gap this follow-up closes: the error a user is far likelier to hit than
// a version refusal. It is INVALID_ARGUMENT, the same code the transport
// stamps on anything it fabricates - which is why provenance, not code, is
// what decides.
func TestAGenuineInvalidArgumentMessageReachesTheUser(t *testing.T) {
	const missing = "pipeline.params: missing required field :model_name"

	err := Parse(fakePayloadError{
		code:    400,
		payload: &models.RPCStatus{Code: codeInvalidArgument, Message: missing},
	}, false)

	if !strings.Contains(err.Error(), missing) {
		t.Errorf("the server named the missing field and the user was not told: %q", err.Error())
	}
	if strings.Contains(err.Error(), "Bad Parameter.") {
		t.Errorf("the canned string replaced a message that named the problem: %q", err.Error())
	}
}
