package cmd

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/log"

	appservice "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client/kafeido_service"
	appmodels "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/models"
	"github.com/footprintai/grandturks-client/v2/pkg/encryption"
)

func NewBasicLoginCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		email       string
		password    string
		save2Config bool
	)

	handler := func() error {
		runCmd := mustNewRunCmd(logger)
		params := &appservice.KafeidoServiceAppLoginParams{
			Body: &appmodels.KafeidoAppLoginRequest{
				Basic: &appmodels.KafeidoAppBasicLoginRequest{
					Email:    email,
					Password: password,
				},
			},
		}
		kafeidoAppLoginOk, err := runCmd.stub.KafeidoService.KafeidoServiceAppLogin(
			params.WithTimeout(runCmd.requestTimeout),
			nil,
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		rawAccessToken := kafeidoAppLoginOk.Payload.Basic.Oauth2Token.AccessToken
		populateAccessTokenViperConfig(rawAccessToken)
		if err := runCmd.createUser(rawAccessToken); err != nil {
			return err
		}
		if save2Config {
			return viper.WriteConfig()
		}
		// TODO: create user
		return nil
	}

	cmd := &cobra.Command{
		Use:   "basic",
		Short: "login with email/password, this can bypass oauth2 flow",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().StringVar(&email, "email", "", "email address for login")
	cmd.Flags().StringVar(&password, "password", "", "password")
	cmd.Flags().BoolVar(&save2Config, "save_to_config", true, "save information to configfile")

	return cmd
}

func populateAccessTokenViperConfig(rawAccessToken string) {
	ConfigKeyAuthToken.Set(rawAccessToken)
}

func populateUserInfo2ViperConfig(userInfo *appmodels.UserstoreUserInfo) {
	ConfigKeyUserEmail.Set(userInfo.UserEmail)
	ConfigKeyUserGroups.Set(userInfo.Groups)
	ConfigKeyUserId.Set(userInfo.UserID)
}

func NewOauth2LoginCommand(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command {
	var (
		save2Config bool
	)

	handler := func() error {
		runCmd := mustNewRunCmd(logger)

		currentRequestId := uuid.NewString()
		done := make(chan struct{})
		testserver := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				done <- struct{}{}
			}()
			credentials := r.URL.Query().Get("credentials")
			queryParams, err := encryption.NewEncryption(GetEncryptor()).DecodeStr(credentials)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Oops! something wrong, details:%+v\n", err)))
				return
			}
			values, err := url.ParseQuery(string(queryParams))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Oops! bad parameters\n"))
				return
			}
			if currentRequestId != values.Get("reqId") {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Oops! bad parameters\n"))
				return
			}
			timestamp, _ := strconv.Atoi(values.Get("timestamp"))
			if time.Now().Sub(time.Unix(int64(timestamp), 0)) > time.Hour {
				w.WriteHeader(http.StatusRequestTimeout)
				w.Write([]byte("Oops! too old data, please try again\n"))
				return
			}
			token := values.Get("token")
			populateAccessTokenViperConfig(token)
			if err := runCmd.createUser(token); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(fmt.Sprintf("Oops! login failed, please try again, details:%+v\n", err)))
				return
			}
			if save2Config {
				viper.WriteConfig()
			}
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("Logined. You can close this windows now."))
		}))

		params := &appservice.KafeidoServiceAppLoginParams{
			Body: &appmodels.KafeidoAppLoginRequest{
				Oauth2: &appmodels.KafeidoAppOauth2LoginRequest{
					LocalRedirectURL: testserver.URL,
					RequestID:        currentRequestId,
				},
			},
		}
		kafeidoAppLoginOk, err := runCmd.stub.KafeidoService.KafeidoServiceAppLogin(
			params.WithTimeout(runCmd.requestTimeout),
			runCmd.authInformer(),
		)
		if err != nil {
			return openapiErrorParser(err)
		}
		fmt.Printf("Open the following url on browser: %s\n", kafeidoAppLoginOk.Payload.Oauth2.OpenBrowserURL)

		select {
		case <-time.After(1000 * time.Second):
			panic("timeout")
		case <-done:
			// wait for 5 seconds before we close
			time.Sleep(5 * time.Second)
			return nil
		}
	}

	cmd := &cobra.Command{
		Use:   "login",
		Short: "oauth2 login flow to kafeido",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			cmd.Parent().PersistentPreRun(cmd.Parent(), args)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return handler()
		},
	}

	cmd.Flags().BoolVar(&save2Config, "save_to_config", true, "save information to configfile")

	cmd.AddCommand(NewBasicLoginCommand(logger, ioStreams))
	return cmd
}

func (r *RunCmd) createUser(rawAccessToken string) error {
	params := &appservice.KafeidoServiceCreateUserParams{
		Body: &appmodels.KafeidoCreateUserRequest{
			Oauth2Token: &appmodels.KafeidoOauth2Token{
				AccessToken: rawAccessToken,
			},
		},
	}
	kafeidoCreateUserOk, err := r.stub.KafeidoService.KafeidoServiceCreateUser(
		params.WithTimeout(r.requestTimeout),
		r.authInformer(),
	)
	if err != nil {
		return openapiErrorParser(err)
	}
	populateUserInfo2ViperConfig(kafeidoCreateUserOk.Payload.UserInfo)
	return nil
}
