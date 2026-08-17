package encryption

// The oauth2 callback credential: the blob the authentication service puts in
// the `credentials` query parameter of the loopback redirect that ends
// `kafeido login`, and the CLI reads back.
//
// Designed in docs/architecture/oauth2-callback-credential.md, which is worth
// reading before changing anything here. The short version of why it exists:
//
//   - What it replaces is AES-CBC with no authentication tag, so a modified
//     ciphertext decrypted to *something* rather than failing
//     (grandturks-client#29), and two shapes of malformed input panicked
//     outright (#23, grandturks#1204).
//   - That scheme's key was a source literal compiled into a distributed
//     binary, and its IV was fixed - so it was deterministic, and offered no
//     confidentiality against anyone who had the CLI.
//
// Here the CLI mints an ephemeral X25519 keypair per login and sends the
// public half in the login request; the server seals to it with a keypair of
// its own. Nothing durable is stored, there is no build-time secret, and a
// public key travelling through the browser is harmless - that is the property
// a public key has and a shared symmetric key does not.
//
// This package is the ONE implementation. grandturks seals with SealCredential
// and the CLI opens with CredentialKey.OpenCredential, so the two ends cannot
// drift the way the duplicated AES-CBC code did.

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"errors"
	"fmt"
	"io"

	"golang.org/x/crypto/hkdf"
)

const (
	// CredentialMagic marks the format and its version. A decoder dispatches
	// on it: anything else is the legacy CBC blob. A legacy blob opens with
	// these four bytes with probability 2^-32, and the consequence is a decode
	// error and a retried login rather than a wrong plaintext.
	CredentialMagic = "GTE1"

	credentialMagicLen = len(CredentialMagic)
	// X25519 public keys and shared secrets are 32 bytes.
	credentialKeyLen = 32
	// GCM standard nonce size; cipher.NewGCM's default.
	credentialNonceLen = 12
	credentialTagLen   = 16

	credentialMinLen = credentialMagicLen + credentialKeyLen + credentialNonceLen + credentialTagLen

	// credentialKDFInfo domain-separates this key derivation from any other
	// use of the same ECDH output.
	credentialKDFInfo = "grandturks/oauth2-callback v1"
)

var (
	// ErrBadPublicKey is returned when the peer's public key is not a usable
	// X25519 point. It arrives over the wire, so it is attacker-influenced.
	ErrBadPublicKey = errors.New("encryption: invalid credential public key")

	// ErrNotSealedCredential is returned for input that is not this format at
	// all - including the legacy CBC blob, which the caller should route to
	// the legacy decoder rather than treat as corrupt.
	ErrNotSealedCredential = errors.New("encryption: not a sealed credential")
)

// CredentialKey is one login's ephemeral keypair, held by the CLI for the few
// seconds between asking for a login and the browser calling back.
//
// The private half never leaves the process and is never written to disk.
type CredentialKey struct {
	priv *ecdh.PrivateKey
}

// NewCredentialKey mints the keypair for a single login.
func NewCredentialKey() (*CredentialKey, error) {
	priv, err := ecdh.X25519().GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("encryption: generating credential key: %w", err)
	}
	return &CredentialKey{priv: priv}, nil
}

// PublicKey is the 32 bytes to put in the login request.
//
// A copy, so a caller cannot mutate the key its reply will be sealed to.
func (k *CredentialKey) PublicKey() []byte {
	pub := k.priv.PublicKey().Bytes()
	out := make([]byte, len(pub))
	copy(out, pub)
	return out
}

