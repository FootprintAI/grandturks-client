package encryption

type StrEncodeDecoder interface {
	EncodeToString([]byte) string
	DecodeString(string) ([]byte, error)
}
