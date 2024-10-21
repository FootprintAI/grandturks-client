package cmd

import (
	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
)

func getProjectById(cmd *RunCmd, projectId string) (*appservice.KafeidoServiceGetProjectOK, error) {
	params := &appservice.KafeidoServiceGetProjectParams{
		ProjectID: projectId,
	}
	return cmd.stub.KafeidoService.KafeidoServiceGetProject(
		params.WithTimeout(cmd.requestTimeout),
		cmd.authInformer(),
	)
}
