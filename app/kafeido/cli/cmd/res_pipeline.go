package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/go-openapi/strfmt"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
)

func NewCreatePipelineCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId    string
		pipelineName string
		pipelineDesc string
		manifestPath string // workflow manifest path
	)
	handler := func() error {
		manifestBytes, err := ioutil.ReadFile(manifestPath)
		if err != nil {
			return fmt.Errorf("failed to read file, err:%+v\n", err)
		}
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoCreateProjectPipelineParams{
			Body: appservice.KafeidoCreateProjectPipelineBody{
				Name:               pipelineName,
				Desc:               pipelineDesc,
				KfPipelineWorkflow: strfmt.Base64(manifestBytes),
			},
			ProjectID: projectId,
		}

		kafeidoCreateProjectPipelineOk, err := runCmd.stub.Kafeido.KafeidoCreateProjectPipeline(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(kafeidoCreateProjectPipelineOk.Payload.NamedPipelineID, "pipeline created").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&pipelineName, "name", "", "pipeline name")
	cmd.MarkFlagRequired("name")
	cmd.Flags().StringVar(&manifestPath, "manifest_path", "", "manifeset manifest local path")
	cmd.MarkFlagRequired("manifest_path")

	return cmd
}

func NewGetPipelineCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId       string
		namedPipelineId string
		saveDir         string // workflow manifest save dir
	)
	handler := func() error {
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoGetProjectPipelineParams{
			ProjectID:       projectId,
			NamedPipelineID: namedPipelineId,
		}

		kafeidoGetProjectPipelineOk, err := runCmd.stub.Kafeido.KafeidoGetProjectPipeline(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewPipelineDetailsFormatter([]*appmodels.KafeidoGetProjectPipelineResponse{kafeidoGetProjectPipelineOk.Payload}, saveDir).Write(ioStreams.Out)

	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}
	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&namedPipelineId, "named_pipeline_id", "", "named pipeline id")
	cmd.MarkFlagRequired("named_pipeline_id")
	cmd.Flags().StringVar(&saveDir, "save_dir", "", "temporary save dir")
	return cmd
}

func NewListPipelineCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
		saveDir   string // workflow manifest save dir
	)
	handler := func() error {
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoListProjectPipelineParams{
			ProjectID: projectId,
		}
		kafeidoListProjectPipelineOk, err := runCmd.stub.Kafeido.KafeidoListProjectPipeline(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return err
		}
		var listResp []*appmodels.KafeidoGetProjectPipelineResponse
		for _, namedPipelineId := range kafeidoListProjectPipelineOk.Payload.NamedPipelineID {
			subParams := &appservice.KafeidoGetProjectPipelineParams{
				ProjectID:       projectId,
				NamedPipelineID: namedPipelineId,
			}
			kafeidoGetProjectPipelineOk, err := runCmd.stub.Kafeido.KafeidoGetProjectPipeline(
				subParams.WithTimeout(runCmd.requestTimeout),
				runCmd.authInformer(),
			)
			if err == nil {
				listResp = append(listResp, kafeidoGetProjectPipelineOk.Payload)
			}
		}

		return format.NewPipelineDetailsFormatter(listResp, saveDir).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}
	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.Flags().StringVar(&saveDir, "save_dir", "", "temporary save dir")
	cmd.MarkFlagRequired("project_id")
	return cmd
}

func NewRunPipelineCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId       string
		namedPipelineId string
		jsonParams      string
	)
	handler := func() error {
		jsonMap, err := ensureJsonMarshalable([]byte(jsonParams))
		if err != nil {
			return err
		}

		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoRunProjectPipelineParams{
			Body: appservice.KafeidoRunProjectPipelineBody{
				Parameters: jsonMap,
			},
			ProjectID:       projectId,
			NamedPipelineID: namedPipelineId,
		}

		kafeidoRunProjectPipelineOk, err := runCmd.stub.Kafeido.KafeidoRunProjectPipeline(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewPipelineRunMessageFormatter(kafeidoRunProjectPipelineOk.Payload, "Pipeline Run has been submitted").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&namedPipelineId, "named_pipeline_id", "", "named pipeline id")
	cmd.MarkFlagRequired("named_pipeline_id")
	cmd.Flags().StringVar(&jsonParams, "run_parameters", "{}", "run parameters in json string format:")

	return cmd
}

func ensureJsonMarshalable(input []byte) (map[string]string, error) {
	var v map[string]string
	err := json.Unmarshal(input, &v)
	return v, err
}

func NewPutPipelineCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId       string
		namedPipelineId string
		manifestPath    string // workflow manifest path
	)
	handler := func() error {
		manifestBytes, err := ioutil.ReadFile(manifestPath)
		if err != nil {
			return fmt.Errorf("failed to read file, err:%+v\n", err)
		}
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoPutProjectPipelineParams{
			Body: appservice.KafeidoPutProjectPipelineBody{
				PipelineWorkflowInRawYaml: strfmt.Base64(manifestBytes),
			},
			ProjectID:       projectId,
			NamedPipelineID: namedPipelineId,
		}

		kafeidoPutProjectPipelineOk, err := runCmd.stub.Kafeido.KafeidoPutProjectPipeline(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewPipelineUpdateMessageFormatter(kafeidoPutProjectPipelineOk.Payload, "pipeline is updated. check for new pipeline version Id").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&namedPipelineId, "named_pipeline_id", "", "pipeline id")
	cmd.MarkFlagRequired("named_pipeline_id")
	cmd.Flags().StringVar(&manifestPath, "manifest_path", "", "manifeset manifest local path")
	cmd.MarkFlagRequired("manifest_path")

	return cmd
}

func NewDeletePipelineCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId       string
		namedPipelineId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoDeleteProjectPipelineParams{
			ProjectID:       projectId,
			NamedPipelineID: namedPipelineId,
		}

		_, err := runCmd.stub.Kafeido.KafeidoDeleteProjectPipeline(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(namedPipelineId, "pipeline is deleted.").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "pipeline",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&namedPipelineId, "named_pipeline_id", "", "pipeline id")
	cmd.MarkFlagRequired("named_pipeline_id")

	return cmd
}
