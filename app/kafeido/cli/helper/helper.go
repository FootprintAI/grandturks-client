package clihelper

import (
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
)

type CliJobType string

func (c CliJobType) String() string {
	return string(c)
}

var (
	CliJobNone  CliJobType = "none"
	CliJobImage CliJobType = "image"
	CliJobAudio CliJobType = "audio"
)

func MakeOpenAPIModelInferenceJob(j CliJobType) *appmodels.TaskworkerInferenceJob {
	var taskJobKind appmodels.TaskworkerInferenceJob
	if j == CliJobImage {
		taskJobKind = appmodels.TaskworkerInferenceJobINFERENCEJOBIMAGE
	} else if j == CliJobAudio {
		taskJobKind = appmodels.TaskworkerInferenceJobINFERENCEJOBAUDIO
	}
	return &taskJobKind
}

func MakeCliJobType(j *appmodels.TaskworkerInferenceJob) CliJobType {
	if j == nil {
		return CliJobNone
	}

	if *j == appmodels.TaskworkerInferenceJobINFERENCEJOBIMAGE {
		return CliJobImage
	} else if *j == appmodels.TaskworkerInferenceJobINFERENCEJOBAUDIO {
		return CliJobAudio
	}
	return CliJobNone
}
