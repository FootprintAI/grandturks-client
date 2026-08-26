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

	// A SERVER-AUTHORED INSTRUCTION WINS, and only that.
	//
	// The strings below are generic by necessity: they are all this package
	// knows when the server said nothing specific. But when the server wrote
	// an instruction for a human, replacing it with "Bad Parameter." throws
	// away the only actionable part - and Details is redacted unless --debug,
	// so it was not merely buried, it was gone.
	//
	// The case that prompted this: a deployment enforcing a minimum CLI
	// version answers with
	//
	//   kafeido CLI 2.6.0 is too old for this deployment, which requires
	//   2.7.0 or newer. Upgrade with: go install ...@v2.7.0
	//
	// and the user saw "Bad Parameter.(details:<redacted>)".
	//
	// NARROW ON PURPOSE - gated on FAILED_PRECONDITION, not on 4xx generally.
	// I tried 4xx first and it broke pkg/http/openapi/transport's tests, which
	// were right to fail it: that layer FABRICATES an RPCStatus for responses
	// whose body was never RPCStatus at all, filling Message with the raw
	// upstream body - "token is expired\n", or an entire HTML error page from
	// a proxy. Surfacing those instead of "Token Expired. Require Login
	// first." is strictly worse, and is the grandturks#1092 failure mode the
	// transport exists to remove.
	//
	// FAILED_PRECONDITION distinguishes them without a marker or a heuristic:
	// grpcCodeForHTTPStatus maps 400 to INVALID_ARGUMENT, so the transport
	// never synthesises a 9. A 9 can only have come from a server that chose
	// it - and it means precisely "the caller must change something before
	// retrying", which is exactly when its own words are the useful ones.
	// Both gates. The gRPC code says the server chose this deliberately; the
	// HTTP status says it is the caller's to fix. A 5xx carrying a 9 is still
	// the server's problem and its internals are still not the caller's
	// business - which a test here caught me getting wrong.
	const codeFailedPrecondition = 9
	if code >= 400 && code < 500 {
		if payload := serverPayload(err); payload != nil &&
			payload.Code == codeFailedPrecondition && payload.Message != "" {
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
