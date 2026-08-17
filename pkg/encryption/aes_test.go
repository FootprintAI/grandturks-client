package encryption

import (
	"bytes"
	"crypto/aes"
	"encoding/base64"
	"strings"
	"testing"
)

// This file stopped compiling in 6e439b9 (2024-05-14), when the implementation
// moved out from under it in four separate places (grandturks-client#16): the
// aesKey fixture went away, NewAESEncryptor took (cipher.Block, string)
// instead, and EncodeStr/DecodeStr moved off aesEncryptor onto Encryption.
// Nothing reported it, because `go build ./...` stays green on a broken _test
// file and this repo had no CI to run `go test` (grandturks-client#15).
//
// It is restored rather than deleted on purpose. This is the AES-CBC encryptor
// the CLI uses for credential handling - app/kafeido/cli/cmd/cmd_login.go
// feeds the oauth2 callback's `credentials` query parameter straight into
// Encryption.DecodeStr - so deleting the file would have closed the build
// error and widened the gap it stands for.
//
// Two things the old test did that these do not: it built its cipher from
// []byte("cipher-for-unittest"), which is 19 bytes and not a legal AES key
// size, and it discarded the error saying so - so rawCipherBlock was nil and
// the test could only ever have panicked had it compiled. The keys below are
// 16/24/32 bytes, and every aes.NewCipher error is checked.

const (
	// One AES block. CBC requires the IV to be exactly BlockSize.
	testInitialVector = "1234567890123456"
	testKey256        = "0123456789abcdef0123456789abcdef"
)

func mustCipher(t *testing.T, key string) *aesEncryptorFixture {
	t.Helper()
	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		t.Fatalf("aes.NewCipher(%d-byte key): %v", len(key), err)
	}
	e := NewAESEncryptor(block, testInitialVector)
	return &aesEncryptorFixture{encryptor: e, encryption: NewEncryption(e)}
}

type aesEncryptorFixture struct {
	encryptor  aesEncryptor
	encryption *Encryption
}

// TestEncodeDecodeRoundTrip is the assertion the package never had: what
// Encode produces, Decode gives back - across the plaintext lengths where
// PKCS#5 padding is easy to get wrong (empty, sub-block, exactly one block,
// exactly two blocks, and either side of a block boundary).
func TestEncodeDecodeRoundTrip(t *testing.T) {
	blockSize := aes.BlockSize

	for _, tc := range []struct {
		name      string
		plaintext []byte
	}{
		{"empty", []byte("")},
		{"single byte", []byte("x")},
		{"helloworld", []byte("helloworld")},
		{"one byte short of a block", bytes.Repeat([]byte("a"), blockSize-1)},
		{"exactly one block", bytes.Repeat([]byte("b"), blockSize)},
		{"one byte over a block", bytes.Repeat([]byte("c"), blockSize+1)},
		{"exactly two blocks", bytes.Repeat([]byte("d"), 2*blockSize)},
		{"non-ascii", []byte("台北 ☕ kafeido")},
		{"binary with nul bytes", []byte{0x00, 0x01, 0x00, 0xff, 0x7f, 0x00}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := mustCipher(t, testKey256)

			ciphertext, err := f.encryptor.Encode(tc.plaintext)
			if err != nil {
				t.Fatalf("Encode(%q) returned error: %v", tc.plaintext, err)
			}
			if len(ciphertext)%blockSize != 0 {
				t.Errorf("len(ciphertext) = %d, want a multiple of the %d-byte block size",
					len(ciphertext), blockSize)
			}
			// A padded plaintext is always strictly longer than the plaintext,
			// so this also catches an Encode that forgot to encrypt at all.
			if bytes.Equal(ciphertext, tc.plaintext) {
				t.Errorf("Encode returned its input unchanged (%q) - nothing was encrypted", ciphertext)
			}

			got, err := f.encryptor.Decode(ciphertext)
			if err != nil {
				t.Fatalf("Decode returned error: %v", err)
			}
			if !bytes.Equal(got, tc.plaintext) {
				t.Errorf("round trip = %q, want %q", got, tc.plaintext)
			}
		})
	}
}

