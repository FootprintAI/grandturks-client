package cmd

// Step 3 of the rollout in docs/architecture/oauth2-callback-credential.md:
// the CLI side of #29.
//
// The login command mints an ephemeral X25519 keypair per login, and the
// callback handler routes whatever arrives in the `credentials` query
// parameter to the decoder that matches its format. Two formats exist at once
// on purpose, and for as long as it takes: a CLI talking to a server that has
// not been updated gets the legacy AES-CBC blob, and one talking to an updated
// server gets GTE1. Neither side has to know which the other is.

import (
	"strings"
	"testing"

	"github.com/footprintai/grandturks-client/v2/pkg/encryption"
)

const testCallbackRequestID = "0f4b9a2e-7c31-4f5a-9d0e-2b7a1c6e5d43"

// fixedKeyEncryptor is a stand-in for the encryptor grandturks' main injects
// via SetEncryptor. This module deliberately ships no key of its own - the
// legacy one is shared with the authentication service and lives there.
type fixedKeyEncryptor struct{ shift byte }

func (f fixedKeyEncryptor) Encode(plaintext []byte) ([]byte, error) {
	out := make([]byte, len(plaintext))
	for i, b := range plaintext {
		out[i] = b + f.shift
	}
	return out, nil
}

func (f fixedKeyEncryptor) Decode(ciphertext []byte) ([]byte, error) {
	out := make([]byte, len(ciphertext))
	for i, b := range ciphertext {
		out[i] = b - f.shift
	}
	return out, nil
}

func withLegacyEncryptor(t *testing.T, e encryption.Encryptor) {
	t.Helper()
	previous := GetEncryptor()
	t.Cleanup(func() { SetEncryptor(previous) })
	SetEncryptor(e)
}

// TestDecodeCallbackCredentialRoutesByFormat is the dispatch itself. The
// marker decides, not configuration and not a flag: a blob the server sealed
// to this login's key opens with that key, and anything else goes to the
// legacy decoder, which is what every server that predates #29 sends.
func TestDecodeCallbackCredentialRoutesByFormat(t *testing.T) {
	withLegacyEncryptor(t, fixedKeyEncryptor{shift: 3})

	key, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	payload := "reqId=" + testCallbackRequestID + "&token=ya29.a0&timestamp=1700000000"

	t.Run("sealed", func(t *testing.T) {
		sealed, err := encryption.SealCredential(key.PublicKey(), testCallbackRequestID, []byte(payload))
		if err != nil {
			t.Fatalf("SealCredential: %v", err)
		}

		got, err := decodeCallbackCredential(key, testCallbackRequestID, sealed)
		if err != nil {
			t.Fatalf("decodeCallbackCredential: %v", err)
		}
		if string(got) != payload {
			t.Errorf("plaintext = %q, want %q", got, payload)
		}
	})

	t.Run("legacy", func(t *testing.T) {
		legacy, err := encryption.NewEncryption(GetEncryptor()).EncodeStr([]byte(payload))
		if err != nil {
			t.Fatalf("EncodeStr: %v", err)
		}

		got, err := decodeCallbackCredential(key, testCallbackRequestID, legacy)
		if err != nil {
			t.Fatalf("decodeCallbackCredential: %v", err)
		}
		if string(got) != payload {
			t.Errorf("plaintext = %q, want %q", got, payload)
		}
	})
}

// TestDecodeCallbackCredentialSealedPathNeedsNoInjectedEncryptor: the sealed
// format carries its own key exchange, so it does not touch the injected
// legacy encryptor at all. A build without one - which is any binary built
// from this module (#33) - can still finish a login against a server that
// seals.
func TestDecodeCallbackCredentialSealedPathNeedsNoInjectedEncryptor(t *testing.T) {
	withLegacyEncryptor(t, nil)

	key, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	sealed, err := encryption.SealCredential(key.PublicKey(), testCallbackRequestID, []byte("reqId=x&token=y"))
	if err != nil {
		t.Fatalf("SealCredential: %v", err)
	}

	got, err := decodeCallbackCredential(key, testCallbackRequestID, sealed)
	if err != nil {
		t.Fatalf("decodeCallbackCredential: %v", err)
	}
	if string(got) != "reqId=x&token=y" {
		t.Errorf("plaintext = %q", got)
	}
}

// TestDecodeCallbackCredentialRejectsGarbage: the handler's error branch has
// to be reachable. This is the input an unrelated local process can send to
// the callback port, and it is what used to panic (#23, grandturks#1204).
func TestDecodeCallbackCredentialRejectsGarbage(t *testing.T) {
	withLegacyEncryptor(t, fixedKeyEncryptor{shift: 3})

	key, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}

	for _, tc := range []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"not base64", "not base64!!"},
		{"marker but truncated", "R1RFMQ=="},
		{"sealed to another key", sealedToAnotherKey(t)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := decodeCallbackCredential(key, testCallbackRequestID, tc.input)
			if err == nil {
				t.Fatalf("decodeCallbackCredential(%q) = %q, want an error", tc.input, got)
			}
		})
	}
}

// TestDecodeCallbackCredentialBindsTheRequestID: a sealed blob from an earlier
// login does not open in a later one, which is checked before the plaintext is
// parsed rather than after.
func TestDecodeCallbackCredentialBindsTheRequestID(t *testing.T) {
	withLegacyEncryptor(t, fixedKeyEncryptor{shift: 3})

	key, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	sealed, err := encryption.SealCredential(key.PublicKey(), "an-earlier-login", []byte("reqId=x&token=y"))
	if err != nil {
		t.Fatalf("SealCredential: %v", err)
	}

	if _, err := decodeCallbackCredential(key, testCallbackRequestID, sealed); err == nil {
		t.Error("a credential from another login was accepted")
	}
}

// TestNewCredentialKeyIsFreshPerLogin: the key is per login, so two logins
// cannot open each other's callbacks. Cheap to assert, and the whole scheme
// rests on it.
func TestNewCredentialKeyIsFreshPerLogin(t *testing.T) {
	first, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	second, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}

	if string(first.PublicKey()) == string(second.PublicKey()) {
		t.Fatal("two logins got the same keypair")
	}

	sealed, err := encryption.SealCredential(first.PublicKey(), testCallbackRequestID, []byte("reqId=x&token=y"))
	if err != nil {
		t.Fatalf("SealCredential: %v", err)
	}
	if _, err := second.OpenCredential(testCallbackRequestID, sealed); err == nil {
		t.Error("one login's key opened another login's credential")
	}
}

// TestDecodeCallbackCredentialErrorDoesNotLeakTheCredential: the handler
// writes this error into an HTTP response body that the browser renders, so
// whatever arrived must not be echoed back into it.
func TestDecodeCallbackCredentialErrorDoesNotLeakTheCredential(t *testing.T) {
	withLegacyEncryptor(t, fixedKeyEncryptor{shift: 3})

	key, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	const attacker = "R1RFMQthis-is-not-a-real-credential-but-it-is-attacker-supplied"

	_, err = decodeCallbackCredential(key, testCallbackRequestID, attacker)
	if err == nil {
		t.Fatal("expected an error")
	}
	if strings.Contains(err.Error(), attacker) {
		t.Errorf("error %q echoes the supplied credential", err)
	}
}

func sealedToAnotherKey(t *testing.T) string {
	t.Helper()
	other, err := encryption.NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	sealed, err := encryption.SealCredential(other.PublicKey(), testCallbackRequestID, []byte("reqId=x&token=y"))
	if err != nil {
		t.Fatalf("SealCredential: %v", err)
	}
	return sealed
}
