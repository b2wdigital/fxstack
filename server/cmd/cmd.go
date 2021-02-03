package cmd

import (
	"os"

	"github.com/b2wdigital/fxstack/cmd"
	"github.com/b2wdigital/fxstack/fx/server/cloudevents"
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
	c.SetAppModule(h.appModule)
	c.SetLogger(h.logger)

	c.Load()

	switch os.Getenv("CMD_DEFAULT") {
	case "server.cloudevents":
		return cloudevents.Start(h.handlerModule)
	default:
		return c.RunCobra()
	}

}

func (h *Cmd) innerCmd() []*cobra.Command {

	serverCmd := &cobra.Command{
		Use:   "server",
		Short: "server",
		Long:  "",
	}

	ceCmd := &cobra.Command{
		Use:   "cloudevents",
		Short: "cloudevents",
		Long:  "",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cloudevents.Start(h.handlerModule)
		},
	}

	serverCmd.AddCommand(ceCmd)

	return []*cobra.Command{serverCmd}
}
