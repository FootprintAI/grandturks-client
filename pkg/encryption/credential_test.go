package encryption

import (
	"bytes"
	"crypto/aes"
	"crypto/ecdh"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"testing"

	"github.com/footprintai/grandturks-client/v2/pkg/encryption/credentialvectors"
)

// The oauth2 callback credential, per
// docs/architecture/oauth2-callback-credential.md. The properties below are
// the ones the design's tamper matrix names; each is an assertion the AES-CBC
// scheme this replaces could not make at all.

const testRequestID = "0f4b9a2e-7c31-4f5a-9d0e-2b7a1c6e5d43"

func mustKey(t *testing.T) *CredentialKey {
	t.Helper()
	k, err := NewCredentialKey()
	if err != nil {
		t.Fatalf("NewCredentialKey: %v", err)
	}
	return k
}

func mustSeal(t *testing.T, k *CredentialKey, requestID string, plaintext []byte) string {
	t.Helper()
	sealed, err := SealCredential(k.PublicKey(), requestID, plaintext)
	if err != nil {
		t.Fatalf("SealCredential: %v", err)
	}
	return sealed
}

// TestCredentialRoundTrip: what the authentication service seals, the CLI that
// asked for it opens.
func TestCredentialRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name      string
		plaintext []byte
	}{
		{"empty", []byte("")},
		{"single byte", []byte("x")},
		// What the callback actually carries.
		{"callback payload", []byte("reqId=" + testRequestID + "&token=ya29.a0AfH6SM&timestamp=1700000000")},
		{"non-ascii", []byte("台北 ☕ kafeido")},
		{"binary with nul bytes", []byte{0x00, 0x01, 0x00, 0xff, 0x7f, 0x00}},
		{"large", bytes.Repeat([]byte("t"), 4096)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			k := mustKey(t)

			sealed := mustSeal(t, k, testRequestID, tc.plaintext)
			// Against the raw blob rather than its base64 text, and only for
			// payloads long enough for a match to mean something: a one-byte
			// plaintext appears in almost any base64 string by chance.
			if len(tc.plaintext) >= 8 {
				raw, decErr := base64.URLEncoding.DecodeString(sealed)
				if decErr != nil {
					t.Fatalf("sealed output is not base64url: %v", decErr)
				}
				if bytes.Contains(raw, tc.plaintext) {
					t.Error("sealed output contains the plaintext")
				}
			}

			got, err := k.OpenCredential(testRequestID, sealed)
			if err != nil {
				t.Fatalf("OpenCredential: %v", err)
			}
			if !bytes.Equal(got, tc.plaintext) {
				t.Errorf("round trip = %q, want %q", got, tc.plaintext)
			}
		})
	}
}

// TestSealIsNotDeterministic: the scheme this replaces used a fixed IV, so the
// same plaintext always produced the same blob. A fresh ephemeral key and
// nonce per call is what removes that.
func TestSealIsNotDeterministic(t *testing.T) {
	k := mustKey(t)
	plaintext := []byte("reqId=x&token=y&timestamp=1")

	first := mustSeal(t, k, testRequestID, plaintext)
	second := mustSeal(t, k, testRequestID, plaintext)

	if first == second {
		t.Error("sealing the same plaintext twice produced identical output")
	}
	for _, sealed := range []string{first, second} {
		got, err := k.OpenCredential(testRequestID, sealed)
		if err != nil || !bytes.Equal(got, plaintext) {
			t.Errorf("OpenCredential = %q, %v; want the plaintext back", got, err)
		}
	}
}

// TestOpenRejectsTamperedCredential is the tamper matrix: every region of the
// blob is covered by the tag or the AAD, so a flipped bit anywhere fails to
// open. This is the property unauthenticated CBC could not provide - there, a
// modified ciphertext decrypted to *something*.
func TestOpenRejectsTamperedCredential(t *testing.T) {
	k := mustKey(t)
	sealed := mustSeal(t, k, testRequestID, []byte("reqId=x&token=secret&timestamp=1"))

	raw, err := base64.URLEncoding.DecodeString(sealed)
	if err != nil {
		t.Fatalf("sealed output is not base64url: %v", err)
	}

	for _, tc := range []struct {
		region string
		index  int
	}{
		{"magic", 1},
		{"server public key", credentialMagicLen + 2},
		{"nonce", credentialMagicLen + credentialKeyLen + 3},
		{"ciphertext", credentialMagicLen + credentialKeyLen + credentialNonceLen + 1},
		{"tag", len(raw) - 1},
	} {
		t.Run(tc.region, func(t *testing.T) {
			tampered := append([]byte(nil), raw...)
			tampered[tc.index] ^= 0x01

			got, err := k.OpenCredential(testRequestID, base64.URLEncoding.EncodeToString(tampered))
			if err == nil {
				t.Fatalf("a flipped bit in the %s was accepted, returning %q", tc.region, got)
			}
			if got != nil {
				t.Errorf("got %q alongside the error, want nil", got)
			}
		})
	}
}

