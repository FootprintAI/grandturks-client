package cmd

import (
	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido"
)

func getProjectById(cmd *RunCmd, projectId string) (*appservice.KafeidoGetProjectOK, error) {
	params := &appservice.KafeidoGetProjectParams{
		ProjectID: projectId,
	}
	return cmd.stub.Kafeido.KafeidoGetProject(
		params.WithTimeout(cmd.requestTimeout),
		cmd.authInformer(),
	)
}
