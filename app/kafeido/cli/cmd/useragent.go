package cmd

import (
	"fmt"
	"net/http"

	"github.com/footprintai/grandturks-client/v2/pkg/version"
)

// userAgentTransport stamps every request with the CLI's identity.
//
// WHY (FootprintAI/manifests#308, and the lesson of
// FootprintAI/grandturks#1251).
//
// The API's path prefix is configurable as of this change, so a deployment can
// move off /api - which it must, because /api is also the Kubeflow central
// dashboard's. But moving is only half of it: the old prefix has to be retired
// eventually, and retiring it blind breaks every client still using it.
//
// #1251 is the cautionary case. Its committed AES key could not be rotated
// because "every kafeido CLI distributed since 2024 has these constants
// compiled in and nothing reports which CLI versions are deployed". The fix
// there was to log the credential format on every login, so retirement could
// rest on observation. This is the same move, one layer up.
//
// Deliberately not a header of our own invention: User-Agent is where this
// belongs, every proxy and access log already records it, and it costs a
// server nothing to start reading.
type userAgentTransport struct {
	rt http.RoundTripper
}

// UserAgent is the value sent. Version and commit both, because a version
// alone does not distinguish a release from a local build of the same tag -
// and during a migration "which build is that caller" is the whole question.
func UserAgent() string {
	return fmt.Sprintf("kafeido-cli/%s (%s)", version.GetVersion(), version.GetCommitHash())
}

func (t *userAgentTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone before mutating: RoundTrippers must not modify the request they
	// are given, and a retry that reuses it would otherwise accumulate state.
	r := req.Clone(req.Context())
	r.Header.Set("User-Agent", UserAgent())

	rt := t.rt
	if rt == nil {
		rt = http.DefaultTransport
	}
	return rt.RoundTrip(r)
}