// TestRoundTripAcrossAESKeySizes covers all three legal key lengths, since the
// old fixture's illegal 19-byte key is what made the file worth reading twice.
func TestRoundTripAcrossAESKeySizes(t *testing.T) {
	for _, tc := range []struct {
		name string
		key  string
	}{
		{"aes-128", "0123456789abcdef"},
		{"aes-192", "0123456789abcdef01234567"},
		{"aes-256", testKey256},
	} {
		t.Run(tc.name, func(t *testing.T) {
			f := mustCipher(t, tc.key)
			message := []byte("helloworld")

			ciphertext, err := f.encryptor.Encode(message)
			if err != nil {
				t.Fatalf("Encode: %v", err)
			}
			got, err := f.encryptor.Decode(ciphertext)
			if err != nil {
				t.Fatalf("Decode: %v", err)
			}
			if !bytes.Equal(got, message) {
				t.Errorf("round trip = %q, want %q", got, message)
			}
		})
	}
}

// TestEncryptionStrRoundTrip exercises the path the CLI actually runs:
// Encryption wraps the encryptor with base64url, and EncodeStr/DecodeStr -
// the two methods the old test called on the wrong type - live here.
func TestEncryptionStrRoundTrip(t *testing.T) {
	f := mustCipher(t, testKey256)
	// Shaped like what cmd_login.go round-trips: a url.Values query string.
	message := []byte("reqId=abc-123&timestamp=1700000000&token=ya29.a0")

	encoded, err := f.encryption.EncodeStr(message)
	if err != nil {
		t.Fatalf("EncodeStr: %v", err)
	}
	if strings.Contains(encoded, string(message)) {
		t.Errorf("EncodeStr output %q contains the plaintext", encoded)
	}
	if _, err := base64.URLEncoding.DecodeString(encoded); err != nil {
		t.Errorf("EncodeStr output %q is not base64url: %v", encoded, err)
	}

	got, err := f.encryption.DecodeStr(encoded)
	if err != nil {
		t.Fatalf("DecodeStr: %v", err)
	}
	if !bytes.Equal(got, message) {
		t.Errorf("round trip = %q, want %q", got, message)
	}
}

// TestDecodeStrRejectsMalformedBase64 pins the error path. DecodeStr's input
// is the `credentials` query parameter of a request to the CLI's local oauth2
// callback server, so a value that is not base64url at all is reachable from
// outside the process and must come back as an error.
func TestDecodeStrRejectsMalformedBase64(t *testing.T) {
	f := mustCipher(t, testKey256)

	for _, tc := range []struct {
		name  string
		input string
	}{
		{"not base64", "not base64!!"},
		{"base64 with standard-alphabet chars", "a+b/c="},
		// A fourth case belongs here - {"truncated", "YWJj"}, valid base64url
		// that decodes to 3 bytes - but Decode PANICS on it today
		// ("crypto/cipher: input not full blocks") instead of returning an
		// error, and so does any ciphertext that decrypts to garbage padding.
		// Filed as grandturks-client#23 rather than fixed here, since #16 is
		// about restoring these tests; the case goes in with that fix.
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, err := f.encryption.DecodeStr(tc.input); err == nil {
				t.Errorf("DecodeStr(%q) = nil error, want an error", tc.input)
			}
		})
	}
}

// TestDecodeWithWrongInitialVector pins that the IV is load-bearing: CBC
// XORs it into the first block, so decrypting with a different one must not
// return the plaintext.
func TestDecodeWithWrongInitialVector(t *testing.T) {
	block, err := aes.NewCipher([]byte(testKey256))
	if err != nil {
		t.Fatalf("aes.NewCipher: %v", err)
	}
	message := bytes.Repeat([]byte("e"), 2*aes.BlockSize)

	ciphertext, err := NewAESEncryptor(block, testInitialVector).Encode(message)
	if err != nil {
		t.Fatalf("Encode: %v", err)
	}
	got, err := NewAESEncryptor(block, "6543210987654321").Decode(ciphertext)
	if err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if bytes.Equal(got, message) {
		t.Error("decoding with a different initial vector returned the plaintext")
	}
}
