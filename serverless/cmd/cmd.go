package cmd

import (
	"os"

	"github.com/b2wdigital/fxstack/cmd"
	"github.com/b2wdigital/fxstack/fx/listener/nats"
	"github.com/b2wdigital/fxstack/fx/serverless/lambda"
	gilog "github.com/b2wdigital/goignite/log"
	"go.uber.org/fx"

	"github.com/spf13/cobra"
)

type Cmd struct {
	name          string
	description   string
	handlerModule func() fx.Option
	appModule     func() fx.Option
	logger        func() gilog.Logger
}

func New(name string, description string,
	handlerModule func() fx.Option, appModule func() fx.Option, logger func() gilog.Logger) *Cmd {
	return &Cmd{name: name, description: description, handlerModule: handlerModule, appModule: appModule, logger: logger}
}

func (h *Cmd) SetLogger(logger func() gilog.Logger) {
	h.logger = logger
}

func (h *Cmd) Run() error {

	c := cmd.New(h.name, h.description)
	c.SetCmds(h.innerCmd())
	c.SetLogger(h.logger)
	c.SetAppModule(h.appModule)
	c.Load()

	switch os.Getenv("CMD_DEFAULT") {
	case "listener.nats":
		return nats.Start(h.handlerModule)
	case "serverless.lambda":
		return lambda.Start(h.handlerModule)
	default:
		return c.RunCobra()
	}

}

func (h *Cmd) innerCmd() []*cobra.Command {

	listenerCmd := &cobra.Command{
		Use:   "listener",
		Short: "listeners for messaging system",
		Long:  "",
	}

	natsCmd := &cobra.Command{
		Use:   "nats",
		Short: "listener for nats (http://nats.io)",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return nats.Start(h.handlerModule)
		},
	}

	listenerCmd.AddCommand(natsCmd)

	serverLessCmd := &cobra.Command{
		Use:   "serverless",
		Short: "integration for serverless plataforms",
		Long:  "",
	}

	lambdaCmd := &cobra.Command{
		Use:   "lambda",
		Short: "AWS Lambda",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return lambda.Start(h.handlerModule)
		},
	}

	serverLessCmd.AddCommand(lambdaCmd)

	return []*cobra.Command{listenerCmd, serverLessCmd}
}
