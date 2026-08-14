package openapierrors

import (
	"errors"
	"fmt"
)

type openapiError interface {
	Code() int
	Error() string
}

// errRedacted stands in for the underlying error when the caller asked not to
// surface details. It reaches the user verbatim, inside "(details:...)", so it
// is a user-facing string despite looking like an internal sentinel.
//
// Was "<detained>", which is not a word that applies to a suppressed error
// detail and read as a bug in its own right to anyone who quoted it
// (grandturks-client#18).
var (
	errRedacted = errors.New("<redacted>")
)

func Parse(err error, hasDetail bool) error {
	_, isOpenapiErrorType := err.(openapiError)
	if !isOpenapiErrorType {
		return err
	}
	var actualErr error = err
	if !hasDetail {
		actualErr = errRedacted
	}
	openapierr := err.(openapiError)
	switch openapierr.Code() {
	case 200:
		return nil
	case 401:
		return newError("Token Expired. Require Login first.", actualErr)
	case 400:
		return newError("Bad Parameter.", actualErr)
	case 403:
		return newError("Permission denied. You either don't have enough permission or haven't login first.", actualErr)
	case 404:
		return newError("Resource not found.", actualErr)
	case 409:
		return newError("Status Conflicted.", actualErr)
	case 500:
		return newError("Internal error. Please contact your system administrator.", actualErr)
	case 504:
		return newError("Bad Gateway. Please try later.", actualErr)
	default:
		return newError("unknown error", actualErr)
	}
}

func newError(plainText string, err error) *Error {
	return &Error{
		PlainText: plainText,
		Details:   err,
	}
}

type Error struct {
	PlainText string
	Details   error
}

func (m *Error) Error() string {
	return fmt.Sprintf("%s(details:%s)", m.PlainText, m.Details.Error())
}
