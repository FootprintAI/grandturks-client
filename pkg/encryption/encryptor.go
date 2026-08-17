package encryption

import (
	"encoding/base64"
	"errors"
)

// errNoEncryptor is what a build with no injected encryptor reports.
//
// The concrete encryptor is injected - the kafeido CLI's is set by grandturks'
// main via SetEncryptor, because the key is shared with the authentication
// service and does not belong in this module. A binary built from this repo
// alone has none, and the oauth2 callback used to call straight into the nil
// interface and panic in its handler goroutine (grandturks-client#33).
var errNoEncryptor = errors.New("encryption: no encryptor configured - this build cannot encrypt or decrypt; " +
	"use the kafeido CLI distributed by grandturks")

type Encryptor interface {
	Encode(plaintext []byte) ([]byte, error)
	Decode(encryptedplaintext []byte) ([]byte, error)
}

type Encryption struct {
	encryptor Encryptor
	codec     StrEncodeDecoder
}

func (e *Encryption) EncodeStr(plaintext []byte) (string, error) {
	if e.encryptor == nil {
		return "", errNoEncryptor
	}
	encodedMessage, err := e.encryptor.Encode(plaintext)
	if err != nil {
		return "", err
	}
	return e.codec.EncodeToString(encodedMessage), nil
}

func (e *Encryption) DecodeStr(encryptedBase64Str string) ([]byte, error) {
	if e.encryptor == nil {
		return nil, errNoEncryptor
	}
	encryptedRawMessage, err := e.codec.DecodeString(encryptedBase64Str)
	if err != nil {
		return nil, err
	}

	return e.encryptor.Decode(encryptedRawMessage)
}

func NewEncryption(encryptor Encryptor) *Encryption {
	return &Encryption{
		encryptor: encryptor,
		codec:     base64.URLEncoding,
	}
}
