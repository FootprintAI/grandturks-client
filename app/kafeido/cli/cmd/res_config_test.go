package cmd

import (
	"strings"
	"testing"
)

// isValidEndpoint used to require the https scheme and nothing else, with no
// exception for loopback (grandturks#901).
//
// Requiring https for a REMOTE endpoint is right, and stays: the CLI sends a
// bearer token, and this is what stops someone pointing it at a plaintext
// endpoint and leaking credentials. The bug was that the check was
// scheme-only, so it could not tell "plaintext to a remote host" (dangerous)
// from "plaintext to the developer's own machine" (fine).
//
// The consequence was that the CLI could not be configured against any of the
// project's own stacks through its own documented command - the sandbox binds
// SeaweedFS S3 on 127.0.0.1:8333 and the demo webportal on 127.0.0.1:3000, both
// plain HTTP. The only ways in were hand-writing ~/.kafeidoconfig.json or
// setting env vars whose names contain a literal dot. That is plausibly why
// grandturks#897 had to seed the config file directly rather than exercise
// `config set endpoint` as a user would.
func TestIsValidEndpoint(t *testing.T) {
	for _, tc := range []struct {
		name     string
		endpoint string
		wantErr  bool
		why      string
	}{
		// https is always fine, loopback or not.
		{"https remote", "https://api.example.com", false,
			"https to a remote host is the normal case"},
		{"https with port", "https://api.example.com:8443", false,
			"an explicit port must not change the scheme decision"},
		{"https loopback", "https://127.0.0.1:8443", false,
			"https to loopback was already allowed and must stay allowed"},

		// http to loopback is the exception this adds.
		{"http loopback v4", "http://127.0.0.1:8333", false,
			"the sandbox binds SeaweedFS S3 here over plain HTTP"},
		{"http loopback v4 other in 127/8", "http://127.0.0.2:9000", false,
			"the whole 127.0.0.0/8 block is loopback, not just .1"},
		{"http localhost", "http://localhost:3000", false,
			"the demo webportal is reached this way"},
		{"http localhost uppercase", "http://LOCALHOST:3000", false,
			"host names are case-insensitive; a capitalised one is the same machine"},
		{"http loopback v6", "http://[::1]:8080", false,
			"::1 is loopback and must survive the brackets-and-port parsing"},

		// http to anywhere else is still refused - this is the whole point of
		// the original check.
		{"http remote", "http://api.example.com", true,
			"plaintext to a remote host leaks the bearer token - still refused"},
		{"http remote ip", "http://203.0.113.10:8080", true,
			"a routable address is not loopback however it is spelled"},
		{"http private lan", "http://192.168.1.10:8080", true,
			"a LAN address is still someone else's network"},

		// Near-misses. These are the reason the host is parsed and compared
		// properly rather than matched by prefix or substring.
		{"host merely starting with a loopback ip", "http://127.0.0.1.evil.com", true,
			"a prefix match here would hand the token to evil.com"},
		{"host merely containing localhost", "http://localhost.evil.com", true,
			"same trick with the name instead of the address"},
		{"host ending in localhost", "http://evil-localhost", true,
			"a suffix match would be just as wrong"},

		// Any other scheme is refused, as before.
		{"ftp", "ftp://127.0.0.1", true, "only http and https are endpoints"},
		{"no scheme", "127.0.0.1:8333", true,
			"unchanged behaviour: a bare host:port has no scheme to trust"},
		{"empty", "", true, "an empty endpoint is not a valid one"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			err := isValidEndpoint(tc.endpoint)
			if tc.wantErr && err == nil {
				t.Errorf("isValidEndpoint(%q) = nil, want an error - %s", tc.endpoint, tc.why)
			}
			if !tc.wantErr && err != nil {
				t.Errorf("isValidEndpoint(%q) = %v, want nil - %s", tc.endpoint, err, tc.why)
			}
		})
	}
}

// The refusal has to say what would be accepted. The old message was "require
// https", which is not actionable for someone pointing the CLI at their own
// sandbox - it reads as "never http" rather than "not http to a remote host".
func TestRemoteHTTPErrorNamesTheLoopbackException(t *testing.T) {
	err := isValidEndpoint("http://api.example.com")
	if err == nil {
		t.Fatal("http to a remote host must be refused")
	}
	for _, want := range []string{"https", "loopback"} {
		if !strings.Contains(err.Error(), want) {
			t.Errorf("refusal %q does not mention %q - a caller cannot tell from this "+
				"whether their own machine is reachable", err.Error(), want)
		}
	}
}