// TestOpenBindsTheRequestID: the request id is authenticated data, so a
// credential captured from one login cannot be replayed into another. Without
// this the only thing standing in the way is the handler's own reqId
// comparison, one layer above, on plaintext it has already decrypted.
func TestOpenBindsTheRequestID(t *testing.T) {
	k := mustKey(t)
	sealed := mustSeal(t, k, testRequestID, []byte("reqId=x&token=secret&timestamp=1"))

	if _, err := k.OpenCredential("a-different-request-id", sealed); err == nil {
		t.Error("a credential sealed for one request id opened under another")
	}
}

// TestOpenRejectsAnotherKeysCredential: only the CLI that generated the
// keypair can open what was sealed to it.
func TestOpenRejectsAnotherKeysCredential(t *testing.T) {
	mine, theirs := mustKey(t), mustKey(t)
	sealed := mustSeal(t, theirs, testRequestID, []byte("not for me"))

	if _, err := mine.OpenCredential(testRequestID, sealed); err == nil {
		t.Error("a credential sealed to another key opened")
	}
}

// TestOpenRejectsMalformedInput covers what arrives when the query parameter
// is not a credential at all. Every case must be an error - the panics in
// grandturks-client#23 are the reason this test exists in the shape it does.
func TestOpenRejectsMalformedInput(t *testing.T) {
	k := mustKey(t)
	valid, err := base64.URLEncoding.DecodeString(mustSeal(t, k, testRequestID, []byte("x")))
	if err != nil {
		t.Fatalf("decoding a sealed credential: %v", err)
	}

	for _, tc := range []struct {
		name  string
		input string
	}{
		{"empty", ""},
		{"not base64", "not base64!!"},
		{"base64 of nothing", base64.URLEncoding.EncodeToString(nil)},
		{"magic only", base64.URLEncoding.EncodeToString([]byte(CredentialMagic))},
		{"truncated before the public key", base64.URLEncoding.EncodeToString(valid[:credentialMagicLen+10])},
		{"truncated before the nonce", base64.URLEncoding.EncodeToString(valid[:credentialMagicLen+credentialKeyLen+2])},
		{"truncated inside the tag", base64.URLEncoding.EncodeToString(valid[:len(valid)-4])},
		{"legacy cbc blob", base64.URLEncoding.EncodeToString(bytes.Repeat([]byte{0x11}, 2*aes.BlockSize))},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := k.OpenCredential(testRequestID, tc.input)
			if err == nil {
				t.Fatalf("OpenCredential(%q) = %q, want an error", tc.input, got)
			}
			if got != nil {
				t.Errorf("got %q alongside the error, want nil", got)
			}
		})
	}
}

// TestIsSealedCredential is the dispatch the CLI will make in step 3 of the
// rollout: a blob with the marker is opened with the ephemeral key, anything
// else goes to the legacy decoder.
func TestIsSealedCredential(t *testing.T) {
	k := mustKey(t)

	for _, tc := range []struct {
		name  string
		input string
		want  bool
	}{
		{"sealed", mustSeal(t, k, testRequestID, []byte("x")), true},
		{"legacy cbc", base64.URLEncoding.EncodeToString(bytes.Repeat([]byte{0x11}, 2*aes.BlockSize)), false},
		{"empty", "", false},
		{"not base64", "not base64!!", false},
		{"magic but too short to be a credential", base64.URLEncoding.EncodeToString([]byte(CredentialMagic + "short")), false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := IsSealedCredential(tc.input); got != tc.want {
				t.Errorf("IsSealedCredential(%q) = %v, want %v", tc.input, got, tc.want)
			}
		})
	}
}

// TestSealRejectsABadPublicKey: the public key arrives over the wire from a
// client, so the server seals against attacker-influenced bytes.
func TestSealRejectsABadPublicKey(t *testing.T) {
	for _, tc := range []struct {
		name string
		key  []byte
	}{
		{"nil", nil},
		{"empty", []byte{}},
		{"too short", bytes.Repeat([]byte{0x01}, credentialKeyLen-1)},
		{"too long", bytes.Repeat([]byte{0x01}, credentialKeyLen+1)},
		{"all zeroes", make([]byte, credentialKeyLen)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := SealCredential(tc.key, testRequestID, []byte("x"))
			if err == nil {
				t.Fatalf("SealCredential with a %s key returned %q, want an error", tc.name, got)
			}
			if !errors.Is(err, ErrBadPublicKey) {
				t.Errorf("error = %v, want ErrBadPublicKey", err)
			}
		})
	}
}

