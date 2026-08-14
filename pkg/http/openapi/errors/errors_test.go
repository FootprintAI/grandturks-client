package openapierrors

import (
	"errors"
	"strings"
	"testing"
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
