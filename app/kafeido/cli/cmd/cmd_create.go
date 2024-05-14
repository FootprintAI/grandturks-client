package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"
)

func NewCreateCommand(logger log.Logger, ioStreams genericclioptions.IOStreams, newCmdFuncs ...NewCommandFunc) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "create",
		Short: "create command",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
	}

	cmd.AddCommand(NewCreateProjectCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateInferenceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateInferenceJobCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreatePipelineCommand(logger, ioStreams))
	cmd.AddCommand(NewCreatePredictionCommand(logger, ioStreams))
	for _, newCmdFunc := range newCmdFuncs {
		cmd.AddCommand(newCmdFunc(logger, ioStreams))
	}

	return cmd
}
