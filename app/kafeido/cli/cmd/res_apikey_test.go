package cmd

import (
	"bytes"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	kindcmd "sigs.k8s.io/kind/pkg/cmd"

	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

// TestConvertShortRoleToFull maps what a person types to what the contract
// takes. ADMIN is deliberately absent: the server refuses it, because it has
// no project tier distinct from READWRITE in this authz model, and an operator
// handing an outside party an "admin" key would believe they granted something
// they did not. Refusing it locally means that surprise arrives as a CLI
// message rather than as a 400 after the fact.
func TestConvertShortRoleToFull(t *testing.T) {
	for _, tc := range []struct {
		name    string
		short   string
		want    appmodels.KafeidocommonpbAPIKeyRole
		wantErr bool
	}{
		{name: "read", short: "read", want: appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEREAD},
		{name: "readwrite", short: "readwrite", want: appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEREADWRITE},
		{name: "case insensitive", short: "ReadWrite", want: appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEREADWRITE},
		{name: "admin is refused", short: "admin", wantErr: true},
		{name: "unknown", short: "superuser", wantErr: true},
		{name: "empty", short: "", wantErr: true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, err := convertShortRoleToFull(tc.short)
			if tc.wantErr {
				if err == nil {
					t.Fatalf("convertShortRoleToFull(%q) = %v, want an error", tc.short, got)
				}
				return
			}
			if err != nil {
				t.Fatalf("convertShortRoleToFull(%q): %v", tc.short, err)
			}
			if got != tc.want {
				t.Errorf("convertShortRoleToFull(%q) = %q, want %q", tc.short, got, tc.want)
			}
		})
	}
}

// TestConvertShortRoleToFullAdminMessage: "admin is not a valid role" would
// send someone looking for a typo. The message has to say the role exists and
// is refused, and what to use instead.
func TestConvertShortRoleToFullAdminMessage(t *testing.T) {
	_, err := convertShortRoleToFull("admin")
	if err == nil {
		t.Fatal("admin was accepted")
	}
	if !strings.Contains(err.Error(), "readwrite") {
		t.Errorf("error %q does not point at readwrite", err)
	}
}

// TestNewCreateAPIKeyBody covers the lifetime rules, which the server also
// enforces (userStore resolveExpiry). Enforcing them here too means a confused
// request fails before a credential is minted, not after.
func TestNewCreateAPIKeyBody(t *testing.T) {
	for _, tc := range []struct {
		name         string
		keyName      string
		role         string
		ttl          time.Duration
		neverExpires bool
		wantTTL      string
		wantNever    bool
		wantErr      string
	}{
		{
			name:    "default lifetime is the server's",
			keyName: "acme-ci",
			role:    "read",
			wantTTL: "", // omitted, so the server applies its own default
		},
		{
			name:    "explicit ttl in seconds",
			keyName: "acme-ci",
			role:    "readwrite",
			ttl:     24 * time.Hour,
			wantTTL: "86400",
		},
		{
			name:         "never expires",
			keyName:      "acme-ci",
			role:         "read",
			neverExpires: true,
			wantNever:    true,
		},
		{
			// Refused rather than resolved in someone's favour - a request
			// carrying both is a confused caller, and guessing is how a
			// credential ends up with a lifetime nobody chose.
			name:         "ttl and never_expires together",
			keyName:      "acme-ci",
			role:         "read",
			ttl:          time.Hour,
			neverExpires: true,
			wantErr:      "mutually exclusive",
		},
		{
			name:    "negative ttl",
			keyName: "acme-ci",
			role:    "read",
			ttl:     -time.Hour,
			wantErr: "positive",
		},
		{
			name:    "sub-second ttl rounds to zero and is refused",
			keyName: "acme-ci",
			role:    "read",
			ttl:     500 * time.Millisecond,
			wantErr: "positive",
		},
		{
			name:    "empty name",
			role:    "read",
			wantErr: "name",
		},
		{
			name:    "bad role",
			keyName: "acme-ci",
			role:    "admin",
			wantErr: "admin",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			body, err := newCreateAPIKeyBody(tc.keyName, tc.role, tc.ttl, tc.neverExpires)
			if tc.wantErr != "" {
				if err == nil {
					t.Fatalf("newCreateAPIKeyBody(...) = %+v, want an error containing %q", body, tc.wantErr)
				}
				if !strings.Contains(strings.ToLower(err.Error()), tc.wantErr) {
					t.Errorf("error %q does not contain %q", err, tc.wantErr)
				}
				return
			}
			if err != nil {
				t.Fatalf("newCreateAPIKeyBody(...): %v", err)
			}
			if body.Name != tc.keyName {
				t.Errorf("Name = %q, want %q", body.Name, tc.keyName)
			}
			if body.TTLSeconds != tc.wantTTL {
				t.Errorf("TTLSeconds = %q, want %q", body.TTLSeconds, tc.wantTTL)
			}
			if body.NeverExpires != tc.wantNever {
				t.Errorf("NeverExpires = %v, want %v", body.NeverExpires, tc.wantNever)
			}
			if body.Role == nil {
				t.Fatal("Role is nil")
			}
		})
	}
}

