package encryption

import (
	"crypto/aes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAESEncodeDecode(t *testing.T) {

	rawCipherBlock, _ := aes.NewCipher([]byte("cipher-for-unittest"))
	initialVector := "1234567890123456"

	key := aesKey{
		block:         rawCipherBlock,
		initialVector: initialVector,
	}

	message := []byte("helloworld")
	e := NewAESEncryptor(key)
	encryptedMessage, err := e.EncodeStr(message)
	assert.NoError(t, err)
	fmt.Printf("%s\n", string(encryptedMessage))
	message2, err := e.DecodeStr(encryptedMessage)
	assert.NoError(t, err)
	assert.EqualValues(t, message, message2)
}
