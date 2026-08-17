package cmd

import (
	"errors"

	"github.com/go-openapi/strfmt"

	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/pkg/encryption"
)

// decodeCallbackCredential turns the `credentials` query parameter of the
// oauth2 callback into the plaintext the login flow parses.
//
// Two formats coexist, deliberately and for as long as it takes
// (docs/architecture/oauth2-callback-credential.md):
//
//   - GTE1, sealed to this login's ephemeral key. It carries its own key
//     exchange, so it needs nothing injected and cannot be read by anyone who
//     merely holds a copy of the CLI.
//   - the legacy AES-CBC blob, which is what a server that predates #29
//     sends, decrypted with the encryptor grandturks' main injects via
//     SetEncryptor.
//
// The marker in the blob decides, so neither side needs to know the other's
// version - which is what lets the two repositories roll out separately.
func decodeCallbackCredential(key *encryption.CredentialKey, requestID, credentials string) ([]byte, error) {
	if len(credentials) == 0 {
		// Never valid, and worth its own error: an absent parameter and a
		// corrupt one are different mistakes, and the injected legacy
		// encryptor is not guaranteed to distinguish them.
		return nil, errors.New("login: the callback carried no credentials parameter")
	}
	// Dispatch on the MARKER, not on well-formedness: a truncated blob that
	// claims to be a sealed credential must be rejected as one rather than
	// handed to the legacy decoder, which would report the wrong error and is
	// a downgrade path.
	if encryption.HasCredentialMarker(credentials) {
		return key.OpenCredential(requestID, credentials)
	}
	return encryption.NewEncryption(GetEncryptor()).DecodeStr(credentials)
}

// newOauth2LoginRequest builds the login request, including the public half of
// this login's ephemeral keypair.
//
// Sending the key is what asks the server for the sealed callback format: a
// request without one is how every CLI older than #29 looks, and gets the
// legacy AES-CBC blob in reply. So this field is both the key and the
// capability signal, and dropping it silently downgrades the login rather than
// failing it - which is why it has a test of its own.
func newOauth2LoginRequest(localRedirectURL, requestID string, key *encryption.CredentialKey) *appmodels.KafeidoAppOauth2LoginRequest {
	return &appmodels.KafeidoAppOauth2LoginRequest{
		LocalRedirectURL:    localRedirectURL,
		RequestID:           requestID,
		CredentialPublicKey: strfmt.Base64(key.PublicKey()),
	}
}
