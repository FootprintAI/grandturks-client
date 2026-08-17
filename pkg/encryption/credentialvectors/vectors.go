// Package credentialvectors holds golden vectors for the oauth2 callback
// credential format.
//
// It exists so the two ends of the exchange can be tested against the SAME
// bytes. grandturks depends on this module, so its tests import this package
// and assert that what its authentication service seals matches what these
// vectors say, and that it can open them. Drift in either direction fails a
// test rather than a login.
//
// Deliberately a normal (non-_test) package: a _test file cannot be imported
// across modules, and a testdata file cannot be read from the module cache by
// another repository.
//
// The private key here is a fixed literal for reproducibility. It is a test
// fixture and nothing else: real keys are generated per login by
// encryption.NewCredentialKey and never leave the process.
package credentialvectors

// PrivateKeyHex is the X25519 private key every vector below is sealed to.
const PrivateKeyHex = "4f1a5c9e2d8b3706a1c4e59f8d2b603147ae9c58d0f3b26719ac4e8d5f2b0361"

// Vector is one sealed credential and what it must open to.
type Vector struct {
	// Name identifies the case in test output.
	Name string
	// RequestID is authenticated data: opening with any other value fails.
	RequestID string
	// Sealed is the base64url blob as it appears in the `credentials` query
	// parameter of the loopback redirect.
	Sealed string
	// Plaintext is what Sealed must decrypt to.
	Plaintext string
}

// Vectors are stable by contract. Adding is fine; changing one means the wire
// format changed, which is a new version marker and a new rollout - see
// docs/architecture/oauth2-callback-credential.md.
var Vectors = []Vector{
	{
		Name:      "callback payload",
		RequestID: "0f4b9a2e-7c31-4f5a-9d0e-2b7a1c6e5d43",
		Sealed: "R1RFMQ4HzlXtsTJjfGt6oM74XFOCCmx96sNLT-Edp2XVEjVuZPTwn3eeTwsC_svteZ7RLLPWtcgoUAzdiKFm" +
			"u_AQkzTy45ovMr4X6LxOxa1oWIG8R2dU4gdmO4HRm1xJHEvK9NIPYIpP6UFFuKeV-rW9Vwo2kGpcEqiFI35t" +
			"yPVOndfwjyKHuUoadeHp6YU1rB4XW6Pr3wi6BO2H",
		Plaintext: "reqId=0f4b9a2e-7c31-4f5a-9d0e-2b7a1c6e5d43&token=ya29.test-access-token&timestamp=1700000000",
	},
	{
		// An empty payload is representable and must round trip; it is also
		// the shortest legal blob, which is where a length check gets it
		// wrong.
		Name:      "empty payload",
		RequestID: "a-second-login",
		Sealed:    "R1RFMY7-nryqlufPdIvnRHjoFe_EKVOF2CdMIxpp3tb3DodkS8ezNJcc4ayjB2cLBzO6K61rP-Xm8EK6VWZPmQ==",
		Plaintext: "",
	},
}
