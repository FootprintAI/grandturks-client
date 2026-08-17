package cmd

import (
	"fmt"
	"os"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/footprintai/grandturks-client/v2/pkg/apikey"
)

// apiKeyHeader is the header an api key travels in. See pkg/apikey for why it
// is not Authorization; the short version is that an ingress rejects anything
// in Authorization that will not parse as a JWT, and a gtk_ token will not.
const apiKeyHeader = apikey.HTTPHeader

// apiKeyFlag is the value of the root command's --api_key flag.
var apiKeyFlag string

// resolveAPIKey returns the api key to authenticate with, or "".
//
// Precedence is explicit-beats-stored: the flag, then KAFEIDO_API_KEY, then
// ~/.kafeidoconfig. An operator who just typed a key on the command line must
// not be outranked by whatever a previous session left in the config file.
//
// The environment is read here rather than through viper.BindEnv on purpose -
// see apiKeyEnvVar in config.go: a bound value ends up in the config file
// viper writes on first run, which puts a credential on disk that its holder
// deliberately kept in the environment.
func resolveAPIKey() string {
	if len(apiKeyFlag) > 0 {
		return apiKeyFlag
	}
	if fromEnv := os.Getenv(apiKeyEnvVar); len(fromEnv) > 0 {
		return fromEnv
	}
	return ConfigKeyApiKey.GetString()
}

// newAuthInformer picks the ONE credential a request carries.
//
// One, not both, and this is the part worth reading twice. The server's
// PromoteToAuthorization moves an api key from X-Api-Key into the bearer
// credential only when there is no Authorization header, and lets
// Authorization win when both are present - reasonably, since resolving it the
// other way would let a weaker credential replace a stronger one on a request
// that already authenticated. So a CLI that sent a saved JWT alongside an
// explicitly-supplied key would run as the human and silently ignore the key.
//
// A malformed key is refused here rather than sent. This is also the only
// place it can be refused politely: every command builds its client through
// mustNewRunCmd, which panics on error, whereas an error from an auth writer
// comes back from the operation call as an ordinary command error.
func newAuthInformer(apiKeyToken, authToken string) runtime.ClientAuthInfoWriter {
	return runtime.ClientAuthInfoWriterFunc(func(clientRequest runtime.ClientRequest, registry strfmt.Registry) error {
		if len(apiKeyToken) > 0 {
			if _, err := apikey.Parse(apiKeyToken); err != nil {
				return fmt.Errorf("api key is not usable: %w - expected the form gtk_<keyId>_<secret>, "+
					"as returned by `kafeido create apikey`", err)
			}
			return clientRequest.SetHeaderParam(apiKeyHeader, apiKeyToken)
		}
		if len(authToken) > 0 {
			return clientRequest.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", authToken))
		}
		return nil
	})
}
