package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"
)

func NewCancelCommand(logger log.Logger, ioStreams genericclioptions.IOStreams, newCmdFuncs ...NewCommandFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "cancel",
		Short: "cancel command",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
	}

	cmd.AddCommand(NewCancelInferenceJobCommand(logger, ioStreams))
	for _, newCmdFunc := range newCmdFuncs {
		cmd.AddCommand(newCmdFunc(logger, ioStreams))
	}

	return cmd
}
