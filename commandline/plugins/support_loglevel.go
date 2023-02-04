package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

func SupportLogLevel(defaultLevel logger.Level) Plugin {
	return func(p *PluginParameter) error {
		p.NewFlags(flags.Int{
			Name:    "log-level",
			Aliases: []string{"l"},
			Default: int64(defaultLevel),
			Usage:   "setup log level; 0 is silent and 4 is verbose",
			Action: func(data int64) mapper.Mapper {
				return mapper.New().Set("internal.log.level", data)
			},
		})

		p.NewFlags(flags.Bool{
			Name:    "debug",
			Aliases: []string{"D"},
			Default: false,
			Usage:   "mark current log to debug mode",
			Action: func(data bool) mapper.Mapper {
				var m = mapper.New()
				if data {
					logger.SetLevel(logger.DEBUG) // force if --debug is exist
					return m.Set("internal.log.level", 4)
				}
				return m
			},
		})

		p.NewHook(hooks.BEFORE_COMMAND, func(config mapper.Mapper) error {
			var level, err = config.Mi("internal").Mi("log").Ne("level")
			if err != nil {
				return err
			}

			logger.SetLevel(level)
			return nil
		})

		return nil
	}
}
