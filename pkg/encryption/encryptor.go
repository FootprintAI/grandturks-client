package encryption

import "encoding/base64"

type Encryptor interface {
	Encode(plaintext []byte) ([]byte, error)
	Decode(encryptedplaintext []byte) ([]byte, error)
}

type Encryption struct {
	encryptor Encryptor
	codec     StrEncodeDecoder
}

func (e *Encryption) EncodeStr(plaintext []byte) (string, error) {
	encodedMessage, err := e.encryptor.Encode(plaintext)
	if err != nil {
		return "", err
	}
	return e.codec.EncodeToString(encodedMessage), nil
}

func (e *Encryption) DecodeStr(encryptedBase64Str string) ([]byte, error) {
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
