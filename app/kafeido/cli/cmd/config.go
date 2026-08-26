package cmd

import (
	"os"
	"strings"
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
	// ConfigKeyApiBasePath is the path prefix the API is served under, between
	// the host and the /v1/ routes: https://<host>/<basePath>/v1/projects.
	//
	// It was the literal "api", compiled into root.go. That is a problem the
	// server side cannot solve alone: `/api` is ALSO the Kubeflow central
	// dashboard's own API prefix, and on a single-host deployment the two
	// collide (FootprintAI/manifests#308 - the dashboard's requests were
	// answered by appkafeido and it rendered "No Namespaces").
	//
	// Moving the product off /api is the real fix, and it could not even be
	// discussed while every distributed binary had the prefix baked in with no
	// way to know which versions were deployed. Same wall
	// FootprintAI/grandturks#1251 hit with its AES constants.
	//
	// DEFAULT UNCHANGED. Every existing config file, every scripted install and
	// every binary already in the field keeps talking to /api. This makes the
	// move possible later; it does not make it.
	ConfigKeyApiBasePath viperConfigKey = "endpoint.apiBasePath"
)

// DefaultApiBasePath is what the API has always been served under. Named
// rather than repeated so the day it changes is one edit and one grep.
const DefaultApiBasePath = "api"

// apiBasePathEnvVar lets a caller override the prefix for one invocation.
//
// Read with os.Getenv rather than viper.BindEnv, for the reason apiKeyEnvVar
// documents above: initConfig calls viper.SafeWriteConfig() on a fresh
// machine, which persists everything viper knows about. A value supplied for
// one command should not silently become that machine's stored setting - and
// during a migration the whole point is to try the new prefix WITHOUT
// committing to it.
const apiBasePathEnvVar = "KAFEIDO_API_BASE_PATH"

// resolveApiBasePath returns the prefix to sit between host and routes, in
// precedence order: environment, config file, default.
//
// Normalised because the three spellings a person will type - "api", "/api",
// "api/" - must not produce three different URLs. An explicitly empty value is
// honoured as "no prefix", which is what a deployment serving the API at the
// host root needs; that is why this cannot simply be `if s == "" { default }`
// on the environment variable.
func resolveApiBasePath() string {
	if v, ok := os.LookupEnv(apiBasePathEnvVar); ok {
		return strings.Trim(v, "/")
	}
	if v := ConfigKeyApiBasePath.GetString(); v != "" {
		return strings.Trim(v, "/")
	}
	return DefaultApiBasePath
}

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
	ConfigKeyApiBasePath.setDefault(DefaultApiBasePath)
}
