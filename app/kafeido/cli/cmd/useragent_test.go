package cmd

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The point of the header is that a server can tell a kafeido CLI from
// anything else, and tell one build from another. Both halves are asserted,
// because a User-Agent that says only "kafeido-cli" answers "is anyone still
// on the old prefix" but not "can we retire it yet".
func TestUserAgentIdentifiesTheClientAndTheBuild(t *testing.T) {
	ua := UserAgent()
	assert.True(t, strings.HasPrefix(ua, "kafeido-cli/"),
		"a server reading its access log must be able to pick these out: %q", ua)
	assert.NotEqual(t, "kafeido-cli/ ()", ua, "version and commit must actually be interpolated")
}

func TestUserAgentIsSentOnEveryRequest(t *testing.T) {
	var got string
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		got = r.Header.Get("User-Agent")
	}))
	defer srv.Close()

	client := &http.Client{Transport: &userAgentTransport{rt: http.DefaultTransport}}
	resp, err := client.Get(srv.URL)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Equal(t, UserAgent(), got)
}

// A RoundTripper must not mutate the request it is handed - the caller may
// still be holding it, and a retry would otherwise re-send an object this
// transport had already edited.
func TestUserAgentTransportDoesNotMutateTheCallersRequest(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
	require.NoError(t, err)

	tr := &userAgentTransport{rt: http.DefaultTransport}
	resp, err := tr.RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	assert.Empty(t, req.Header.Get("User-Agent"),
		"the caller's request was modified in place")
}

// A zero-value transport must still work rather than nil-panic: it is one
// missing field away in any future wiring change.
func TestUserAgentTransportFallsBackToTheDefaultRoundTripper(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {}))
	defer srv.Close()

	req, err := http.NewRequest(http.MethodGet, srv.URL, nil)
	require.NoError(t, err)

	resp, err := (&userAgentTransport{}).RoundTrip(req)
	require.NoError(t, err)
	defer resp.Body.Close()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
