package cmd

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
)

// convertShortRoleToFull maps the role a person types to the enum the contract
// takes, in the style of convertShortVisibilityToFull.
//
// ADMIN is a value of the enum and is not offered here: the server refuses it,
// because it has no project tier distinct from READWRITE in this authz model.
// An operator who asked for an "admin" key and got one would believe they had
// granted something they had not, so the refusal is spelled out rather than
// left to a 400.
func convertShortRoleToFull(s string) (appmodels.KafeidocommonpbAPIKeyRole, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "read":
		return appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEREAD, nil
	case "readwrite":
		return appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEREADWRITE, nil
	case "admin":
		return appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEUNSPECIFIED, errors.New(
			"role admin is not available for api keys - it has no project tier distinct " +
				"from readwrite, so use readwrite")
	default:
		return appmodels.KafeidocommonpbAPIKeyRoleAPIKEYROLEUNSPECIFIED, fmt.Errorf(
			"unknown role %q, one of read, readwrite", s)
	}
}

// newCreateAPIKeyBody validates a create request before a credential exists.
//
// The lifetime rules are the server's (userStore resolveExpiry) and are
// repeated here on purpose: a request that carries both a ttl and
// --never_expires is a confused caller, and the cheapest place to say so is
// before a key is minted rather than after.
func newCreateAPIKeyBody(name, shortRole string, ttl time.Duration, neverExpires bool) (appservice.KafeidoServiceCreateAPIKeyBody, error) {
	var body appservice.KafeidoServiceCreateAPIKeyBody

	if len(strings.TrimSpace(name)) == 0 {
		return body, errors.New("a name is required - it is what identifies the party holding the key, e.g. acme-ci")
	}
	role, err := convertShortRoleToFull(shortRole)
	if err != nil {
		return body, err
	}
	if neverExpires && ttl != 0 {
		return body, errors.New("--ttl and --never_expires are mutually exclusive")
	}

	body.Name = name
	body.Role = role.Pointer()
	body.NeverExpires = neverExpires

	if !neverExpires && ttl != 0 {
		// Seconds is the wire unit; a duration is what a person can write.
		// Anything under a second truncates to zero, which the wire cannot
		// distinguish from "unset", so it is refused rather than silently
		// turned into the server's default lifetime.
		seconds := int64(ttl / time.Second)
		if seconds <= 0 {
			return body, fmt.Errorf("--ttl must be positive and at least one second, got %s", ttl)
		}
		body.TTLSeconds = strconv.FormatInt(seconds, 10)
	}
	return body, nil
}

// NewCreateApiKeyCommand mints a key for an external party.
func NewCreateApiKeyCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId    string
		keyName      string
		role         string
		ttl          time.Duration
		neverExpires bool
	)

	handler := func() error {
		body, err := newCreateAPIKeyBody(keyName, role, ttl, neverExpires)
		if err != nil {
			return err
		}

		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceCreateAPIKeyParams{
			ProjectID: projectId,
			Body:      body,
		}
		createApiKeyOk, err := runCmd.stub.KafeidoService.KafeidoServiceCreateAPIKey(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewCreatedApiKeyFormatter(
			createApiKeyOk.Payload.Key,
			createApiKeyOk.Payload.Token,
		).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "create an api key an external party can call the api with",
		Long: "Create an api key for machine-to-machine access to a project.\n\n" +
			"The token is returned exactly once, here, and is never recoverable afterwards - " +
			"the listing endpoint deliberately cannot return it. Use it with --api_key, or by " +
			"setting " + apiKeyEnvVar + ".",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project the key is scoped to")
	cmd.Flags().StringVar(&keyName, "name", "", "name identifying the party holding the key, e.g. acme-ci")
	// Least privilege by default: a key that can only read is the safer thing
	// to hand out by accident.
	cmd.Flags().StringVar(&role, "role", "read", "one of read, readwrite")
	cmd.Flags().DurationVar(&ttl, "ttl", 0, "lifetime, e.g. 720h (default: the server's default lifetime)")
	cmd.Flags().BoolVar(&neverExpires, "never_expires", false, "mint a key with no expiry (mutually exclusive with --ttl)")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("name")

	return cmd
}

// NewListApiKeyCommand lists a project's keys. Never their secrets - the
// server cannot return those, by design.
func NewListApiKeyCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var projectId string

	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceListAPIKeysParams{
			ProjectID: projectId,
		}
		listApiKeysOk, err := runCmd.stub.KafeidoService.KafeidoServiceListAPIKeys(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewListApiKeyDetailsFormatter(listApiKeysOk.Payload.Keys).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "list a project's api keys",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project whose keys to list")
	cmd.MarkFlagRequired("project_id")

	return cmd
}

// NewDeleteApiKeyCommand revokes a key. Named delete to match the CLI's verbs;
// the endpoint is a revoke, and the distinction matters - the record survives
// with a revokedAt so an audit trail of what the key wrote stays readable.
func NewDeleteApiKeyCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
		keyId     string
	)

	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceRevokeAPIKeyParams{
			ProjectID: projectId,
			KeyID:     keyId,
		}
		revokeApiKeyOk, err := runCmd.stub.KafeidoService.KafeidoServiceRevokeAPIKey(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		revokedId := keyId
		if revokeApiKeyOk.Payload.Key != nil && len(revokeApiKeyOk.Payload.Key.KeyID) > 0 {
			revokedId = revokeApiKeyOk.Payload.Key.KeyID
		}
		return format.NewIdAndMessageFormatter(revokedId, "Api key is revoked").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "apikey",
		Short: "revoke an api key",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project the key belongs to")
	cmd.Flags().StringVar(&keyId, "key_id", "", "id of the key to revoke, as shown by `list apikey`")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("key_id")

	return cmd
}
