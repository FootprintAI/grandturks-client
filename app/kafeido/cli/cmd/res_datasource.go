package cmd

import (
	"strconv"
	"time"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
	pkgproto "github.com/footprintai/grandturks-client/v2/pkg/grpc/proto"
)

func NewCreateDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {

	cmd := &cobra.Command{
		Use:   "datasource",
		Short: "",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
	}
	cmd.AddCommand(NewCreateImageDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateVideoDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateStreamingDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateYoutubeDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateImageUrlDataSourceCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateAudioFilesDataSourceCommand(logger, ioStreams))
	return cmd
}

func NewCreateImageDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId                   string
		protocol                    string
		endpoint                    string
		bucketName                  string
		indexObject                 string
		repeatedTimes               int32
		delayInDurationForNextFrame time.Duration
	)

	handler := func() error {
		storageInfo := &appmodels.DatastreamStorageInfo{
			ObjectStoreInfo: &appmodels.CompdatastreamObjectStoreInfo{
				Endpoint: endpoint,
				Protocol: protocol,
			},
		}
		params := &appservice.KafeidoCreateDataSourceParams{
			Body: appservice.KafeidoCreateDataSourceBody{
				DataSourceInfo: &appmodels.KafeidoDataSourceInfo{
					Type: appmodels.DatastreamDataSourceDATASOURCEPHOTO.Pointer(),
					PhotoDataSource: &appmodels.DatastreamPhotoDataSource{
						BucketName:    bucketName,
						IndexObject:   indexObject,
						RepeatedTimes: repeatedTimes,
						StorageInfo:   storageInfo,
					},
					DelayInDurationForNextFrame: pkgproto.SerializeDuration(delayInDurationForNextFrame),
				},
			},
			ProjectID: projectId,
		}
		return runCreateDataSourceCmd(params, ioStreams)
	}

	cmd := &cobra.Command{
		Use:   "image",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&protocol, "protocol", "http", "data source protocol (default: http")
	cmd.Flags().StringVar(&endpoint, "endpoint", "minio-service.grandturks.svc.cluster.local:9000", "data source protocol (default: minio-service.grandturks.svc.cluster.local:9000")
	cmd.Flags().StringVar(&bucketName, "bucket_name", "", "data source bucket name (default: ")
	cmd.Flags().StringVar(&indexObject, "index_object", "", "data source index object (default: ")
	cmd.Flags().Int32Var(&repeatedTimes, "repeated_times", 1, "data source repeated times(default: 1")
	cmd.Flags().DurationVar(&delayInDurationForNextFrame, "next_frame_delay", time.Second, "duration per frame (default: 1s")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("bucket")
	cmd.MarkFlagRequired("index_object")

	return cmd

}

func runCreateDataSourceCmd(params *appservice.KafeidoCreateDataSourceParams, ioStreams genericclioptions.IOStreams) error {
	runCmd := mustNewRunCmd()
	kafeidoCreateDataSourceOk, err := runCmd.stub.Kafeido.KafeidoCreateDataSource(
		params.WithTimeout(runCmd.requestTimeout),
		runCmd.authInformer(),
	)
	if err != nil {
		return openapiErrorParser(err)
	}
	return format.NewIdAndMessageFormatter(kafeidoCreateDataSourceOk.Payload.DataSourceID, "DataSource is created").Write(ioStreams.Out)
}

func NewCreateVideoDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId                   string
		protocol                    string
		endpoint                    string
		bucketName                  string
		videoPath                   string
		repeatedTimes               int32
		framesPerSecond             string // convert to float32
		delayInDurationForNextFrame time.Duration
	)

	handler := func() error {
		storageInfo := &appmodels.DatastreamStorageInfo{
			ObjectStoreInfo: &appmodels.CompdatastreamObjectStoreInfo{
				Endpoint: endpoint,
				Protocol: protocol,
			},
		}

		params := &appservice.KafeidoCreateDataSourceParams{
			Body: appservice.KafeidoCreateDataSourceBody{
				DataSourceInfo: &appmodels.KafeidoDataSourceInfo{
					Type: appmodels.DatastreamDataSourceDATASOURCEVIDEO.Pointer(),
					VideoDataSource: &appmodels.DatastreamVideoDataSource{
						BucketName:      bucketName,
						VideoPath:       videoPath,
						RepeatedTimes:   repeatedTimes,
						FramesPerSecond: convertStr2Float32(framesPerSecond),
						StorageInfo:     storageInfo,
					},
					DelayInDurationForNextFrame: pkgproto.SerializeDuration(delayInDurationForNextFrame),
				},
			},
			ProjectID: projectId,
		}
		return runCreateDataSourceCmd(params, ioStreams)
	}

	cmd := &cobra.Command{
		Use:   "video",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&protocol, "protocol", "http", "data source protocol (default: http")
	cmd.Flags().StringVar(&endpoint, "endpoint", "", "data source protocol (default: ")
	cmd.Flags().StringVar(&bucketName, "bucket_name", "", "data source bucket name (default: ")
	cmd.Flags().StringVar(&videoPath, "video_path", "", "video path (default: ")
	cmd.Flags().StringVar(&framesPerSecond, "fps", "1", "frames per second (default: 1")
	cmd.Flags().Int32Var(&repeatedTimes, "repeated_times", 1, "data source repeated times(default: 1")
	cmd.Flags().DurationVar(&delayInDurationForNextFrame, "next_frame_delay", time.Second, "duration per frame (default: 1s")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("bucket")
	cmd.MarkFlagRequired("video_path")

	return cmd

}

func NewCreateStreamingDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId                   string
		endpoint                    string
		protocol                    string
		channel                     string
		user                        string
		password                    string
		framesPerSecond             string // convert to float32
		delayInDurationForNextFrame time.Duration
	)

	handler := func() error {
		params := &appservice.KafeidoCreateDataSourceParams{
			Body: appservice.KafeidoCreateDataSourceBody{
				DataSourceInfo: &appmodels.KafeidoDataSourceInfo{
					Type: appmodels.DatastreamDataSourceDATASOURCESTREAMING.Pointer(),
					StreamingDataSource: &appmodels.DatastreamStreamingDataSource{
						ChannelName:     channel,
						FramesPerSecond: convertStr2Float32(framesPerSecond),
						StreamingInfo: &appmodels.DatastreamStreamingInfo{
							Account:  user,
							Endpoint: endpoint,
							Password: password,
							Protocol: format.NormalizedStreamingInfoProtocol(protocol).Pointer(),
						},
					},
					DelayInDurationForNextFrame: pkgproto.SerializeDuration(delayInDurationForNextFrame),
				},
			},
			ProjectID: projectId,
		}
		return runCreateDataSourceCmd(params, ioStreams)
	}

	cmd := &cobra.Command{
		Use:   "streaming",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&endpoint, "endpoint", "rtsp.streamingmedia.svc.cluster.local:1935", "protocol (default: rtsp.streamingmedia.svc.cluster.local:1935")
	cmd.Flags().StringVar(&protocol, "protocol", "rtmp", "streaming protocol (default: rtmp)")
	cmd.Flags().StringVar(&channel, "channel", "", "streaming channel name(default: ")
	cmd.Flags().StringVar(&user, "user", "", "user account(default: ")
	cmd.Flags().StringVar(&password, "password", "", "user account(default: ")
	cmd.Flags().StringVar(&framesPerSecond, "fps", "1", "frames per second (default: 1")
	cmd.Flags().DurationVar(&delayInDurationForNextFrame, "next_frame_delay", time.Second, "duration per frame (default: 1s")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("channel")

	return cmd

}

func convertStr2Float32(s string) float32 {
	v, _ := strconv.ParseFloat(s, 32)
	return float32(v)
}

func NewCreateYoutubeDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId                   string
		videoId                     string
		streamingDuration           time.Duration // duration
		delayInDurationForNextFrame time.Duration
	)

	handler := func() error {
		params := &appservice.KafeidoCreateDataSourceParams{
			Body: appservice.KafeidoCreateDataSourceBody{
				DataSourceInfo: &appmodels.KafeidoDataSourceInfo{
					Type: appmodels.DatastreamDataSourceDATASOURCEYOUTUBE.Pointer(),
					YoutubeDataSource: &appmodels.DatastreamYoutubeDataSource{
						VideoID:           videoId,
						StreamingDuration: pkgproto.SerializeDuration(streamingDuration),
					},
					DelayInDurationForNextFrame: pkgproto.SerializeDuration(delayInDurationForNextFrame),
				},
			},
			ProjectID: projectId,
		}
		return runCreateDataSourceCmd(params, ioStreams)
	}

	cmd := &cobra.Command{
		Use:   "youtube",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().DurationVar(&streamingDuration, "duration", time.Hour, "streaming duration (default: 1h")
	cmd.Flags().StringVar(&videoId, "video_id", "", "video Id(default: ")
	cmd.Flags().DurationVar(&delayInDurationForNextFrame, "next_frame_delay", time.Second, "duration per frame (default: 1s")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("video_id")

	return cmd

}

func NewCreateImageUrlDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId                   string
		imageUrl                    string
		streamingDuration           time.Duration // duration
		delayInDurationForNextFrame time.Duration
	)

	handler := func() error {
		params := &appservice.KafeidoCreateDataSourceParams{
			Body: appservice.KafeidoCreateDataSourceBody{
				DataSourceInfo: &appmodels.KafeidoDataSourceInfo{
					Type: appmodels.DatastreamDataSourceDATASOURCEIMAGEURL.Pointer(),
					ImageURLDataSource: &appmodels.DatastreamImageURLDataSource{
						URL:               imageUrl,
						StreamingDuration: pkgproto.SerializeDuration(streamingDuration),
					},
					DelayInDurationForNextFrame: pkgproto.SerializeDuration(delayInDurationForNextFrame),
				},
			},
			ProjectID: projectId,
		}
		return runCreateDataSourceCmd(params, ioStreams)
	}

	cmd := &cobra.Command{
		Use:   "imageurl",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().DurationVar(&streamingDuration, "streaming_duration", time.Hour, "streaming duration (default: 1h")
	cmd.Flags().StringVar(&imageUrl, "image_url", "", "url for image resource(default: ")
	cmd.Flags().DurationVar(&delayInDurationForNextFrame, "next_frame_delay", time.Second, "duration per frame (default: 1s")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("video_id")

	return cmd

}

func NewCreateAudioFilesDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId                   string
		protocol                    string
		endpoint                    string
		bucketName                  string
		indexObject                 string // a file path containing a newline delitered file, each line is with an audio file
		repeatedTimes               int32
		delayInDurationForNextFrame time.Duration
	)

	handler := func() error {
		storageInfo := &appmodels.DatastreamStorageInfo{
			ObjectStoreInfo: &appmodels.CompdatastreamObjectStoreInfo{
				Endpoint: endpoint,
				Protocol: protocol,
			},
		}

		params := &appservice.KafeidoCreateDataSourceParams{
			Body: appservice.KafeidoCreateDataSourceBody{
				DataSourceInfo: &appmodels.KafeidoDataSourceInfo{
					Type: appmodels.DatastreamDataSourceDATASOURCEAUDIOFILES.Pointer(),
					AudioFilesDataSource: &appmodels.DatastreamAudioFilesDataSource{
						BucketName:    bucketName,
						IndexObject:   indexObject,
						RepeatedTimes: repeatedTimes,
						StorageInfo:   storageInfo,
					},
					DelayInDurationForNextFrame: pkgproto.SerializeDuration(delayInDurationForNextFrame),
				},
			},
			ProjectID: projectId,
		}
		return runCreateDataSourceCmd(params, ioStreams)
	}

	cmd := &cobra.Command{
		Use:   "audiofile",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&protocol, "protocol", "http", "data source protocol (default: http")
	cmd.Flags().StringVar(&endpoint, "endpoint", "", "data source protocol (default: ")
	cmd.Flags().StringVar(&bucketName, "bucket_name", "", "data source bucket name (default: ")
	cmd.Flags().StringVar(&indexObject, "index_object", "", "audio path (default: ")
	cmd.Flags().Int32Var(&repeatedTimes, "repeated_times", 1, "data source repeated times(default: 1")
	cmd.Flags().DurationVar(&delayInDurationForNextFrame, "next_frame_delay", time.Second, "duration in delay for next streaming frame(default: 1s")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("bucket")
	cmd.MarkFlagRequired("index_object")

	return cmd
}

func NewListDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoListDataSourceParams{
			ProjectID: projectId,
		}
		kafeidoListDataSourceOk, err := runCmd.stub.Kafeido.KafeidoListDataSource(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		var listResp []*appmodels.AppkafeidoGetDataSourceResponse
		for _, dataSourceId := range kafeidoListDataSourceOk.Payload.DataSourceIds {
			subParams := &appservice.KafeidoGetDataSourceParams{
				ProjectID:    projectId,
				DataSourceID: dataSourceId,
			}
			kafeidoGetDataSourceOk, err := runCmd.stub.Kafeido.KafeidoGetDataSource(
				subParams.WithTimeout(runCmd.requestTimeout),
				runCmd.authInformer(),
			)
			if err == nil {
				listResp = append(listResp, kafeidoGetDataSourceOk.Payload)
			}
		}

		return format.NewDataSourceDetailsFormatter(listResp).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "datasource",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.MarkFlagRequired("project_id")

	return cmd

}

func NewGetDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId    string
		dataSourceId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoGetDataSourceParams{
			ProjectID:    projectId,
			DataSourceID: dataSourceId,
		}
		kafeidoGetDataSourceOk, err := runCmd.stub.Kafeido.KafeidoGetDataSource(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewDataSourceDetailsFormatter([]*appmodels.AppkafeidoGetDataSourceResponse{kafeidoGetDataSourceOk.Payload}).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "datasource",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&dataSourceId, "datasource_id", "", "data source Id")
	cmd.MarkFlagRequired("datasource_id")

	return cmd

}

func NewDeleteDataSourceCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId    string
		dataSourceId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd()

		params := &appservice.KafeidoDeleteDataSourceParams{
			ProjectID:    projectId,
			DataSourceID: dataSourceId,
		}
		_, err := runCmd.stub.Kafeido.KafeidoDeleteDataSource(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(dataSourceId, "DataSource has been deleted").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "datasource",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.MarkFlagRequired("project_id")
	cmd.Flags().StringVar(&dataSourceId, "datasource_id", "", "data source Id")
	cmd.MarkFlagRequired("datasource_id")

	return cmd

}
