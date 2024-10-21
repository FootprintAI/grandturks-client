package cmd

import (
	goflag "flag"
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/mitchellh/go-homedir"
	commonhttp "github.com/sdinsure/agent/pkg/http"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"k8s.io/cli-runtime/pkg/genericclioptions"
	"sigs.k8s.io/kind/pkg/cmd"
	"sigs.k8s.io/kind/pkg/log"

	swaggerclient "github.com/footprintai/grandturks-client/v2/api/app/kafeido/proto/go-openapiv2/client"
	"github.com/footprintai/grandturks-client/v2/app/kafeido/cli/format"
	openapierrors "github.com/footprintai/grandturks-client/v2/pkg/http/openapi/errors"
)

var (
	cfgFile      string
	debug        bool
	outputFormat string
	lintLength   int
)

// NewCommandFunc new and returns a *cobra.Command
type NewCommandFunc func(logger log.Logger, ioStreams genericclioptions.IOStreams) *cobra.Command

type PluggableCommand interface {
	apply(*plugableCommand)
}

type plugableCommand struct {
	createCommands []NewCommandFunc
	getCommands    []NewCommandFunc
	listCommands   []NewCommandFunc
	deleteCommands []NewCommandFunc
	putCommands    []NewCommandFunc
	runCommands    []NewCommandFunc
	cancelCommands []NewCommandFunc
	uploadCommands []NewCommandFunc
}

func newPlugableCommands(pcs ...PluggableCommand) *plugableCommand {
	pcConfig := &plugableCommand{}
	for _, pc := range pcs {
		pc.apply(pcConfig)
	}
	return pcConfig
}

type createCommands struct {
	cmds []NewCommandFunc
}

func (c createCommands) apply(pc *plugableCommand) {
	pc.createCommands = append(pc.createCommands, c.cmds...)
}

func WithCreateCommands(cmds ...NewCommandFunc) createCommands {
	return createCommands{cmds: cmds}
}

type getCommands struct {
	cmds []NewCommandFunc
}

func (c getCommands) apply(pc *plugableCommand) {
	pc.getCommands = append(pc.getCommands, c.cmds...)
}

func WithGetCommands(cmds ...NewCommandFunc) getCommands {
	return getCommands{cmds: cmds}
}

type listCommands struct {
	cmds []NewCommandFunc
}

func (c listCommands) apply(pc *plugableCommand) {
	pc.listCommands = append(pc.listCommands, c.cmds...)
}

func WithListCommands(cmds ...NewCommandFunc) listCommands {
	return listCommands{cmds: cmds}
}

type deleteCommands struct {
	cmds []NewCommandFunc
}

func (c deleteCommands) apply(pc *plugableCommand) {
	pc.deleteCommands = append(pc.deleteCommands, c.cmds...)
}

func WithDeleteCommands(cmds ...NewCommandFunc) deleteCommands {
	return deleteCommands{cmds: cmds}
}

type putCommands struct {
	cmds []NewCommandFunc
}

func (c putCommands) apply(pc *plugableCommand) {
	pc.putCommands = append(pc.putCommands, c.cmds...)
}

func WithPutCommands(cmds ...NewCommandFunc) putCommands {
	return putCommands{cmds: cmds}
}

type runCommands struct {
	cmds []NewCommandFunc
}

func (c runCommands) apply(pc *plugableCommand) {
	pc.runCommands = append(pc.runCommands, c.cmds...)
}

func WithRunCommands(cmds ...NewCommandFunc) runCommands {
	return runCommands{cmds: cmds}
}

type cancelCommands struct {
	cmds []NewCommandFunc
}

func (c cancelCommands) apply(pc *plugableCommand) {
	pc.cancelCommands = append(pc.cancelCommands, c.cmds...)
}

func WithCancelCommands(cmds ...NewCommandFunc) cancelCommands {
	return cancelCommands{cmds: cmds}
}

type uploadCommands struct {
	cmds []NewCommandFunc
}

func (c uploadCommands) apply(pc *plugableCommand) {
	pc.uploadCommands = append(pc.uploadCommands, c.cmds...)
}

func WithUploadCommands(cmds ...NewCommandFunc) uploadCommands {
	return uploadCommands{cmds: cmds}
}

