package encryption

import (
	"bytes"
	"crypto/cipher"
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

func (a aesEncryptor) Decode(encryptedtext []byte) ([]byte, error) {
	ecb := cipher.NewCBCDecrypter(a.block, []byte(a.initialVector))
	decrypted := make([]byte, len(encryptedtext))
	ecb.CryptBlocks(decrypted, encryptedtext)
	return pKCS5Trimming(decrypted), nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5Trimming(encrypt []byte) []byte {
	if len(encrypt) == 0 {
		return encrypt
	}
	padding := encrypt[len(encrypt)-1]
	return encrypt[:len(encrypt)-int(padding)]
}
