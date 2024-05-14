package cmd

import (
	"errors"
	"fmt"
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

func isValidEndpoint(endpoint string) error {
	u, err := url.Parse(endpoint)
	if err != nil {
		return err
	}
	if strings.ToLower(u.Scheme) != "https" {
		return errors.New("require https")
	}
	return nil
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
