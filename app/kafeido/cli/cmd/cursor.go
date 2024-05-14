package cmd

import (
	b64 "encoding/base64"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

var (
	errInvalidCursor = errors.New("invalid cursor")
)

func decodeCursor(cursor string) (maxKeysReturned int, nextKey string, err error) {
	composedKey, err := b64.StdEncoding.DecodeString(cursor)
	if err != nil {
		return 0, "", errInvalidCursor
	}
	tokens := strings.SplitN(string(composedKey), ":", 2)
	if len(tokens) != 2 {
		return 0, "", errInvalidCursor
	}
	nextKey = tokens[1]
	maxKeysReturned, err = strconv.Atoi(tokens[0])
	if err != nil {
		err = errInvalidCursor
		return
	}
	return
}

func encodeCursor(maxKeysReturned int, nextKey string) string {
	encodedKey := fmt.Sprintf("%d:%s", maxKeysReturned, nextKey)
	return b64.StdEncoding.EncodeToString([]byte(encodedKey))
}
