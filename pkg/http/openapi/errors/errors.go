package openapierrors

import (
	"errors"
	"fmt"

	"github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

type openapiError interface {
	Code() int
	Error() string
}

// payloadCarrier is every generated *XDefault: go-swagger gives them all a
// GetPayload returning the RPCStatus grpc-gateway serialises an error into.
// That is where a server-authored message lives.
//
// Typed rather than scraped out of Error(), whose string embeds the same JSON.
// pkg/http/openapi/transport/errorbody.go calls string-matching on error text
// "a #1092 problem", and it would be exactly that here.
type payloadCarrier interface {
	GetPayload() *models.RPCStatus
}

// SynthesizedBodyMarker labels an RPCStatus that pkg/http/openapi/transport
// invented, rather than one a server sent.
//
// It travels in `details` because that is the only field of RPCStatus that can
// carry something extra without lying about the others: `code` is filled in
// honestly by grpcCodeForHTTPStatus, and overloading it with a sentinel would
// make a reader who quotes it quote something untrue.
const SynthesizedBodyMarker = "kafeido.dev/synthesized-error-body"

// isSynthesized reports whether the transport wrote this payload rather than a
// server. Its Message is then the raw upstream body - an HTML page, a plain
// sentence - and must not displace the canned text (grandturks#1092).
func isSynthesized(payload *models.RPCStatus) bool {
	for _, d := range payload.Details {
		if d != nil && d.AtType == SynthesizedBodyMarker {
			return true
		}
	}
	return false
}

// serverPayload returns the RPCStatus the server sent, or nil.
func serverPayload(err error) *models.RPCStatus {
	carrier, ok := err.(payloadCarrier)
	if !ok {
		return nil
	}
	return carrier.GetPayload()
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
	code := openapierr.Code()

	// A SERVER-AUTHORED MESSAGE WINS on a 4xx, unless the transport made it up.
	//
	// The strings below are generic by necessity: they are all this package
	// knows when the server said nothing specific. When it did say something,
	// replacing it throws away the only actionable part - and Details is
	// redacted unless --debug, so the message was not buried, it was gone.
	// Two real cases:
	//
	//   pipeline.params: missing required field :model_name   -> "Bad Parameter."
	//   kafeido CLI 2.6.0 is too old ... upgrade with ...     -> "Bad Parameter."
	//
	// WHY THE MARKER, rather than a code allow-list. pkg/http/openapi/transport
	// FABRICATES an RPCStatus for responses whose body was never one, filling
	// Message with the raw upstream text - "token is expired\n", or a proxy's
	// HTML error page - and stamping it with grpcCodeForHTTPStatus, which maps
	// 400 to INVALID_ARGUMENT. So a genuine INVALID_ARGUMENT from the server and
	// a fabricated one are indistinguishable by code, and relaying the
	// fabricated ones re-opens grandturks#1092.
	//
	// An earlier revision dodged that by allow-listing FAILED_PRECONDITION,
	// which no synthesised body carries. That worked but only for the version
	// check - it left "missing required field :model_name", the error a user
	// is far likelier to hit, still reading "Bad Parameter." The transport now
	// labels what it invents, so this can key on provenance rather than on a
	// code that means two different things.
	//
	// 4xx ONLY. A 5xx is the server's problem, not the caller's, and its
	// internals are not something to relay - which is what the canned
	// "contact your system administrator" is for.
	if code >= 400 && code < 500 {
		if payload := serverPayload(err); payload != nil &&
			payload.Message != "" && !isSynthesized(payload) {
			return newError(payload.Message, actualErr)
		}
	}

	switch code {
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
