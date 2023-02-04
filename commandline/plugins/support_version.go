package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportVersion will add '--version', '-v', 'version' to get metadata version
func SupportVersion() Plugin {
	return func(p *PluginParameter) error {
		p.NewFlags(flags.Bool{
			Name:    "version",
			Aliases: []string{"v"},
			Default: false,
			Usage:   "show current application version",
			Action: func(data bool) mapper.Mapper {
				var m = mapper.New()
				if data {
					return m.Set("internal.command", "version")
				}
				return m
			},
		})

		p.NewCommand(&commands.Command{
			Name:  "version",
			Usage: "show current application version",
			Executor: func(p *commands.ExecutorParameter) error {
				p.Logger.Log(p.Meta.String())
				return nil
			},
		})

		return nil
	}
}
