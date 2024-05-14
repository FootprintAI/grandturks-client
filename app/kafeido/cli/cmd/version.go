package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	"github.com/footprintai/grandturks-client/v2/pkg/version"
)

func NewVersionCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "version",
		Short: "version of kafeido cli",
		RunE: func(cmd *cobra.Command, args []string) error {
			version.Print()
			return nil
		},
	}
	return cmd
}
