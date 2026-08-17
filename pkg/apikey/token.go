// Package apikey holds the client half of the api key credential: its shape,
// and the header it travels in.
//
// It deliberately mirrors, rather than imports, grandturks'
// common/apikey/token.go - this module is the CLI and does not depend on the
// server. The rules below are the server's, so a key this package accepts is
// one the server will at least attempt to look up, and a key it rejects is one
// that could not have been minted.
//
// Nothing here verifies anything. A well-formed token from an attacker parses
// fine; parsing exists so a mistyped key fails locally with a message that
// says what was expected, instead of as a 401 indistinguishable from an
// expired login.
package apikey

import (
	"encoding/hex"
	"errors"
	"strings"
)

const (
	// Prefix makes a key recognisable on sight, greppable in a support bundle,
	// and registrable as a pattern with a secret scanner.
	Prefix = "gtk"

	sep = "_"

	// idHexLen is the id's width: 8 random bytes, hex-encoded. Hex rather than
	// base64url deliberately - base64url's alphabet contains "_", which is the
	// separator, so a base64url id would make the token ambiguous to parse.
	idHexLen = 16
)

// HTTPHeader is the header an api key MUST be sent in (grandturks#1175).
//
// Istio's RequestAuthentication inspects `Authorization: Bearer <token>` and
// rejects anything that will not parse as a JWT, before the request reaches
// any service. A gtk_ token is not a JWT, so a key sent in Authorization dies
// at the ingress with
//
//	401 Jwt is not in the form of Header.Payload.Signature with two dots and 3 sections
//
// on every route. The server accepts both headers - it promotes this one into
// the bearer credential before authentication runs - which means Authorization
// works in-cluster and fails through a gateway: the split that looks like a
// working feature in dev and a broken one in production. Istio does not look
// at this header, so it passes through untouched and needs no ingress change.
const HTTPHeader = "X-Api-Key"

var (
	ErrMalformed = errors.New("apikey: malformed token")

	// ErrEmptySecret is separate from ErrMalformed because "gtk_<id>_" is
	// well-formed in shape and empty where it matters. Collapsing the two
	// would report "malformed" for a token that is structurally fine and
	// dangerously empty.
	ErrEmptySecret = errors.New("apikey: token carries no secret")
)

// Parse splits a presented token and returns its public key id.
//
// The secret is intentionally not returned: nothing in this CLI needs it apart
// from the header value, which is the whole token. The id is safe to print -
// on its own it authenticates nothing - and is what `list apikey` and `delete
// apikey` speak in.
func Parse(token string) (id string, err error) {
	// SplitN with n=3 matters. The secret is base64url, whose alphabet
	// includes "_", so it may itself contain separators; taking the remainder
	// as a single field is what keeps such a secret from being truncated into
	// an unparseable token.
	parts := strings.SplitN(token, sep, 3)
	if len(parts) != 3 || parts[0] != Prefix {
		return "", ErrMalformed
	}
	if len(parts[1]) != idHexLen {
		return "", ErrMalformed
	}
	if _, decErr := hex.DecodeString(parts[1]); decErr != nil {
		return "", ErrMalformed
	}
	if parts[2] == "" {
		return "", ErrEmptySecret
	}
	return parts[1], nil
}
