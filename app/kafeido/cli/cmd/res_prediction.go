package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	pkgio "github.com/footprintai/grandturks-client/v2/pkg/io"
	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"
)

func NewCreatePredictionCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId        string
		modelInferenceId string
		queryFile        string
		queryLang        string
		queryType        string
	)
	handler := func() error {

		rc, err := loadFile(queryFile)
		if err != nil {
			return err
		}
		defer rc.Close()

		runCmd := mustNewRunCmd()
		params := &appservice.KafeidoCreatePredictionParams{
			ProjectID:        projectId,
			ModelInferenceID: modelInferenceId,
			Body: appservice.KafeidoCreatePredictionBody{
				Concurrency: 1,
			},
		}
		if queryType == "image" {
			params.Body.Type = appmodels.KafeidoTypePredictionDataPHOTO.Pointer()
			params.Body.ImageRequests = []*appmodels.KafeidoPredictImageRequestBody{
				&appmodels.KafeidoPredictImageRequestBody{
					Key: filepath.Base(queryFile),
					ImageBytes: &appmodels.KafeidoRequestInBytes{
						B64: pkgio.MustReadAll2Base64(rc),
					},
				},
			}
		} else if queryType == "audio" {
			params.Body.Type = appmodels.KafeidoTypePredictionDataAUDIO.Pointer()
			params.Body.AudioRequests = []*appmodels.KafeidoPredictAudioRequestBody{
				&appmodels.KafeidoPredictAudioRequestBody{
					Key:  filepath.Base(queryFile),
					Lang: []string{queryLang},
					AudioBytes: &appmodels.KafeidoRequestInBytes{
						B64: pkgio.MustReadAll2Base64(rc),
					},
				},
			}
		}
		started := time.Now()
		kafeidoCreatePredictionOk, err := runCmd.stub.Kafeido.KafeidoCreatePrediction(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		timeDur := time.Now().Sub(started)
		if err != nil {
			return openapiErrorParser(err)
		}
		var jsonStringResp string
		if len(kafeidoCreatePredictionOk.Payload.ImageResponses) > 0 {
			jsonStringResp = kafeidoCreatePredictionOk.Payload.ImageResponses[0].Prediction.JSON
		} else if len(kafeidoCreatePredictionOk.Payload.AudioResponses) > 0 {
			jsonStringResp = kafeidoCreatePredictionOk.Payload.AudioResponses[0].Prediction.JSON
		}
		// decode json string

		j := make(map[string]interface{})
		json.Unmarshal([]byte(jsonStringResp), &j)
		jsonBlob, _ := json.Marshal(j)
		ioStreams.Out.Write(jsonBlob)
		ioStreams.Out.Write([]byte(fmt.Sprintf("\n cost :%s\n", timeDur)))
		return err
	}

	cmd := &cobra.Command{
		Use:   "predict",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project Id")
	cmd.Flags().StringVar(&modelInferenceId, "inference_id", "", "inference id")
	cmd.Flags().StringVar(&queryFile, "query_file", "", "query file path")
	cmd.Flags().StringVar(&queryType, "query_type", "image", "query type: image or audio (default: image")
	cmd.Flags().StringVar(&queryLang, "query_lang", "en", "query lang: en or zh(default: en")
	cmd.MarkFlagRequired("project_id")
	cmd.MarkFlagRequired("inference_id")
	cmd.MarkFlagRequired("query_type")
	cmd.MarkFlagRequired("query_file")

	return cmd
}

func loadFile(f string) (io.ReadCloser, error) {
	return os.Open(f)
}
