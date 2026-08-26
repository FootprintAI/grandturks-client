package cmd

// FootprintAI/manifests#308.
//
// The API's path prefix was the literal "api", compiled into root.go. That is
// not a preference - it decides whether a deployment can move the product API
// off a prefix Kubeflow also owns, and it could not be discussed while every
// binary in the field had it baked in with no way to know which versions were
// deployed. FootprintAI/grandturks#1251 hit the same wall with its AES
// constants and had to keep them forever.
//
// The tests that matter here are the ones asserting NOTHING CHANGES by
// default. A migration aid that quietly moves existing installs is worse than
// no migration aid.

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func resetBasePath(t *testing.T) {
	t.Helper()
	viper.Reset()
	viper.SetTypeByDefaultValue(true)
	ConfigKeyApiBasePath.setDefault(DefaultApiBasePath)
}

// The whole point: an install that says nothing keeps talking to /api.
func TestBasePathDefaultsToApiSoNothingInTheFieldMoves(t *testing.T) {
	resetBasePath(t)
	t.Setenv(apiBasePathEnvVar, "")
	// t.Setenv sets it to empty, which is a DELIBERATE value - see the empty
	// case below. Unset it so this test sees a machine with no override.
	assert.NoError(t, osUnsetenv(apiBasePathEnvVar))

	assert.Equal(t, "api", resolveApiBasePath(),
		"a config that names no prefix must keep using the one every deployment serves today")
}

func TestBasePathPrecedence(t *testing.T) {
	for _, tc := range []struct {
		name   string
		config string
		env    *string
		want   string
		why    string
	}{
		{
			name: "config only", config: "gateway", want: "gateway",
			why: "a deployment that has moved its API says so once, in its config file",
		},
		{
			name: "env beats config", config: "gateway", env: strp("other"), want: "other",
			why: "trying a new prefix for one command is the point of the variable; it must win",
		},
		{
			name: "env only", env: strp("gateway"), want: "gateway",
			why: "an ephemeral caller - CI, a container - has no config file to edit",
		},
		{
			name: "env explicitly empty", env: strp(""), want: "",
			why: "a deployment serving the API at the host root has no prefix, and \"\" is how it says so - " +
				"this is why the lookup tests presence, not truthiness",
		},
		{
			name: "config empty falls back", config: "", want: "api",
			why: "an absent key and an empty key are the same thing in a config file, unlike an env var",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			resetBasePath(t)
			assert.NoError(t, osUnsetenv(apiBasePathEnvVar))
			if tc.config != "" {
				ConfigKeyApiBasePath.Set(tc.config)
			}
			if tc.env != nil {
				t.Setenv(apiBasePathEnvVar, *tc.env)
			}
			assert.Equal(t, tc.want, resolveApiBasePath(), tc.why)
		})
	}
}

// "api", "/api" and "api/" are the same prefix and must not produce three
// different URLs. filepath.Join in root.go tolerates some of this; relying on
// that would make the behaviour an accident of the joining function.
func TestBasePathIsNormalised(t *testing.T) {
	for _, in := range []string{"gateway", "/gateway", "gateway/", "/gateway/"} {
		t.Run(in, func(t *testing.T) {
			resetBasePath(t)
			assert.NoError(t, osUnsetenv(apiBasePathEnvVar))
			t.Setenv(apiBasePathEnvVar, in)
			assert.Equal(t, "gateway", resolveApiBasePath())
		})
	}
}

func strp(s string) *string { return &s }
