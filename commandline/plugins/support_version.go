package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/mapper"
)

func SupportVersion(p *PluginParameter) error {
	p.NewFlags(flags.Bool{
		Name:    "version",
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
		Name: "version",
		Executor: func(p *commands.ExecutorParameter) error {
			p.Logger.Log(p.Meta.String())
			return nil
		},
	})

	return nil
}
