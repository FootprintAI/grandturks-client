package encryption

import (
	"strings"
	"testing"
)

// TestEncryptionWithoutAnEncryptor: the concrete encryptor is INJECTED - the
// kafeido CLI's is set by grandturks' main via SetEncryptor, because the key
// is shared with the authentication service and does not belong in this
// module. A binary built from this repo alone never sets one, so
// GetEncryptor() returns a nil interface and the oauth2 callback handler
// called straight into it:
//
//	panic: runtime error: invalid memory address or nil pointer dereference
//
// in the handler goroutine, mid-login (grandturks-client#33). A build that
// cannot decrypt has to say so.
func TestEncryptionWithoutAnEncryptor(t *testing.T) {
	e := NewEncryption(nil)

	t.Run("DecodeStr", func(t *testing.T) {
		got, err := e.DecodeStr("YWJjZGVmZ2hpamtsbW5vcA==")
		if err == nil {
			t.Fatalf("DecodeStr = %q, want an error rather than a panic", got)
		}
		if !strings.Contains(err.Error(), "encryptor") {
			t.Errorf("error %q does not say what is missing", err)
		}
	})

	t.Run("EncodeStr", func(t *testing.T) {
		if _, err := e.EncodeStr([]byte("anything")); err == nil {
			t.Error("EncodeStr = nil error, want an error rather than a panic")
		}
	})
}
