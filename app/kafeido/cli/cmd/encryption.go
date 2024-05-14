package cmd

import "github.com/footprintai/grandturks-client/v2/pkg/encryption"

var (
	defaultEncryptor encryption.Encryptor
)

func SetEncryptor(e encryption.Encryptor) {
	defaultEncryptor = e
}

func GetEncryptor() encryption.Encryptor {
	return defaultEncryptor
}
