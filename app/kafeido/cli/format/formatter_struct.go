package format

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	pkgurl "net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/minio/minio-go/v7"

	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	appswaggermodel "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	clihelper "github.com/footprintai/grandturks-client/v2/app/kafeido/cli/helper"
)

type IdAndMessageFormatter struct {
	id  string
	msg string
}

func (i *IdAndMessageFormatter) Header() []string {
	return []string{"id", "message"}
}

func (i *IdAndMessageFormatter) Rows() [][]string {
	return [][]string{
		[]string{i.id, i.msg},
	}
}

func (i *IdAndMessageFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

func NewIdAndMessageFormatter(id string, msg string) *IdAndMessageFormatter {
	return &IdAndMessageFormatter{
		id:  id,
		msg: msg,
	}
}

type ListProjectDetailsFormatter struct {
	pbBodys []*appswaggermodel.KafeidoGetProjectResponse
}

func (p *ListProjectDetailsFormatter) Header() []string {
	return []string{
		"id",
		"name",
		"desc",
		"status",
		"public_bucket",
		"private_bucket",
	}
}

func (p *ListProjectDetailsFormatter) Rows() [][]string {
	var vals [][]string
	for _, pbBody := range p.pbBodys {
		vals = append(vals, []string{
			pbBody.ProjectID,
			pbBody.Name,
			pbBody.Desc,
			pbBody.Status,
			pbBody.ObjectStoreProjectPublicBucketName,
			pbBody.ObjectStoreProjectPrivateBucketName,
		})
	}
	return vals
}

func (p *ListProjectDetailsFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(p).Write(w)
}

func NewListProjectDetailsFormatter(pbBodys []*appswaggermodel.KafeidoGetProjectResponse) *ListProjectDetailsFormatter {
	return &ListProjectDetailsFormatter{
		pbBodys: pbBodys,
	}
}

type PipelineDetailsFormatter struct {
	pbBodys []*appswaggermodel.KafeidoGetProjectPipelineResponse
	saveDir string
}

func (p *PipelineDetailsFormatter) Header() []string {
	return []string{
		"project_Id",
		"named_pipeline_id",
		"named_pipeline_name",
		"named_pipeline_type",
		"workflow_manifest",
		"run_parameters",
	}
}

func (p *PipelineDetailsFormatter) Rows() [][]string {
	var resp [][]string
	for _, row := range p.pbBodys {
		resp = append(resp, p.Row(row))
	}
	return resp
}

func (p *PipelineDetailsFormatter) Row(resp *appswaggermodel.KafeidoGetProjectPipelineResponse) []string {
	var shortPipelineWorkflowStr string
	if len(p.saveDir) == 0 {
		// use abbreviation
		shortPipelineWorkflowStr = "..."

	} else {
		fname := mustGetTempFileWithDir(p.saveDir, "*.workflow.yaml")
		_, err := fname.Write([]byte(resp.KfPipelineWorkflow))
		if err != nil {
			fmt.Printf("failed to write: %s\n", err)
		}
		shortPipelineWorkflowStr = fname.Name()
	}

	return []string{
		resp.ProjectID,
		resp.NamedPipelineID,
		resp.NamedPipelineName,
		resp.NamedPipelineType,
		shortPipelineWorkflowStr,
		mustEncodeJson(resp.RunParameters),
	}
}

func mustGetTempFileWithDir(dir, pattern string) *os.File {
	f, err := ioutil.TempFile(dir, pattern)
	if err != nil {
		panic(err)
	}
	return f
}

func mustEncodeJson(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func (p *PipelineDetailsFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(p).Write(w)
}

func NewPipelineDetailsFormatter(pbBodys []*appswaggermodel.KafeidoGetProjectPipelineResponse, saveDir string) *PipelineDetailsFormatter {
	return &PipelineDetailsFormatter{
		pbBodys: pbBodys,
		saveDir: saveDir,
	}
}

type PipelineUpdateMessageFormatter struct {
	pbBody *appswaggermodel.KafeidoPutProjectPipelineResponse
	msg    string
}

func (i *PipelineUpdateMessageFormatter) Header() []string {
	return []string{
		"project_id",
		"named_pipeline_id",
		"kf_pipeline_id",
		"kf_pipeline_version_id",
		"message",
	}
}

func (i *PipelineUpdateMessageFormatter) Rows() [][]string {
	return [][]string{
		[]string{
			i.pbBody.ProjectID,
			i.pbBody.NamedPipelineID,
			i.pbBody.KfPipelineID,
			i.pbBody.KfPipelineVersionID,
			i.msg,
		},
	}
}

func (i *PipelineUpdateMessageFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

func NewPipelineUpdateMessageFormatter(pbBody *appswaggermodel.KafeidoPutProjectPipelineResponse, msg string) *PipelineUpdateMessageFormatter {
	return &PipelineUpdateMessageFormatter{
		pbBody: pbBody,
		msg:    msg,
	}
}

type PipelineRunMessageFormatter struct {
	pbBody *appswaggermodel.KafeidoRunProjectPipelineResponse
	msg    string
}

func (i *PipelineRunMessageFormatter) Header() []string {
	return []string{
		"project_id",
		"named_pipeline_id",
		"kf_pipeline_id",
		"kf_pipeline_version_id",
		"experiment_id",
		"run_id",
		"message",
	}
}

func (i *PipelineRunMessageFormatter) Rows() [][]string {
	return [][]string{
		[]string{
			i.pbBody.ProjectID,
			i.pbBody.NamedPipelineID,
			i.pbBody.KfPipelineID,
			i.pbBody.KfPipelineVersionID,
			i.pbBody.ExperimentID,
			i.pbBody.RunID,
			i.msg,
		},
	}
}

func (i *PipelineRunMessageFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

func NewPipelineRunMessageFormatter(pbBody *appswaggermodel.KafeidoRunProjectPipelineResponse, msg string) *PipelineRunMessageFormatter {
	return &PipelineRunMessageFormatter{
		pbBody: pbBody,
		msg:    msg,
	}
}

type ModelInferenceDetailsFormatter struct {
	pbBodys []*appswaggermodel.KafeidoGetModelInferenceResponse
}

func (i *ModelInferenceDetailsFormatter) Header() []string {
	return []string{
		"project_id",
		"named_pipeline_id",
		"model_inference_id",
		"model_name",
		"model_uri",
		"model_status",
		"model_supported",
	}
}

func (i *ModelInferenceDetailsFormatter) Rows() [][]string {
	var out [][]string
	for _, pbBody := range i.pbBodys {
		out = append(out, []string{
			pbBody.ProjectID,
			pbBody.NamedPipelineID,
			pbBody.ModelInferenceID,
			pbBody.ModelName,
			pbBody.ModelURI,
			fmt.Sprintf("%v", pbBody.IsModelReady),
			fmt.Sprintf("%v", pbBody.IsInputSupported),
		})
	}
	return out
}

func (i *ModelInferenceDetailsFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

func NewModelInferenceDetailsFormatter(pbBodys []*appswaggermodel.KafeidoGetModelInferenceResponse) *ModelInferenceDetailsFormatter {
	return &ModelInferenceDetailsFormatter{
		pbBodys: pbBodys,
	}
}

type ModelInferenceJobDetailsFormatter struct {
	pbBodys []*appmodels.KafeidoGetModelInferenceJobResponse
}

func (i *ModelInferenceJobDetailsFormatter) Header() []string {
	return []string{
		"project_id",
		"inference_id",
		"datasource_id",
		"job_id",
		"job_kind",
		"job_status",
		"job_concurrency",
		"job_started_at",
		"job_ended_at",
		"output_storage_prefix",
		"output_redisstream",
	}
}

func (i *ModelInferenceJobDetailsFormatter) Rows() [][]string {
	var vals [][]string
	for _, pbBody := range i.pbBodys {
		vals = append(vals, []string{
			pbBody.ProjectID,
			pbBody.ModelInferenceID,
			pbBody.DataSourceID,
			pbBody.ModelInferenceJobID,
			clihelper.MakeCliJobType(pbBody.JobKind).String(),
			pbBody.ModelInferenceJobStatus,
			fmt.Sprintf("%d", pbBody.ConcurrentRequests),
			pbBody.ModelInferenceJobStartedAt.String(),
			pbBody.ModelInferenceJobEndedAt.String(),
			pbBody.SinkDestination.ObjectStoreDataSink.ObjectPrefix,
			pbBody.SinkDestination.RedisDataSink.StreamName,
		})
	}
	return vals
}

func (i *ModelInferenceJobDetailsFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

func NewModelInferenceJobDetailsFormatter(pbBodys []*appmodels.KafeidoGetModelInferenceJobResponse) *ModelInferenceJobDetailsFormatter {
	return &ModelInferenceJobDetailsFormatter{
		pbBodys: pbBodys,
	}
}

type DataSourceDetailsFormatter struct {
	pbBodys []*appswaggermodel.AppkafeidoGetDataSourceResponse
}

func (i *DataSourceDetailsFormatter) Header() []string {
	return []string{
		"project_id",
		"datasource_id",
		"data_type",
		"protocol",
		"endpoint",
		"path",
		"details",
	}
}

func (i *DataSourceDetailsFormatter) Rows() [][]string {
	var vals [][]string
	for _, pbBody := range i.pbBodys {
		var val []string
		if appswaggermodel.DatastreamDataSourceDATASOURCEPHOTO == *pbBody.DataSourceInfo.Type {
			val = serializedPhotoDataSource(pbBody)
		} else if appswaggermodel.DatastreamDataSourceDATASOURCEVIDEO == *pbBody.DataSourceInfo.Type {
			val = serializedVideoDataSource(pbBody)
		} else if appswaggermodel.DatastreamDataSourceDATASOURCESTREAMING == *pbBody.DataSourceInfo.Type {
			val = serializedStreamingDataSource(pbBody)
		} else if appswaggermodel.DatastreamDataSourceDATASOURCEYOUTUBE == *pbBody.DataSourceInfo.Type {
			val = serializedYoutubeDataSource(pbBody)
		} else if appswaggermodel.DatastreamDataSourceDATASOURCEIMAGEURL == *pbBody.DataSourceInfo.Type {
			val = serializedImageUrlDataSource(pbBody)
		} else if appswaggermodel.DatastreamDataSourceDATASOURCEAUDIOFILES == *pbBody.DataSourceInfo.Type {
			val = serializedAudioFilesDataSource(pbBody)
		}
		vals = append(vals, val)
	}
	return vals
}

func serializedPhotoDataSource(p *appswaggermodel.AppkafeidoGetDataSourceResponse) []string {
	return []string{
		p.ProjectID,
		p.DataSourceID,
		DeNormalizedDataSource(*p.DataSourceInfo.Type),
		p.DataSourceInfo.PhotoDataSource.StorageInfo.ObjectStoreInfo.Protocol,
		p.DataSourceInfo.PhotoDataSource.StorageInfo.ObjectStoreInfo.Endpoint,
		filepath.Join(p.DataSourceInfo.PhotoDataSource.BucketName, p.DataSourceInfo.PhotoDataSource.IndexObject),
		fmt.Sprintf("repeated: %d, next frame: %s", p.DataSourceInfo.PhotoDataSource.RepeatedTimes, p.DataSourceInfo.DelayInDurationForNextFrame),
	}
}

func serializedVideoDataSource(p *appswaggermodel.AppkafeidoGetDataSourceResponse) []string {
	return []string{
		p.ProjectID,
		p.DataSourceID,
		DeNormalizedDataSource(*p.DataSourceInfo.Type),
		p.DataSourceInfo.VideoDataSource.StorageInfo.ObjectStoreInfo.Protocol,
		p.DataSourceInfo.VideoDataSource.StorageInfo.ObjectStoreInfo.Endpoint,
		filepath.Join(p.DataSourceInfo.VideoDataSource.BucketName, p.DataSourceInfo.VideoDataSource.VideoPath),
		fmt.Sprintf("repeated: %d, fps: %0.2f, next frame: %s", p.DataSourceInfo.VideoDataSource.RepeatedTimes, p.DataSourceInfo.VideoDataSource.FramesPerSecond, p.DataSourceInfo.DelayInDurationForNextFrame),
	}
}

func serializedStreamingDataSource(p *appswaggermodel.AppkafeidoGetDataSourceResponse) []string {
	return []string{
		p.ProjectID,
		p.DataSourceID,
		DeNormalizedDataSource(*p.DataSourceInfo.Type),
		DeNormalizedStreamingInfoProtocol(*p.DataSourceInfo.StreamingDataSource.StreamingInfo.Protocol),
		p.DataSourceInfo.StreamingDataSource.StreamingInfo.Endpoint,
		p.DataSourceInfo.StreamingDataSource.ChannelName,
		fmt.Sprintf("fps: %0.2f, next frame: %s", p.DataSourceInfo.StreamingDataSource.FramesPerSecond, p.DataSourceInfo.DelayInDurationForNextFrame),
	}
}

func mustConvertDuration(s string) time.Duration {
	dur, _ := runtime.Duration(s)
	if dur == nil {
		return time.Duration(-1)
	}
	return dur.AsDuration()
}

func serializedYoutubeDataSource(p *appswaggermodel.AppkafeidoGetDataSourceResponse) []string {
	return []string{
		p.ProjectID,
		p.DataSourceID,
		DeNormalizedDataSource(*p.DataSourceInfo.Type),
		"hls",
		"www.youtube.com",
		p.DataSourceInfo.YoutubeDataSource.VideoID,
		fmt.Sprintf("fps:1, streaming duration: %s, next frame: %s\n", p.DataSourceInfo.YoutubeDataSource.StreamingDuration, p.DataSourceInfo.DelayInDurationForNextFrame),
	}
}

func serializedImageUrlDataSource(p *appswaggermodel.AppkafeidoGetDataSourceResponse) []string {
	parsedUrl, _ := pkgurl.Parse(p.DataSourceInfo.ImageURLDataSource.URL)
	return []string{
		p.ProjectID,
		p.DataSourceID,
		DeNormalizedDataSource(*p.DataSourceInfo.Type),
		parsedUrl.Scheme,
		parsedUrl.Host,
		parsedUrl.Path,
		fmt.Sprintf("fps:1, streaming duration: %s, next frame: %s\n", p.DataSourceInfo.ImageURLDataSource.StreamingDuration, p.DataSourceInfo.DelayInDurationForNextFrame),
	}
}

func serializedAudioFilesDataSource(p *appswaggermodel.AppkafeidoGetDataSourceResponse) []string {
	return []string{
		p.ProjectID,
		p.DataSourceID,
		DeNormalizedDataSource(*p.DataSourceInfo.Type),
		p.DataSourceInfo.AudioFilesDataSource.StorageInfo.ObjectStoreInfo.Protocol,
		p.DataSourceInfo.AudioFilesDataSource.StorageInfo.ObjectStoreInfo.Endpoint,
		filepath.Join(p.DataSourceInfo.AudioFilesDataSource.BucketName, p.DataSourceInfo.AudioFilesDataSource.IndexObject),
		fmt.Sprintf("repeated: %d, next frame: %s", p.DataSourceInfo.AudioFilesDataSource.RepeatedTimes, p.DataSourceInfo.DelayInDurationForNextFrame),
	}
}

func (i *DataSourceDetailsFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

func NewDataSourceDetailsFormatter(pbBodys []*appswaggermodel.AppkafeidoGetDataSourceResponse) *DataSourceDetailsFormatter {
	return &DataSourceDetailsFormatter{
		pbBodys: pbBodys,
	}
}

func NewListObjectFormatter(objectInfoSlice []minio.ObjectInfo) *ListObjectFormatter {
	return &ListObjectFormatter{
		objectInfoSlice: objectInfoSlice,
	}
}

type ListObjectFormatter struct {
	objectInfoSlice []minio.ObjectInfo
}

func (i *ListObjectFormatter) Header() []string {
	return []string{
		"key",
		"modified_time",
		"size",
	}
}

func (i *ListObjectFormatter) Rows() [][]string {
	var vals [][]string
	for _, objectInfo := range i.objectInfoSlice {
		vals = append(vals, []string{
			objectInfo.Key,
			formatTime(objectInfo.LastModified),
			formatSizeInBytes(objectInfo.Size),
		})
	}
	return vals
}

func (i *ListObjectFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(i).Write(w)
}

// CreatedApiKeyFormatter prints a freshly minted api key.
//
// The token is on its own row, labelled, because this is the ONLY time it is
// ever printed: the server returns it once at creation and the listing
// endpoint deliberately cannot return it. Someone who does not notice it here
// has to revoke the key and mint another.
type CreatedApiKeyFormatter struct {
	key   *appswaggermodel.AppkafeidoAPIKeyInfo
	token string
}

func (a *CreatedApiKeyFormatter) Header() []string {
	return []string{"key_id", "name", "role", "expires_at", "token"}
}

func (a *CreatedApiKeyFormatter) Rows() [][]string {
	var keyId, name, role, expiresAt string
	if a.key != nil {
		keyId = a.key.KeyID
		name = a.key.Name
		role = apiKeyRoleString(a.key.Role)
		expiresAt = a.key.ExpiresAt
		if len(expiresAt) == 0 {
			expiresAt = "never"
		}
	}
	return [][]string{
		{keyId, name, role, expiresAt, a.token},
	}
}

func (a *CreatedApiKeyFormatter) Write(w io.Writer) error {
	if err := NewDefaultFormatter(a).Write(w); err != nil {
		return err
	}
	// The table formatter truncates every cell to --lint_length (20 by
	// default), so the table alone prints "gtk_0123456789abcdef..." - an
	// unusable credential for the one command that can ever show it. Repeat it
	// in full underneath, in table form only: json and csv are not linted, and
	// an extra line there would break the parser an integration is piping this
	// into.
	if DefaultTypeFormatter == TypeFormatTable {
		if _, err := fmt.Fprintf(w, "token: %s\n", a.token); err != nil {
			return err
		}
	}
	return nil
}

func NewCreatedApiKeyFormatter(key *appswaggermodel.AppkafeidoAPIKeyInfo, token string) *CreatedApiKeyFormatter {
	return &CreatedApiKeyFormatter{key: key, token: token}
}

// ListApiKeyDetailsFormatter lists keys without their secrets, which the
// server cannot return in the first place.
type ListApiKeyDetailsFormatter struct {
	keys []*appswaggermodel.AppkafeidoAPIKeyInfo
}

func (a *ListApiKeyDetailsFormatter) Header() []string {
	return []string{
		"key_id",
		"name",
		"role",
		"created_at",
		"created_by",
		"expires_at",
		"last_used_at",
		"revoked_at",
	}
}

func (a *ListApiKeyDetailsFormatter) Rows() [][]string {
	var vals [][]string
	for _, key := range a.keys {
		if key == nil {
			continue
		}
		expiresAt := key.ExpiresAt
		if len(expiresAt) == 0 {
			expiresAt = "never"
		}
		// An empty last_used_at means a key that has never been presented,
		// which is worth reading as such when deciding whether to revoke it.
		lastUsedAt := key.LastUsedAt
		if len(lastUsedAt) == 0 {
			lastUsedAt = "never"
		}
		vals = append(vals, []string{
			key.KeyID,
			key.Name,
			apiKeyRoleString(key.Role),
			key.CreatedAt,
			key.CreatedBy,
			expiresAt,
			lastUsedAt,
			key.RevokedAt,
		})
	}
	return vals
}

func (a *ListApiKeyDetailsFormatter) Write(w io.Writer) error {
	return NewDefaultFormatter(a).Write(w)
}

func NewListApiKeyDetailsFormatter(keys []*appswaggermodel.AppkafeidoAPIKeyInfo) *ListApiKeyDetailsFormatter {
	return &ListApiKeyDetailsFormatter{keys: keys}
}

// apiKeyRoleString renders the enum as the word a user types into --role,
// rather than as API_KEY_ROLE_READWRITE.
func apiKeyRoleString(role *appswaggermodel.KafeidocommonpbAPIKeyRole) string {
	if role == nil {
		return ""
	}
	switch *role {
	case appswaggermodel.KafeidocommonpbAPIKeyRoleAPIKEYROLEREAD:
		return "read"
	case appswaggermodel.KafeidocommonpbAPIKeyRoleAPIKEYROLEREADWRITE:
		return "readwrite"
	case appswaggermodel.KafeidocommonpbAPIKeyRoleAPIKEYROLEADMIN:
		return "admin"
	default:
		return string(*role)
	}
}