func NewRootCommand(logger log.Logger, ioStreams genericclioptions.IOStreams, pcs ...PluggableCommand) *cobra.Command {
	plugableCommandOptions := newPlugableCommands(pcs...)

	cmd := &cobra.Command{
		Use:   "kafiedo",
		Short: "a kafiedo cli tool",
		Long:  `kafiedo is a command-line tool.`,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			// For cobra + glog flags. Available to all subcommands.
			goflag.Parse()
			populate2Goflags()
			populate2DefaultVars()
		},
	}
	cmd.AddCommand(NewVersionCommand(logger, ioStreams))
	cmd.AddCommand(NewOauth2LoginCommand(logger, ioStreams))
	cmd.AddCommand(NewConfigCommand(logger, ioStreams))
	cmd.AddCommand(NewCreateCommand(logger, ioStreams, plugableCommandOptions.createCommands...))
	cmd.AddCommand(NewGetCommand(logger, ioStreams, plugableCommandOptions.getCommands...))
	cmd.AddCommand(NewListCommand(logger, ioStreams, plugableCommandOptions.listCommands...))
	cmd.AddCommand(NewDeleteCommand(logger, ioStreams, plugableCommandOptions.deleteCommands...))
	cmd.AddCommand(NewPutCommand(logger, ioStreams, plugableCommandOptions.putCommands...))
	cmd.AddCommand(NewRunCommand(logger, ioStreams, plugableCommandOptions.runCommands...))
	cmd.AddCommand(NewUploadCommand(logger, ioStreams, plugableCommandOptions.uploadCommands...))
	cmd.AddCommand(NewCancelCommand(logger, ioStreams, plugableCommandOptions.cancelCommands...))

	cmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.kafeidoconfig)")
	cmd.PersistentFlags().BoolVar(&debug, "debug", false, "debug flag (default: false)")
	cmd.PersistentFlags().StringVar(&outputFormat, "format", "table", "output format(table, csv, json), (default: table)")
	cmd.PersistentFlags().IntVar(&lintLength, "lint_length", 30, "lint length for each variable (default: 30 chars)")
	return cmd
}

func OpenAPIErrorParser(err error) error {
	return openapiErrorParser(err)
}

func openapiErrorParser(err error) error {
	return openapierrors.Parse(err, debug)
}

func populate2Goflags() {
	//httpflags.HTTPDebug = nullable.BoolP(debug)
}

func populate2DefaultVars() {
	format.DefaultTypeFormatter = format.MustParseFormat(outputFormat)
	format.DefaultTableLintLength = lintLength
}

// Execute executes the root command.
func Execute() error {
	return NewRootCommand(
		cmd.NewLogger(),
		genericclioptions.IOStreams{
			In:     os.Stdin,
			Out:    os.Stdout,
			ErrOut: os.Stderr,
		},
	).Execute()
}

func init() {
	cobra.OnInitialize(initConfig)
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.SetConfigType("json")
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			os.Exit(1)
		}
		// Search config in home directory with name ".kafeidoconfig" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("json")
		viper.SetConfigName(".kafeidoconfig")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
	} else {
		viper.SafeWriteConfig()
	}
}

func MustNewRunCmd(logger log.Logger) *RunCmd {
	return mustNewRunCmd(logger)
}

func mustNewRunCmd(logger log.Logger) *RunCmd {
	cmd, err := newRunCmd(logger)
	if err != nil {
		panic(err)
	}
	return cmd
}

func newRunCmd(logger log.Logger) (*RunCmd, error) {
	hostUrl, err := url.Parse(ConfigKeyApiEndpoint.GetString())
	if err != nil {
		return nil, err
	}
	stub := swaggerclient.New(httptransport.NewWithClient(
		hostUrl.Host,
		filepath.Join("api", swaggerclient.DefaultBasePath),
		[]string{hostUrl.Scheme},
		commonhttp.NewHttpClient(commonhttp.WithDebug(debug, newPkgLoggerAdaptor(logger))),
	), nil)
	authInformer := func() runtime.ClientAuthInfoWriter {
		return runtime.ClientAuthInfoWriterFunc(func(clientRequest runtime.ClientRequest, registry strfmt.Registry) error {
			authToken := ConfigKeyAuthToken.GetString()
			if len(authToken) > 0 {
				return clientRequest.SetHeaderParam("Authorization", fmt.Sprintf("Bearer %s", authToken))
			}
			return nil
		})
	}
	r := &RunCmd{
		stub:           stub,
		requestTimeout: ConfigKeyRequestTimeout.GetDuration(),
		authInformer:   authInformer,
	}
	return r, nil
}

type RunCmd struct {
	stub           *swaggerclient.GrandturkKafeidoAPIDocumentations
	requestTimeout time.Duration
	authInformer   func() runtime.ClientAuthInfoWriter
}

func (r *RunCmd) Stub() *swaggerclient.GrandturkKafeidoAPIDocumentations {
	return r.stub
}

func (r *RunCmd) AuthInformer() runtime.ClientAuthInfoWriter {
	return r.authInformer()

}

func (r *RunCmd) RequestTimeout() time.Duration {
	return r.requestTimeout
}
