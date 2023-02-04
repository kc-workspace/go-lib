package plugins

import (
	"fmt"

	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportHelp will add --help, -h, and help command to show help message
func SupportHelp() Plugin {
	return func(p *PluginParameter) error {
		p.NewFlags(flags.Bool{
			Name:    "help",
			Aliases: []string{"h"},
			Default: false,
			Usage:   "show application help",
			Action: func(data bool) mapper.Mapper {
				var m = mapper.New()
				if data {
					return m.Set("internal.command", "help")
				}
				return m
			},
		})

		p.NewHook(hooks.INTERNAL_HELP, func(config mapper.Mapper) error {
			var command = config.Mi("internal").Zi("command").(*commands.Manager)
			var flag = config.Mi("internal").Zi("flag").(*flags.Manager)

			var helpMessage = fmt.Sprintf(
				`Usage of %s:
%s
Commands:
%s
Options:
%s`,
				p.Metadata.Name,
				p.Metadata.Usage,
				command.Help(0),
				flag.Help(0),
			)

			config.Set("internal.helpMessage", helpMessage)
			return nil
		})

		p.NewCommand(&commands.Command{
			Name:  "help",
			Usage: "show application help",
			Executor: func(p *commands.ExecutorParameter) error {
				p.Logger.Log(p.Config.Mi("internal").Sr("helpMessage"))
				return nil
			},
		})

		return nil
	}
}
