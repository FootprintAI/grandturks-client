package cmd

import (
	"time"

	"github.com/spf13/viper"
)

type viperConfigKey string

func (t viperConfigKey) String() string {
	return string(t)
}

func (t viperConfigKey) GetString() string {
	val := viper.Get(t.String())
	_, isStr := val.(string)
	if !isStr {
		return ""
	}
	return val.(string)
}

func (t viperConfigKey) GetDuration() time.Duration {
	val := viper.Get(t.String())
	_, isDur := val.(time.Duration)
	if !isDur {
		return time.Duration(0)
	}
	return val.(time.Duration)
}

func (t viperConfigKey) Get() interface{} {
	return viper.Get(t.String())
}

func (t viperConfigKey) Set(val interface{}) {
	viper.Set(t.String(), val)
}

func (t viperConfigKey) setDefault(val interface{}) {
	viper.SetDefault(t.String(), val)
}

var (
	ConfigKeyApiEndpoint     viperConfigKey = "endpoint.api"
	ConfigKeyStorageEndpoint viperConfigKey = "endpoint.storage"
	ConfigKeyUserId          viperConfigKey = "userInfo.userId"
	ConfigKeyUserGroups      viperConfigKey = "userInfo.groups"
	ConfigKeyUserEmail       viperConfigKey = "userInfo.email"
	ConfigKeyAuthToken       viperConfigKey = "authToken"
	ConfigKeyRequestTimeout  viperConfigKey = "requestTimeout"
	ConfigKeyApiKey          viperConfigKey = "apiKey"
)

// apiKeyEnvVar is how an integration supplies its credential.
//
// An api key exists so a machine can call the API without a human login
// (#21), and a machine has no interactive step in which to write
// ~/.kafeidoconfig.
//
// Read with os.Getenv in resolveAPIKey and deliberately NOT bound with
// viper.BindEnv. initConfig calls viper.SafeWriteConfig() when no config file
// exists, which writes every setting viper knows about - so a bound key meant
// that one `KAFEIDO_API_KEY=gtk_... kafeido list project` on a fresh machine
// wrote the credential to ~/.kafeidoconfig.json in plaintext. A key passed
// through the environment is deliberately not on disk, and persisting it is
// not this CLI's decision to make.
const apiKeyEnvVar = "KAFEIDO_API_KEY"

func init() {
	viper.SetTypeByDefaultValue(true)
	ConfigKeyRequestTimeout.setDefault(45 * time.Second)
}
