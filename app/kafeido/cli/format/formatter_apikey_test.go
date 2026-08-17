package format

import (
	"bytes"
	"encoding/json"
	"strings"
	"testing"

	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

func testAPIKeyInfo() *appmodels.AppkafeidoAPIKeyInfo {
	role := appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEREAD
	return &appmodels.AppkafeidoAPIKeyInfo{
		KeyID: "0123456789abcdef",
		Name:  "acme-ci",
		Role:  &role,
	}
}

// TestCreatedApiKeyFormatterEmitsOnlyTheFormat keeps `--format=json` machine
// readable.
//
// An api key exists for integrations, and an integration mints one by piping
// this command into a parser. A human-readable warning appended to the
// formatter's output - which is what this formatter did first - turns valid
// JSON into
//
//	invalid character 'S' after top-level value
//
// The warning still has to be said, so it is said on stderr by the command
// rather than on stdout by the formatter.
func TestCreatedApiKeyFormatterEmitsOnlyTheFormat(t *testing.T) {
	const token = "gtk_0123456789abcdef_c2VjcmV0"

	original := DefaultTypeFormatter
	t.Cleanup(func() { DefaultTypeFormatter = original })
	DefaultTypeFormatter = TypeFormatJSON

	var buf bytes.Buffer
	if err := NewCreatedApiKeyFormatter(testAPIKeyInfo(), token).Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var rows []map[string]string
	if err := json.Unmarshal(buf.Bytes(), &rows); err != nil {
		t.Fatalf("output is not parseable json (%v), output was:\n%s", err, buf.String())
	}
	if len(rows) != 1 {
		t.Fatalf("got %d rows, want 1", len(rows))
	}
	if rows[0]["token"] != token {
		t.Errorf("token = %q, want %q", rows[0]["token"], token)
	}
}

// TestCreatedApiKeyFormatterShowsTheTokenInTableForm: the token is returned
// exactly once by the server and is never recoverable, so the one command that
// can print it must actually print it.
func TestCreatedApiKeyFormatterShowsTheTokenInTableForm(t *testing.T) {
	const token = "gtk_0123456789abcdef_c2VjcmV0"

	original := DefaultTypeFormatter
	t.Cleanup(func() { DefaultTypeFormatter = original })
	DefaultTypeFormatter = TypeFormatTable

	var buf bytes.Buffer
	if err := NewCreatedApiKeyFormatter(testAPIKeyInfo(), token).Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if !strings.Contains(buf.String(), token) {
		t.Errorf("table output does not contain the token:\n%s", buf.String())
	}
}

// TestListApiKeyDetailsFormatterRendersAbsentTimestamps: a key with no expiry
// and one that has never been used are both meaningful states, and an empty
// cell reads as missing data rather than as either of them.
func TestListApiKeyDetailsFormatterRendersAbsentTimestamps(t *testing.T) {
	original := DefaultTypeFormatter
	t.Cleanup(func() { DefaultTypeFormatter = original })
	DefaultTypeFormatter = TypeFormatJSON

	var buf bytes.Buffer
	if err := NewListApiKeyDetailsFormatter([]*appmodels.AppkafeidoAPIKeyInfo{testAPIKeyInfo()}).Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}

	var rows []map[string]string
	if err := json.Unmarshal(buf.Bytes(), &rows); err != nil {
		t.Fatalf("output is not parseable json (%v), output was:\n%s", err, buf.String())
	}
	if rows[0]["expires_at"] != "never" {
		t.Errorf("expires_at = %q, want %q", rows[0]["expires_at"], "never")
	}
	if rows[0]["last_used_at"] != "never" {
		t.Errorf("last_used_at = %q, want %q", rows[0]["last_used_at"], "never")
	}
	if rows[0]["role"] != "read" {
		t.Errorf("role = %q, want the word a user types into --role", rows[0]["role"])
	}
}

// TestListApiKeyDetailsFormatterSkipsNilKeys: the generated model is a slice
// of pointers, so a nil element is representable and must not panic the
// listing.
func TestListApiKeyDetailsFormatterSkipsNilKeys(t *testing.T) {
	var buf bytes.Buffer
	keys := []*appmodels.AppkafeidoAPIKeyInfo{nil, testAPIKeyInfo(), nil}
	if err := NewListApiKeyDetailsFormatter(keys).Write(&buf); err != nil {
		t.Fatalf("Write: %v", err)
	}
	if !strings.Contains(buf.String(), "acme-ci") {
		t.Errorf("the one real key is missing from the output:\n%s", buf.String())
	}
}