// SealCredential encrypts plaintext to the holder of clientPublicKey. It is
// what the authentication service calls; it lives here so both ends of the
// exchange are the same code.
//
// requestID is authenticated but not encrypted: binding it means a credential
// captured from one login cannot be opened in another.
func SealCredential(clientPublicKey []byte, requestID string, plaintext []byte) (string, error) {
	if len(clientPublicKey) != credentialKeyLen {
		return "", fmt.Errorf("%w: got %d bytes, want %d", ErrBadPublicKey, len(clientPublicKey), credentialKeyLen)
	}
	clientPub, err := ecdh.X25519().NewPublicKey(clientPublicKey)
	if err != nil {
		// Rejects the all-zero point and anything else X25519 will not accept.
		return "", fmt.Errorf("%w: %s", ErrBadPublicKey, err)
	}

	serverPriv, err := ecdh.X25519().GenerateKey(rand.Reader)
	if err != nil {
		return "", fmt.Errorf("encryption: generating ephemeral key: %w", err)
	}
	shared, err := serverPriv.ECDH(clientPub)
	if err != nil {
		return "", fmt.Errorf("%w: %s", ErrBadPublicKey, err)
	}
	serverPub := serverPriv.PublicKey().Bytes()

	gcm, err := newCredentialAEAD(shared)
	if err != nil {
		return "", err
	}
	nonce := make([]byte, credentialNonceLen)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", fmt.Errorf("encryption: reading nonce: %w", err)
	}

	// header is both the blob's prefix and the additional authenticated data,
	// so the version marker and the ephemeral key cannot be altered or
	// stripped to force a downgrade.
	header := make([]byte, 0, credentialMagicLen+credentialKeyLen)
	header = append(header, CredentialMagic...)
	header = append(header, serverPub...)

	out := make([]byte, 0, credentialMinLen+len(plaintext))
	out = append(out, header...)
	out = append(out, nonce...)
	out = gcm.Seal(out, nonce, plaintext, credentialAAD(header, requestID))

	return base64.URLEncoding.EncodeToString(out), nil
}

// OpenCredential decrypts what SealCredential produced for this key.
//
// Total by contract: any string in, a value or an error out, never a panic.
// The scheme this replaces panicked on two shapes of malformed input, and the
// input is a query parameter on a local HTTP server anything can reach.
func (k *CredentialKey) OpenCredential(requestID, encoded string) ([]byte, error) {
	raw, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return nil, fmt.Errorf("encryption: credential is not base64url: %w", err)
	}
	if !isSealedCredentialBytes(raw) {
		return nil, ErrNotSealedCredential
	}

	serverPub, err := ecdh.X25519().NewPublicKey(raw[credentialMagicLen : credentialMagicLen+credentialKeyLen])
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBadPublicKey, err)
	}
	shared, err := k.priv.ECDH(serverPub)
	if err != nil {
		return nil, fmt.Errorf("%w: %s", ErrBadPublicKey, err)
	}

	gcm, err := newCredentialAEAD(shared)
	if err != nil {
		return nil, err
	}
	header := raw[:credentialMagicLen+credentialKeyLen]
	nonce := raw[len(header) : len(header)+credentialNonceLen]
	ciphertext := raw[len(header)+credentialNonceLen:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, credentialAAD(header, requestID))
	if err != nil {
		// Deliberately one message for every failure - wrong key, tampered
		// bytes, wrong request id. Which one it was is not something the
		// caller can act on, and saying is how a padding oracle starts.
		return nil, errors.New("encryption: credential failed authentication")
	}
	return plaintext, nil
}

// IsSealedCredential reports whether encoded is this format, so a caller can
// route anything else to the legacy decoder.
func IsSealedCredential(encoded string) bool {
	raw, err := base64.URLEncoding.DecodeString(encoded)
	if err != nil {
		return false
	}
	return isSealedCredentialBytes(raw)
}

func isSealedCredentialBytes(raw []byte) bool {
	return len(raw) >= credentialMinLen && string(raw[:credentialMagicLen]) == CredentialMagic
}

// credentialAAD is the authenticated-but-not-encrypted data: the header
// (marker and ephemeral public key) plus the request id.
func credentialAAD(header []byte, requestID string) []byte {
	aad := make([]byte, 0, len(header)+len(requestID))
	aad = append(aad, header...)
	return append(aad, requestID...)
}

// newCredentialAEAD derives the message key from the ECDH output and returns
// the AEAD to use with it.
//
// HKDF rather than using the shared secret directly: the X25519 output is a
// curve point, not a uniformly distributed key, and the info string keeps this
// derivation distinct from any other use of the same exchange.
func newCredentialAEAD(shared []byte) (cipher.AEAD, error) {
	key := make([]byte, credentialKeyLen)
	if _, err := io.ReadFull(hkdf.New(sha256.New, shared, nil, []byte(credentialKDFInfo)), key); err != nil {
		return nil, fmt.Errorf("encryption: deriving credential key: %w", err)
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, fmt.Errorf("encryption: creating cipher: %w", err)
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, fmt.Errorf("encryption: creating aead: %w", err)
	}
	return gcm, nil
}
