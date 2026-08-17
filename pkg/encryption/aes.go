package encryption

import (
	"bytes"
	"crypto/cipher"
	"errors"
	"fmt"
)

type aesEncryptor struct {
	block         cipher.Block
	initialVector string
}

var (
	_ Encryptor = aesEncryptor{}
)

func NewAESEncryptor(block cipher.Block, initialVector string) aesEncryptor {
	return aesEncryptor{
		block:         block,
		initialVector: initialVector,
	}
}

func (a aesEncryptor) Encode(plaintext []byte) ([]byte, error) {
	ecb := cipher.NewCBCEncrypter(a.block, []byte(a.initialVector))
	paddedPlaintext := pKCS5Padding(plaintext, a.block.BlockSize())
	encryptedtext := make([]byte, len(paddedPlaintext))
	ecb.CryptBlocks(encryptedtext, paddedPlaintext)
	return encryptedtext, nil

}

// Decode reverses Encode.
//
// Both guards below exist because this function's input is not internal
// (grandturks-client#23): app/kafeido/cli/cmd/cmd_login.go passes the
// `credentials` query parameter of a request to the CLI's local oauth2
// callback server straight into Encryption.DecodeStr, and handles the error
// this returns. Before these checks there was no error to handle - the
// handler goroutine panicked and took the CLI down mid-login.
func (a aesEncryptor) Decode(encryptedtext []byte) ([]byte, error) {
	blockSize := a.block.BlockSize()
	// cipher.CBCDecrypter.CryptBlocks panics on an input that is not a whole
	// number of blocks, before anything here can inspect it.
	if len(encryptedtext) == 0 || len(encryptedtext)%blockSize != 0 {
		return nil, fmt.Errorf("encryption: ciphertext of %d bytes is not a whole number of %d-byte blocks",
			len(encryptedtext), blockSize)
	}
	ecb := cipher.NewCBCDecrypter(a.block, []byte(a.initialVector))
	decrypted := make([]byte, len(encryptedtext))
	ecb.CryptBlocks(decrypted, encryptedtext)
	return pKCS5Trimming(decrypted, blockSize)
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

// pKCS5Trimming removes the padding Encode added, after checking that it is
// padding at all.
//
// The unchecked version took the last byte as a length and sliced by it, so a
// ciphertext that decrypted to garbage - anything encrypted under a different
// key, or simply made up - produced a negative index and a panic: 468 of 500
// wrong-key decodes in the measurement on #23, with the other 32 returning
// silently-wrong plaintext.
//
// Every padding byte is checked, not just the last, which is what turns those
// 32 into errors too. This is not constant time and does not need to be: the
// caller is a single-shot local callback, not a decryption oracle an attacker
// can query repeatedly.
func pKCS5Trimming(decrypted []byte, blockSize int) ([]byte, error) {
	if len(decrypted) == 0 {
		return nil, errors.New("encryption: nothing to trim")
	}
	padding := int(decrypted[len(decrypted)-1])
	if padding == 0 || padding > blockSize || padding > len(decrypted) {
		return nil, errors.New("encryption: ciphertext is not correctly padded")
	}
	for _, b := range decrypted[len(decrypted)-padding:] {
		if int(b) != padding {
			return nil, errors.New("encryption: ciphertext is not correctly padded")
		}
	}
	return decrypted[:len(decrypted)-padding], nil
}
