package cmd

import (
	"errors"
	"fmt"
	"net"
	"net/url"
	"strings"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"
)

func NewConfigSetCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "set",
		Short: "",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
	}

	cmd.AddCommand(NewConfigSetEndpointCommand(logger, ioStreams))
	cmd.AddCommand(NewConfigSetTimeoutCommand(logger, ioStreams))
	return cmd
}

func NewConfigSetEndpointCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		apiEndpoint     string
		storageEndpoint string
	)

	var handler = func() error {
		if err := isValidEndpoint(apiEndpoint); err != nil {
			return err
		}
		if err := isValidEndpoint(storageEndpoint); err != nil {
			return err
		}
		ConfigKeyApiEndpoint.Set(apiEndpoint)
		ConfigKeyStorageEndpoint.Set(storageEndpoint)
		return viper.WriteConfig()
	}

	cmd := &cobra.Command{
		Use:   "endpoint",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&apiEndpoint, "api_endpoint", "", "api endpoint (default: ''")
	cmd.Flags().StringVar(&storageEndpoint, "storage_endpoint", "", "storage endpoint (default: ''")
	cmd.MarkFlagRequired("api_endpoint")
	cmd.MarkFlagRequired("storage_endpoint")
	return cmd
}

// isValidEndpoint refuses any endpoint that would send the CLI's bearer token
// over plaintext to somewhere other than this machine.
//
// https is required for remote hosts, and that is the point of this check: the
// CLI attaches a bearer token to every call, so an http endpoint pointed at a
// remote host leaks credentials to anything on the path.
//
// http to LOOPBACK is allowed, because it is not that risk and forbidding it
// made the CLI unusable against the project's own stacks - the sandbox binds
// SeaweedFS S3 on 127.0.0.1:8333 and the demo webportal on 127.0.0.1:3000, both
// plain HTTP. With a scheme-only check there was no way to point the CLI at
// either through `config set endpoint`; you had to hand-write
// ~/.kafeidoconfig.json or set env vars whose names contain a literal dot
// (grandturks#901).
func isValidEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}

	switch strings.ToLower(u.Scheme) {
	case "https":
		return nil
	case "http":
		if isLoopbackHost(u.Hostname()) {
			return nil
		}
		return fmt.Errorf(
			"require https for %q: plaintext http is allowed only for loopback "+
				"(127.0.0.0/8, ::1, localhost), because the CLI sends a bearer token",
			u.Host)
	default:
		return errors.New("require https")
	}
}

// isLoopbackHost reports whether a URL host refers to this machine.
//
// The host is parsed and compared, never matched by prefix or substring:
// `127.0.0.1.evil.com` and `localhost.evil.com` are ordinary remote names that
// a naive check would wave through, handing the bearer token to whoever owns
// them. u.Hostname() is what strips the port and the IPv6 brackets, so `::1`
// still parses as an address here.
func isLoopbackHost(host string) bool {
	if strings.EqualFold(host, "localhost") {
		return true
	}
	if ip := net.ParseIP(host); ip != nil {
		return ip.IsLoopback()
	}
	return false
}

func NewConfigGetCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get",
		Short: "",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
	}
	cmd.AddCommand(NewConfigGetEndpointCommand(logger, ioStreams))
	return cmd

}

func NewConfigGetEndpointCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {

	var handler = func() error {
		fmt.Printf("endpoint.api: %s\n", ConfigKeyApiEndpoint.GetString())
		fmt.Printf("endpoint.storage: %s\n", ConfigKeyStorageEndpoint.GetString())
		return nil
	}

	cmd := &cobra.Command{
		Use:   "endpoint",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}
	return cmd

}

func NewConfigSetTimeoutCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		requestTimeout string
	)

	var handler = func() error {
		dur, err := time.ParseDuration(requestTimeout)
		if err != nil {
			return err
		}
		ConfigKeyRequestTimeout.Set(dur)
		return viper.WriteConfig()
	}

	cmd := &cobra.Command{
		Use:   "timeout",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&requestTimeout, "request", "", "timeout in sec(default: 45s")
	cmd.MarkFlagRequired("request")
	return cmd
}
