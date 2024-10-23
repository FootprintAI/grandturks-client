package cmd

import (
	"errors"
	"strings"

	"github.com/spf13/cobra"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
)

func NewCreateProjectCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectName                   string
		projectDesc                   string
		preferredProjectId            string
		preferredAnnonymousBucketName string
		preferredPrivateBucketName    string
		projectShortVisibility        string
	)

	handler := func() error {
		pbVisibility, err := convertShortVisibilityToFull(projectShortVisibility)
		if err != nil {
			return err
		}

		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceCreateProjectParams{
			Body: &appmodels.KafeidoCreateProjectRequest{
				Name:                       projectName,
				Desc:                       projectDesc,
				Visibility:                 pbVisibility.Pointer(),
				PreferredProjectID:         preferredProjectId,
				PreferredPrivateBucketName: preferredPrivateBucketName,
				PreferredPublicBucketName:  preferredAnnonymousBucketName,
			},
		}
		kafeidoCreateProjectOk, err := runCmd.stub.KafeidoService.KafeidoServiceCreateProject(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(kafeidoCreateProjectOk.Payload.ProjectID, "Project is created").Write(ioStreams.Out)

	}

	cmd := &cobra.Command{
		Use:   "project",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectName, "name", "", "project name")
	cmd.Flags().StringVar(&projectDesc, "desc", "", "project description")
	cmd.Flags().StringVar(&preferredProjectId, "project_id", "", "preferred project id")
	cmd.Flags().StringVar(&preferredAnnonymousBucketName, "public_bucket_name", "", "preferred public bucket name")
	cmd.Flags().StringVar(&preferredPrivateBucketName, "private_bucket_name", "", "preferred private bucket name")
	cmd.Flags().StringVar(&projectShortVisibility, "visibility", "team", "permission of the project, one of private, team, public, default: team")
	cmd.MarkFlagRequired("name")
	cmd.MarkFlagRequired("desc")

	return cmd
}

func convertShortVisibilityToFull(s string) (appmodels.AuthzTypeVisibility, error) {
	switch strings.ToLower(s) {
	case "owner":
		return appmodels.AuthzTypeVisibilityTYPEVISIBILITYONLYOWNER, nil
	case "team":
		return appmodels.AuthzTypeVisibilityTYPEVISIBILITYTEAM, nil
	case "public":
		return appmodels.AuthzTypeVisibilityTYPEVISIBILITYPUBLIC, nil
	default:
		return appmodels.AuthzTypeVisibilityTYPEVISIBILITYUNSPECIFIED, errors.New("not possible visibility")
	}
}

func NewGetProjectCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		kafeidoGetProjectOk, err := getProjectById(runCmd, projectId)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewListProjectDetailsFormatter([]*appmodels.KafeidoGetProjectResponse{kafeidoGetProjectOk.Payload}).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "project",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")

	return cmd
}

func NewListProjectCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceListProjectsParams{}

		kafeidoListProjectsOk, err := runCmd.stub.KafeidoService.KafeidoServiceListProjects(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}

		var listResp []*appmodels.KafeidoGetProjectResponse
		for _, projectId := range kafeidoListProjectsOk.Payload.ProjectIds {
			subParams := &appservice.KafeidoServiceGetProjectParams{
				ProjectID: projectId,
			}
			kafeidoGetProjectOk, err := runCmd.stub.KafeidoService.KafeidoServiceGetProject(
				subParams.WithTimeout(runCmd.requestTimeout),
				runCmd.authInformer(),
			)
			if err == nil {
				listResp = append(listResp, kafeidoGetProjectOk.Payload)
			}
		}

		return format.NewListProjectDetailsFormatter(listResp).Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "project",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	return cmd
}

func NewDeleteProjectCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		projectId string
	)
	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceDeleteProjectParams{
			ProjectID: projectId,
		}
		_, err := runCmd.stub.KafeidoService.KafeidoServiceDeleteProject(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		return format.NewIdAndMessageFormatter(projectId, "Project is deleted").Write(ioStreams.Out)
	}

	cmd := &cobra.Command{
		Use:   "project",
		Short: "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&projectId, "project_id", "", "project id")
	cmd.MarkFlagRequired("project_id")

	return cmd
}
