package io

import (
	"encoding/base64"
	"io"
)

type Base64 struct{}

// ReadAll implements read all from io.Reader and convert it into base64
func (b *Base64) ReadAll(r io.Reader) (string, error) {
	blob, err := io.ReadAll(r)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(blob), nil
}

var (
	DefaultBase64Stub = &Base64{}
)

func ReadAll2Base64(r io.Reader) (string, error) {
	return DefaultBase64Stub.ReadAll(r)
}

func MustReadAll2Base64(r io.Reader) string {
	s, err := DefaultBase64Stub.ReadAll(r)
	if err != nil {
		panic(err)
	}
	return s
}
