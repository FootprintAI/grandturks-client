// Package transport carries the http.RoundTripper the CLI wraps its client
// with, so that an error status reaches pkg/http/openapi/errors.Parse with its
// code intact.
//
// This is the CLI-side half of grandturks#1092. The other half lives in
// FootprintAI/grandturks at common/http/openapi, which does the same thing for
// that repo's service-to-service clients - two copies for the same reason
// errors.go has two: the CLI builds from THIS module, and a fix that only
// lands in the other one never reaches a user (grandturks#1010 -> #1085 is the
// worked example, and grandturks' drift_test.go exists because of it).
package transport

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// maxNormalizedBody caps how much of an unreadable error body is carried into
// the reconstructed one.
//
// The body being normalized here is, by definition, not something this API
// wrote - an ingress' HTML error page, a proxy's plain text, a mux's "404 page
// not found". Some of those are large, and the whole of one is never the
// useful part. 4 KiB is enough for any of them to be recognisable, and the cut
// is announced in the message rather than made silently.
const maxNormalizedBody = 4 << 10

// New wraps rt so that non-2xx responses reach the generated client in a shape
// it can deserialize.
//
// nil rt means http.DefaultTransport, mirroring http.Client's own handling -
// calling RoundTrip on a nil RoundTripper panics.
func New(rt http.RoundTripper) http.RoundTripper {
	if rt == nil {
		rt = http.DefaultTransport
	}
	return normalizeErrorBodies{next: rt}
}

// normalizeErrorBodies makes a non-2xx response the generated clients can
// actually deserialize (grandturks#1092).
//
// # THE PROBLEM
//
// Every go-swagger reader in this repo has this shape:
//
//	default:
//	    result := NewXDefault(response.Code())
//	    if err := result.readResponse(response, consumer, o.formats); err != nil {
//	        return nil, err            // <- the status code is DROPPED here
//	    }
//	    return nil, result
//
// When the body cannot be deserialized into that component's RPCStatus, the
// caller gets the CONSUMER's error - which carries no Code() - instead of the
// *XDefault that does. openapierrors.Parse returns any error that is not an
// openapiError verbatim, so every message in its switch is unreachable on
// exactly the responses that do not carry a well-formed error body:
//
//	Error: &{0 [] } (*models.RPCStatus) is not supported by the TextConsumer,
//	       can be resolved by supporting TextUnmarshaler interface
//
// which reads as a client/server contract bug. In the reported case it was an
// expired token, and the first three hypotheses were a dependency regression,
// a server outage, and a deployment problem - none of them true, and none of
// them the user's to fix.
//
// # WHY HERE
//
// The status is present and correct in the *http.Response; it is lost one
// layer above, inside generated code that must not be hand-edited. A
// RoundTripper sits below that layer and can hand the reader a body it can
// read, after which the existing switch works unchanged - for 401, and equally
// for the 403, 404 and 502 that were just as opaque.
//
// The alternative is matching on the consumer's error string, which is what
// FootprintAI/manifests' lifecycle_test.py does today and what #1092 calls a
// workaround: it breaks when go-openapi rewords the message, and it can never
// tell a 401 from a 500.
//
// # WHAT IT WILL NOT DO
//
// It never touches a 2xx, and it never touches a body the client can already
// read. Rewriting a well-formed error would replace the API's own message with
// this package's reconstruction of it, which is a worse defect than the one
// being fixed - the server's words are the part an operator needs.
type normalizeErrorBodies struct {
	next http.RoundTripper
}

func (n normalizeErrorBodies) RoundTrip(req *http.Request) (*http.Response, error) {
	resp, err := n.next.RoundTrip(req)
	if err != nil || resp == nil || resp.Body == nil {
		return resp, err
	}
	// 1xx/2xx/3xx are somebody else's business. Redirects are followed by the
	// http.Client above this, so what arrives here is the final response.
	if resp.StatusCode < 400 {
		return resp, nil
	}

	// One byte past the cap, so a body exactly at the cap is not reported as
	// truncated.
	body, readErr := io.ReadAll(io.LimitReader(resp.Body, maxNormalizedBody+1))
	closeErr := resp.Body.Close()
	if readErr != nil {
		// The body could not be read at all - a connection dropped mid-response.
		// Say that, rather than pretending the status arrived with no content:
		// "the server said 500 and nothing else" and "the connection died while
		// it was talking" are different problems.
		return replaceBody(resp, resp.StatusCode, fmt.Sprintf("could not read the response body: %v", readErr)), nil
	}
	_ = closeErr

	truncated := len(body) > maxNormalizedBody
	if truncated {
		body = body[:maxNormalizedBody]
	}

	// An EMPTY body already deserializes into a zero-valued RPCStatus, so the
	// reader already produces an *XDefault for it and Parse already works.
	// Leaving it alone keeps this change to the responses that are actually
	// broken.
	if len(bytes.TrimSpace(body)) == 0 || readsAsRPCStatus(body) {
		resp.Body = io.NopCloser(bytes.NewReader(body))
		return resp, nil
	}

	message := string(bytes.TrimSpace(body))
	if truncated {
		message += fmt.Sprintf(" […truncated at %d bytes]", maxNormalizedBody)
	}
	return replaceBody(resp, resp.StatusCode, message), nil
}

