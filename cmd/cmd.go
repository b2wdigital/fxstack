package cmd

import (
	"context"
	"fmt"
	"net"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	gizerolog "github.com/b2wdigital/goignite/log/zerolog/v1"
	"github.com/spf13/cobra"
	_ "go.uber.org/automaxprocs"
	"go.uber.org/fx"
)

type Cmd struct {
	name        string
	description string
	cmd         []*cobra.Command
	logger      func() gilog.Logger
	appModule   func() fx.Option
}

func New(name string, description string) *Cmd {
	return &Cmd{name: name, description: description}
}

func (h *Cmd) SetCmds(cmds []*cobra.Command) {
	h.cmd = cmds
}

func (h *Cmd) SetLogger(logger func() gilog.Logger) {
	h.logger = logger
}

func (h *Cmd) SetAppModule(module func() fx.Option) {
	h.appModule = module
}

func (h *Cmd) Load() {

	var logger func() gilog.Logger

	if h.logger != nil {
		logger = h.logger
	} else {
		logger = gizerolog.NewLogger
	}

	fx.New(
		fx.Invoke(
			giconfig.Load,
			logger,
		),
	).Start(context.Background())

	if h.appModule != nil {
		fx.New(h.appModule()).
			Start(context.Background())
	}

}

func (h *Cmd) RunCobra() error {

	// cobra.OnInitialize(h.Load)

	rootCmd := &cobra.Command{
		Use:   h.name,
		Short: h.description,
		Long:  "",
	}

	rootCmd.AddCommand(h.cmd...)

	rootCmd.DisableFlagParsing = true

	for _, entry := range giconfig.Entries() {
		h.parseFlag(rootCmd, entry)
	}

	rootCmd.PersistentFlags().StringSlice(giconfig.ConfArgument, nil, "path to one or more config files")

	return rootCmd.Execute()
}

func (h *Cmd) parseFlag(cmd *cobra.Command, c giconfig.Config) { // nolint

	switch t := c.Value.(type) {

	case string:
		cmd.PersistentFlags().String(c.Key, t, c.Description)
	case []string:
		cmd.PersistentFlags().StringSlice(c.Key, t, c.Description)
	case map[string]string:
		var s string
		for key, val := range t {
			s = s + fmt.Sprintf("%s=\"%s\" ", key, val)
		}
		cmd.PersistentFlags().String(c.Key, s, c.Description)
	case bool:
		cmd.PersistentFlags().Bool(c.Key, t, c.Description)
	case []bool:
		cmd.PersistentFlags().BoolSlice(c.Key, t, c.Description)
	case int:
		cmd.PersistentFlags().Int(c.Key, t, c.Description)
	case []int:
		cmd.PersistentFlags().IntSlice(c.Key, t, c.Description)
	case int8:
		cmd.PersistentFlags().Int8(c.Key, t, c.Description)
	case int16:
		cmd.PersistentFlags().Int16(c.Key, t, c.Description)
	case int32:
		cmd.PersistentFlags().Int32(c.Key, t, c.Description)
	case int64:
		cmd.PersistentFlags().Int64(c.Key, t, c.Description)
	case uint:
		cmd.PersistentFlags().Uint(c.Key, t, c.Description)
	case uint64:
		cmd.PersistentFlags().Uint64(c.Key, t, c.Description)
	case time.Duration:
		cmd.PersistentFlags().Duration(c.Key, t, c.Description)
	case []byte:
		cmd.PersistentFlags().BytesBase64(c.Key, t, c.Description)
	case float64:
		cmd.PersistentFlags().Float64(c.Key, t, c.Description)
	case net.IPNet:
		cmd.PersistentFlags().IPNet(c.Key, t, c.Description)
	case net.IP:
		cmd.PersistentFlags().IP(c.Key, t, c.Description)
	case net.IPMask:
		cmd.PersistentFlags().IPMask(c.Key, t, c.Description)
	default:
		gilog.Warnf("type unknown %s", t)
	}

}
