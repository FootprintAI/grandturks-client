package apikey

import (
	"errors"
	"strings"
	"testing"
)

// TestParse pins the token grammar this CLI accepts, which is a mirror of the
// one the server mints against (grandturks common/apikey/token.go). The cases
// that matter are the ones a naive split would get wrong.
func TestParse(t *testing.T) {
	const validID = "0123456789abcdef"

	for _, tc := range []struct {
		name    string
		token   string
		wantID  string
		wantErr error
	}{
		{
			name:   "well formed",
			token:  "gtk_" + validID + "_c2VjcmV0dmFsdWU",
			wantID: validID,
		},
		{
			// The secret is base64url, whose alphabet includes "_" - the same
			// character that separates the fields. Splitting on every "_"
			// truncates such a secret and rejects a perfectly good key, so the
			// split has to stop after the second separator.
			name:   "secret containing the separator",
			token:  "gtk_" + validID + "_c2Vj_cmV0_dmFsdWU",
			wantID: validID,
		},
		{
			// hex.DecodeString accepts either case, so an upper-case id is a
			// real id and must not be refused.
			name:   "upper case hex id",
			token:  "gtk_0123456789ABCDEF_c2VjcmV0",
			wantID: "0123456789ABCDEF",
		},
		{"empty", "", "", ErrMalformed},
		{"no separators", "gtk", "", ErrMalformed},
		{"missing secret field", "gtk_" + validID, "", ErrMalformed},
		{"wrong prefix", "sk_" + validID + "_c2VjcmV0", "", ErrMalformed},
		{"bearer token pasted by mistake", "eyJhbGciOiJIUzI1NiJ9.e30.abc", "", ErrMalformed},
		{"empty id", "gtk__c2VjcmV0", "", ErrMalformed},
		{"short id", "gtk_0123_c2VjcmV0", "", ErrMalformed},
		{"long id", "gtk_0123456789abcdef00_c2VjcmV0", "", ErrMalformed},
		{"non hex id", "gtk_zzzzzzzzzzzzzzzz_c2VjcmV0", "", ErrMalformed},
		{
			// Structurally fine and empty where it counts. Reported apart from
			// ErrMalformed so a caller cannot log "malformed" for a token that
			// is shaped correctly and carries no credential at all.
			name:    "empty secret",
			token:   "gtk_" + validID + "_",
			wantErr: ErrEmptySecret,
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			gotID, err := Parse(tc.token)
			if !errors.Is(err, tc.wantErr) {
				t.Fatalf("Parse(%q) error = %v, want %v", tc.token, err, tc.wantErr)
			}
			if gotID != tc.wantID {
				t.Errorf("Parse(%q) id = %q, want %q", tc.token, gotID, tc.wantID)
			}
		})
	}
}

// TestParseErrorDoesNotLeakTheSecret: Parse's error is printed to the terminal
// and pasted into bug reports. A key that is malformed is still a credential -
// the holder may simply have mistyped the id - so nothing that came after the
// prefix belongs in the message.
func TestParseErrorDoesNotLeakTheSecret(t *testing.T) {
	const secret = "sup3rs3cr3tvalue"

	_, err := Parse("gtk_badid_" + secret)
	if err == nil {
		t.Fatal("Parse of a malformed token returned no error")
	}
	if got := err.Error(); strings.Contains(got, secret) {
		t.Errorf("error message %q contains the secret", got)
	}
}

// TestHTTPHeaderIsNotAuthorization is a one-line test for the whole reason
// this package has a header constant (grandturks#1175): Istio's
// RequestAuthentication parses any Authorization bearer token as a JWT and
// rejects what will not parse, so a gtk_ key sent there dies at the ingress
// with "Jwt is not in the form of Header.Payload.Signature" before reaching
// any service. It works in-cluster and fails through the gateway, which is
// the worst possible split, and this constant is what avoids it.
func TestHTTPHeaderIsNotAuthorization(t *testing.T) {
	if HTTPHeader != "X-Api-Key" {
		t.Errorf("HTTPHeader = %q, want %q - the server reads this exact header", HTTPHeader, "X-Api-Key")
	}
}
