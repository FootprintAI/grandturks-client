package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
	clihelper "github.com/footprintai/grandturks-client/v2/app/kafeido/cli/helper"
)

func NewCreateInferenceJobCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId          string
		modelInferenceId   string
		dataSourceId       string
		concurrentRequests int
		outputWidth        int
		outputHeight       int
		jobKind            string

		outputKey string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceCreateModelInferenceJobParams{
			Body: appservice.KafeidoServiceCreateModelInferenceJobBody{
				JobKind: clihelper.MakeOpenAPIModelInferenceJob(clihelper.CliJobType(jobKind)),
				DataSinkConfig: &appmodels.KafeidoDataSinkConfig{
					TaskWorkerDataSink: &appmodels.ComptaskworkerDataSink{
						RestcolDataSink: &appmodels.TaskworkerRestColDataSink{
							CollectionID: outputKey,
						},
						AppDataSink: &appmodels.TaskworkerAppDataSink{
							Labels: []*appmodels.TaskworkerAppDataSinkLabel{
								&appmodels.TaskworkerAppDataSinkLabel{
									Name:  "tasklabel",
									Value: outputKey,
								},
							},
						},
						ObjectStoreDataSink: &appmodels.ComptaskworkerObjectStoreDataSink{
							ObjectPrefix: outputKey,
						},
						RedisDataSink: &appmodels.ComptaskworkerRedisStreamSink{
							StreamName: outputKey,
						},
					},
				},
				DataSourceID:       dataSourceId,
				ModelInferenceID:   modelInferenceId,
				ConcurrentRequests: int32(concurrentRequests),
				PrometheusJobLabel: outputKey,
				ImageOutputFormat: &appmodels.DatastreamImageOutputFormat{
					Encoding:    appmodels.NewDatastreamImageOutputFormatEncodingType(appmodels.DatastreamImageOutputFormatEncodingTypeENCODINGTYPEUNSPECIFIED /*hsiny: no encoding from datastream*/),
					PhotoFormat: appmodels.NewImageOutputFormatPhotoFormat(appmodels.ImageOutputFormatPhotoFormatPHOTOFORMATPNG),
					Width:       int32(outputWidth),
					Height:      int32(outputHeight),
				},
				AudioOutputFormat: &appmodels.DatastreamAudioOutputFormat{
					Encoding: appmodels.NewDatastreamAudioOutputFormatEncodingType(appmodels.DatastreamAudioOutputFormatEncodingTypeENCODINGTYPEUNSPECIFIED /*hsiny: no encoding from datastream*/),
					Format:   appmodels.NewAudioOutputFormatAudioFormat(appmodels.AudioOutputFormatAudioFormatAUDIOFORMATMPEG),
				},
			},
			ProjectID: projectId,
		}

		kafeidoCreateModelInferenceJobOk, err := runCmd.stub.KafeidoService.KafeidoServiceCreateModelInferenceJob(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(kafeidoCreateModelInferenceJobOk.Payload.ModelInferenceJobID, "Inference job is created").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "job",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelInferenceId, "inference_id", "", "inference id")
	cmd.Flags().StringVar(&dataSourceId, "datasource_id", "", "data source id")
	cmd.Flags().IntVar(&concurrentRequests, "job_concurrency", 1, "concurrent requets to model inferencing (default=  1")
	cmd.Flags().StringVar(&outputKey, "output", "", "key, a raw string, for job output, including streaming, storage, and appdata, ...")
	cmd.Flags().IntVar(&outputWidth, "width", 0, "width in px, default: 0")
	cmd.Flags().IntVar(&outputHeight, "height", 0, "height in px, default: 0")
	cmd.Flags().StringVar(&jobKind, "job_type", clihelper.CliJobImage.String(), "image or audio, default: image")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("inference_id")
	cmd.MarkFlagRequired("data_id")
	cmd.MarkFlagRequired("output")
	cmd.MarkFlagRequired("job_type")

	return cmd
}

func NewListInferenceJobCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceListModelInferenceJobParams{
			ProjectID: projectId,
		}

		kafeidoListModelInferenceJobOk, err := runCmd.stub.KafeidoService.KafeidoServiceListModelInferenceJob(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		var listResp []*appmodels.KafeidoGetModelInferenceJobResponse
		for _, modelInferenceJobId := range kafeidoListModelInferenceJobOk.Payload.ModelInferenceJobIds {
			subParams := &appservice.KafeidoServiceGetModelInferenceJobParams{
				ProjectID:           projectId,
				ModelInferenceJobID: modelInferenceJobId,
			}
			kafeidoGetModelInferenceJobOk, err := runCmd.stub.KafeidoService.KafeidoServiceGetModelInferenceJob(
				subParams.WithTimeout(runCmd.requestTimeout),
				runCmd.authInformer(),
			)
			if err == nil {
				listResp = append(listResp, kafeidoGetModelInferenceJobOk.Payload)
			}
		}

		return format.NewModelInferenceJobDetailsFormatter(listResp).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "job",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.MarkFlagRequired("project_id")

	return cmd
}

func NewGetInferenceJobCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId           string
		modelInferenceJobId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceGetModelInferenceJobParams{
			ProjectID:           projectId,
			ModelInferenceJobID: modelInferenceJobId,
		}
		kafeidoGetModelInferenceJobOk, err := runCmd.stub.KafeidoService.KafeidoServiceGetModelInferenceJob(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewModelInferenceJobDetailsFormatter([]*appmodels.KafeidoGetModelInferenceJobResponse{kafeidoGetModelInferenceJobOk.Payload}).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "job",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelInferenceJobId, "job_id", "", "job id")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("job_id")

	return cmd
}

func NewCancelInferenceJobCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId           string
		modelInferenceJobId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceCancelModelInferenceJobParams{
			ProjectID:           projectId,
			ModelInferenceJobID: modelInferenceJobId,
		}
		_, err := runCmd.stub.KafeidoService.KafeidoServiceCancelModelInferenceJob(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(modelInferenceJobId, "Job has been cancelled").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "job",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelInferenceJobId, "job_id", "", "job id")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("job_id")

	return cmd
}