// readsAsRPCStatus reports whether the generated client could deserialize this
// body into its RPCStatus.
//
// Checked STRUCTURALLY rather than against a models package, because there is
// no single one to check against: every component generates its own RPCStatus
// (authz, userStore, kafeido, ...), all with the same shape, and this package
// is the transport all of them share. Importing one component's models here
// would be both a layering violation and wrong for the other nine.
//
// The probe below is that shared shape - grpc-gateway's `rpcStatus` - with
// every field a pointer or a slice so that an absent field is not a mismatch.
// A JSON array, an HTML page, or an object whose `code` is a string all fail
// it, which is exactly the set that fails in the generated reader.
func readsAsRPCStatus(body []byte) bool {
	var probe struct {
		Code    *int32            `json:"code"`
		Message *string           `json:"message"`
		Details []json.RawMessage `json:"details"`
	}
	return json.Unmarshal(body, &probe) == nil
}

// replaceBody swaps in an RPCStatus the reader can deserialize.
//
// Content-Length is corrected and Content-Type forced to JSON: leaving either
// stale would trade a deserialization failure for a truncated read or for the
// runtime picking the text consumer again, which is the bug this exists to
// remove.
func replaceBody(resp *http.Response, httpStatus int, message string) *http.Response {
	// Marshalled rather than fmt-formatted, so a message containing a quote or
	// a newline - an HTML page always does - cannot produce invalid JSON and
	// land straight back in the failure being fixed.
	payload, err := json.Marshal(struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
		Details []any  `json:"details"`
	}{
		Code:    grpcCodeForHTTPStatus(httpStatus),
		Message: message,
		Details: []any{},
	})
	if err != nil {
		// Unreachable for these types; a bare object still deserializes, which
		// keeps the status reaching Parse even if the words are lost.
		payload = []byte(`{"code":2,"message":"","details":[]}`)
	}

	resp.Body = io.NopCloser(bytes.NewReader(payload))
	resp.ContentLength = int64(len(payload))
	resp.Header.Set("Content-Type", "application/json")
	resp.Header.Set("Content-Length", fmt.Sprintf("%d", len(payload)))
	// Any encoding applied to the ORIGINAL body does not describe this one.
	resp.Header.Del("Content-Encoding")
	return resp
}

// grpcCodeForHTTPStatus fills in the `code` field with the canonical gRPC code
// for a status, per grpc-gateway's own mapping.
//
// Nothing in this repo switches on it - openapierrors.Parse reads the HTTP
// status off the response, not this field - so it is presentation only. It is
// filled in anyway because a hardcoded 0 is OK/"no error" beside a 500, and a
// reader who quotes it would be quoting something untrue.
func grpcCodeForHTTPStatus(status int) int32 {
	switch status {
	case http.StatusBadRequest:
		return 3 // INVALID_ARGUMENT
	case http.StatusUnauthorized:
		return 16 // UNAUTHENTICATED
	case http.StatusForbidden:
		return 7 // PERMISSION_DENIED
	case http.StatusNotFound:
		return 5 // NOT_FOUND
	case http.StatusConflict:
		return 6 // ALREADY_EXISTS
	case http.StatusTooManyRequests:
		return 8 // RESOURCE_EXHAUSTED
	case http.StatusNotImplemented:
		return 12 // UNIMPLEMENTED
	case http.StatusServiceUnavailable, http.StatusBadGateway, http.StatusGatewayTimeout:
		return 14 // UNAVAILABLE
	case http.StatusInternalServerError:
		return 13 // INTERNAL
	default:
		return 2 // UNKNOWN
	}
}
