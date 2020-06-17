package cli

import (
	"fmt"
	"net"
	"time"

	giconfig "github.com/b2wdigital/goignite/config"
	gilog "github.com/b2wdigital/goignite/log"
	"github.com/spf13/cobra"
)

type Cli struct {
	name        string
	description string
}

func New(name string, description string) *Cli {
	return &Cli{name: name, description: description}
}

func (h *Cli) Run() error {

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

func (h *Cli) parseFlag(cmd *cobra.Command, c giconfig.Config) { // nolint

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