// TestPublicKeyIsStable: the value is sent in the login request and then used
// to open the reply, so it cannot change between the two.
func TestPublicKeyIsStable(t *testing.T) {
	k := mustKey(t)

	first, second := k.PublicKey(), k.PublicKey()
	if !bytes.Equal(first, second) {
		t.Fatal("PublicKey returned different values on successive calls")
	}
	if len(first) != credentialKeyLen {
		t.Errorf("len(PublicKey()) = %d, want %d", len(first), credentialKeyLen)
	}

	// Handing the caller the internal array would let a caller mutate the key
	// the reply is sealed to.
	first[0] ^= 0xff
	if bytes.Equal(k.PublicKey(), first) {
		t.Error("PublicKey exposes internal state - mutating the result changed the key")
	}
}

// FuzzOpenCredential: the bug this whole design replaces was a panic on
// malformed input, found by four hand-written cases. Fuzzing is what says
// there is no fifth.
func FuzzOpenCredential(f *testing.F) {
	k, err := NewCredentialKey()
	if err != nil {
		f.Fatalf("NewCredentialKey: %v", err)
	}
	sealed, err := SealCredential(k.PublicKey(), testRequestID, []byte("reqId=x&token=y"))
	if err != nil {
		f.Fatalf("SealCredential: %v", err)
	}

	f.Add(sealed)
	f.Add("")
	f.Add("not base64!!")
	f.Add(base64.URLEncoding.EncodeToString([]byte(CredentialMagic)))
	f.Add(base64.URLEncoding.EncodeToString(bytes.Repeat([]byte{0x11}, 32)))

	f.Fuzz(func(t *testing.T, encoded string) {
		// The contract is total: any string in, a value or an error out,
		// never a panic.
		if _, err := k.OpenCredential(testRequestID, encoded); err == nil && encoded != sealed {
			t.Errorf("OpenCredential(%q) succeeded on input that was not the sealed credential", encoded)
		}
		IsSealedCredential(encoded)
	})
}

// TestGoldenVectors decodes blobs committed in
// pkg/encryption/credentialvectors, which grandturks' tests import and assert
// against too. If a change here alters the wire format, this fails on both
// sides of the boundary rather than in a login.
func TestGoldenVectors(t *testing.T) {
	privBytes, err := hex.DecodeString(credentialvectors.PrivateKeyHex)
	if err != nil {
		t.Fatalf("decoding the vector private key: %v", err)
	}
	priv, err := ecdh.X25519().NewPrivateKey(privBytes)
	if err != nil {
		t.Fatalf("parsing the vector private key: %v", err)
	}
	k := &CredentialKey{priv: priv}

	if len(credentialvectors.Vectors) == 0 {
		t.Fatal("no vectors")
	}
	for _, v := range credentialvectors.Vectors {
		t.Run(v.Name, func(t *testing.T) {
			got, err := k.OpenCredential(v.RequestID, v.Sealed)
			if err != nil {
				t.Fatalf("OpenCredential: %v", err)
			}
			if string(got) != v.Plaintext {
				t.Errorf("plaintext = %q, want %q", got, v.Plaintext)
			}
			// The same vector under a different request id must fail, which
			// is what pins the AAD binding into the vectors themselves.
			if _, err := k.OpenCredential(v.RequestID+"-not", v.Sealed); err == nil {
				t.Error("vector opened under the wrong request id")
			}
		})
	}
}

// TestHasCredentialMarker: dispatch is by MARKER, while IsSealedCredential
// also requires the blob to be long enough to be one. The difference matters
// for exactly one case - a value that carries the marker and is too short -
// and getting it wrong routes that value to the LEGACY decoder, which reports
// a misleading error and is a downgrade path something could aim for.
func TestHasCredentialMarker(t *testing.T) {
	k := mustKey(t)
	sealed := mustSeal(t, k, testRequestID, []byte("x"))
	truncated := base64.URLEncoding.EncodeToString([]byte(CredentialMagic))

	for _, tc := range []struct {
		name           string
		input          string
		wantMarker     bool
		wantWellFormed bool
	}{
		{"sealed", sealed, true, true},
		{"marker but too short", truncated, true, false},
		{"legacy cbc", base64.URLEncoding.EncodeToString(bytes.Repeat([]byte{0x11}, 2*aes.BlockSize)), false, false},
		{"empty", "", false, false},
		{"not base64", "not base64!!", false, false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if got := HasCredentialMarker(tc.input); got != tc.wantMarker {
				t.Errorf("HasCredentialMarker(%q) = %v, want %v", tc.input, got, tc.wantMarker)
			}
			if got := IsSealedCredential(tc.input); got != tc.wantWellFormed {
				t.Errorf("IsSealedCredential(%q) = %v, want %v", tc.input, got, tc.wantWellFormed)
			}
		})
	}
}
