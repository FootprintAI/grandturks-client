package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"
)

func NewListCommand(logger log.Logger, ioStreams genericclioptions.IOStreams, newCmdFuncs ...NewCommandFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "listcommand",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
	}

	cmd.AddCommand(NewListProjectCommand(logger, ioStreams))
	cmd.AddCommand(NewListInferenceCommand(logger, ioStreams))
	cmd.AddCommand(NewListInferenceJobCommand(logger, ioStreams))
	cmd.AddCommand(NewListDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewListPipelineCommand(logger, ioStreams))
	//cmd.AddCommand(NewListStorageCommand(logger, ioStreams))
	for _, newCmdFunc := range newCmdFuncs {
		cmd.AddCommand(newCmdFunc(logger, ioStreams))
	}

	return cmd
}