// TestAPIKeyCommandsAreReachable walks the command tree the way a user does.
// The capability #21 reports missing is not "the client can call the endpoint"
// but "the CLI can": a handler nobody wired into create/list/delete is the
// same as no handler.
func TestAPIKeyCommandsAreReachable(t *testing.T) {
	for _, tc := range []struct {
		verb          string
		requiredFlags []string
	}{
		{verb: "create", requiredFlags: []string{"project_id", "name"}},
		{verb: "list", requiredFlags: []string{"project_id"}},
		{verb: "delete", requiredFlags: []string{"project_id", "key_id"}},
	} {
		t.Run(tc.verb, func(t *testing.T) {
			root := newTestRootCommand()

			cmd, _, err := root.Find([]string{tc.verb, "apikey"})
			if err != nil {
				t.Fatalf("Find(%q apikey): %v", tc.verb, err)
			}
			if cmd.Name() != "apikey" {
				t.Fatalf("%s resolved to %q, want the apikey subcommand", tc.verb, cmd.Name())
			}
			for _, flag := range tc.requiredFlags {
				if cmd.Flags().Lookup(flag) == nil {
					t.Errorf("%s apikey has no --%s flag", tc.verb, flag)
				}
			}
		})
	}
}

// TestAPIKeyFlagIsGlobal: authenticating with a key has to work on every
// command, not just the apikey ones - that is the row in #21's table that
// matters most, since it is what lets the CLI run as an integration at all.
func TestAPIKeyFlagIsGlobal(t *testing.T) {
	root := newTestRootCommand()

	if root.PersistentFlags().Lookup("api_key") == nil {
		t.Fatal("--api_key is not a persistent flag on the root command")
	}

	cmd, _, err := root.Find([]string{"list", "project"})
	if err != nil {
		t.Fatalf("Find(list project): %v", err)
	}
	if cmd.Flags().Lookup("api_key") == nil && cmd.InheritedFlags().Lookup("api_key") == nil {
		t.Error("list project cannot be run with --api_key")
	}
}

// TestResolveAPIKeyPrecedence: an integration sets the environment variable or
// the flag; the config file is what a login writes. The explicit signal wins,
// or a stale ~/.kafeidoconfig silently outranks what the operator just typed.
func TestResolveAPIKeyPrecedence(t *testing.T) {
	const (
		fromFlag   = "gtk_00000000000000ff_ZmxhZw"
		fromConfig = "gtk_00000000000000aa_Y29uZmln"
	)

	for _, tc := range []struct {
		name   string
		flag   string
		config string
		want   string
	}{
		{name: "flag wins over config", flag: fromFlag, config: fromConfig, want: fromFlag},
		{name: "config when no flag", config: fromConfig, want: fromConfig},
		{name: "neither", want: ""},
	} {
		t.Run(tc.name, func(t *testing.T) {
			t.Cleanup(func() {
				apiKeyFlag = ""
				ConfigKeyApiKey.Set("")
			})
			apiKeyFlag = tc.flag
			ConfigKeyApiKey.Set(tc.config)

			if got := resolveAPIKey(); got != tc.want {
				t.Errorf("resolveAPIKey() = %q, want %q", got, tc.want)
			}
		})
	}
}

// newTestRootCommand builds the command tree exactly as main.go does, so these
// tests walk the same tree a user does.
func newTestRootCommand() *cobra.Command {
	var out, errOut bytes.Buffer
	return NewRootCommand(
		kindcmd.NewLogger(),
		genericclioptions.IOStreams{In: strings.NewReader(""), Out: &out, ErrOut: &errOut},
	)
}

// TestAPIKeyFromTheEnvironmentIsNotPersisted guards a credential leak this
// work introduced and nearly shipped.
//
// initConfig calls viper.SafeWriteConfig() whenever no config file exists yet,
// and that writes every setting viper knows about. Binding KAFEIDO_API_KEY
// with viper.BindEnv made the key one of them, so a single
//
//	KAFEIDO_API_KEY=gtk_... kafeido list project
//
// on a fresh machine wrote the credential to ~/.kafeidoconfig.json in
// plaintext - observed, not theorised. An api key handed to a process through
// the environment is deliberately not on disk, and it is not the CLI's call to
// put it there.
//
// So the environment is read directly and viper is never told about it.
func TestAPIKeyFromTheEnvironmentIsNotPersisted(t *testing.T) {
	const envKey = "gtk_00000000000000bb_ZW52"
	t.Setenv(apiKeyEnvVar, envKey)
	t.Cleanup(func() { apiKeyFlag = "" })

	if got := resolveAPIKey(); got != envKey {
		t.Errorf("resolveAPIKey() = %q, want the value from $%s", got, apiKeyEnvVar)
	}
	for name, value := range viper.AllSettings() {
		if str, ok := value.(string); ok && str == envKey {
			t.Errorf("viper setting %q holds the api key from the environment - "+
				"SafeWriteConfig would write it to disk", name)
		}
	}
}

// TestAPIKeyPrecedenceFlagOverEnvironment: the flag is the most explicit
// signal there is, so it outranks an environment variable inherited from a
// shell or a CI runner.
func TestAPIKeyPrecedenceFlagOverEnvironment(t *testing.T) {
	const (
		envKey  = "gtk_00000000000000bb_ZW52"
		flagKey = "gtk_00000000000000cc_ZmxhZw"
	)
	t.Setenv(apiKeyEnvVar, envKey)
	t.Cleanup(func() { apiKeyFlag = "" })
	apiKeyFlag = flagKey

	if got := resolveAPIKey(); got != flagKey {
		t.Errorf("resolveAPIKey() = %q, want the flag value %q", got, flagKey)
	}
}
