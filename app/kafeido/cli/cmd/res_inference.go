package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
)

func NewCreateInferenceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId       string
		modelName       string
		modelUri        string
		runtimeVersion  string
		namedPipelineId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceCreateModelInferenceParams{
			Body: appservice.KafeidoServiceCreateModelInferenceBody{
				ModelName:       modelName,
				NamedPipelineID: namedPipelineId,
				ModelURI:        modelUri,
				RuntimeVersion:  runtimeVersion,
			},
			ProjectID: projectId,
		}
		kafeidoCreateModelInferenceOk, err := runCmd.stub.KafeidoService.KafeidoServiceCreateModelInference(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(kafeidoCreateModelInferenceOk.Payload.ModelInferenceID, "Model Inference is created").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "inference",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelName, "model_name", "", "unqiue model name, leave it blank will use random generated modelname")
	cmd.Flags().StringVar(&modelUri, "model_uri", "", "model uri, like s3:bucketname/objectname or https://... ")
	cmd.Flags().StringVar(&runtimeVersion, "runtime_version", "", "runtime version")
	cmd.Flags().StringVar(&namedPipelineId, "named_pipeline_id", "", "named pipeine id for run")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("named_pipeline_id")

	return cmd
}

func NewGetInferenceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId        string
		modelInferenceId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceGetModelInferenceParams{
			ProjectID:        projectId,
			ModelInferenceID: modelInferenceId,
		}
		kafeidoGetModelInferenceOk, err := runCmd.stub.KafeidoService.KafeidoServiceGetModelInference(
			params.WithRequestTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewModelInferenceDetailsFormatter([]*appmodels.KafeidoGetModelInferenceResponse{kafeidoGetModelInferenceOk.Payload}).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "inference",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelInferenceId, "inference_id", "", "model inference id")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("inference_id")

	return cmd
}

func NewListInferenceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceListModelInferenceParams{
			ProjectID: projectId,
		}
		kafeidoListModelInferenceOk, err := runCmd.stub.KafeidoService.KafeidoServiceListModelInference(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		var listResp []*appmodels.KafeidoGetModelInferenceResponse
		for _, modelInferenceId := range kafeidoListModelInferenceOk.Payload.ModelInferenceID {
			subParams := &appservice.KafeidoServiceGetModelInferenceParams{
				ProjectID:        projectId,
				ModelInferenceID: modelInferenceId,
			}
			kafeidoGetModelInferenceOk, err := runCmd.stub.KafeidoService.KafeidoServiceGetModelInference(
				subParams.WithRequestTimeout(runCmd.requestTimeout),
				runCmd.authInformer(),
			)

			if err == nil {
				listResp = append(listResp, kafeidoGetModelInferenceOk.Payload)
			}
		}
		return format.NewModelInferenceDetailsFormatter(listResp).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "inference",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.MarkFlagRequired("project_id")

	return cmd
}

func NewDeleteInferenceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId        string
		modelInferenceId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceDeleteModelInferenceParams{
			ProjectID:        projectId,
			ModelInferenceID: modelInferenceId,
		}
		_, err := runCmd.stub.KafeidoService.KafeidoServiceDeleteModelInference(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(modelInferenceId, "Inference has been deleted").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "inference",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelInferenceId, "inference_id", "", "model inference id")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("inference_id")

	return cmd
}
