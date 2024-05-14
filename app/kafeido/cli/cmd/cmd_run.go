package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"
)

func NewRunCommand(logger log.Logger, ioStreams genericclioptions.IOStreams, newCmdFuncs ...NewCommandFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "run",
		Short: "run command",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
			fmt.Println("in run pre run")
		},
	}

	cmd.AddCommand(NewRunPipelineCommand(logger, ioStreams))
	for _, newCmdFunc := range newCmdFuncs {
		cmd.AddCommand(newCmdFunc(logger, ioStreams))
	}

	return cmd
}
