package plugins

import (
	"fmt"
	"os"
	"sort"
	"strings"

	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/configs"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportConfig will add several things listed below:
// 1. --pwd for resolve directory (default is current directory)
// 2. --configs for specify configuration files or directories
// 3. command 'config' for list config
// For parameters:
// defaultPrefix is environment prefix
// defaultConfig is default configs if no flag provided
func SupportConfig(defaultPrefix string, defaultConfig []string) Plugin {
	return func(p *PluginParameter) error {
		var wd, err = os.Getwd()
		if err != nil {
			return err
		}

		p.NewFlags(flags.String{
			Name:    "pwd",
			Default: wd,
			Usage:   "current directory",
			Action: func(data string) mapper.Mapper {
				return mapper.New().Set("variables.current", data)
			},
		})

		p.NewFlags(flags.Array{
			Name:    "configs",
			Default: defaultConfig,
			Usage:   "configuration file/directory. both must be either json or yaml",
			Action: func(data []string) mapper.Mapper {
				var result = mapper.New()
				if len(data) > 0 {
					result.
						Set("internal.fs.config.type", "auto").
						Set("internal.fs.config.mode", "multiple").
						Set("internal.fs.config.fullpath", data)
				}
				return result
			},
		})

		p.NewHook(hooks.BEFORE_COMMAND, func(config mapper.Mapper) error {
			var name = p.Metadata.Name
			if name == "" {
				name = defaultPrefix
			}

			var addition, err = configs.New(name, config).Build(os.Environ())
			if err != nil {
				return err
			}

			// load additional config from files
			addition.ForEach(func(key string, value interface{}) {
				config.Set(key, value)
			})

			return nil
		})

		p.NewCommand(&commands.Command{
			Name:  "config",
			Usage: "list all possible config user can set",
			Flags: flags.New(flags.Bool{
				Name:    "data",
				Default: false,
				Usage:   "show config value as well",
				Action: func(data bool) mapper.Mapper {
					return mapper.New().Set("internal.flag.data", data)
				},
			}, flags.Bool{
				Name:    "all",
				Default: false,
				Usage:   "show all configuration, including internal",
				Action: func(data bool) mapper.Mapper {
					return mapper.New().Set("internal.flag.all", data)
				},
			}),
			Executor: func(p *commands.ExecutorParameter) error {
				var withData = p.Config.Mi("internal").Mi("flag").Bo("data", false)
				var all = p.Config.Mi("internal").Mi("flag").Bo("all", false)
				var prefix = p.Meta.Name
				if prefix == "" {
					prefix = defaultPrefix
				}

				var headers = []string{"Key", "Environment"}
				if withData {
					headers = append(headers, "Type", "Value")
				}

				var table = logger.GetTable(uint(len(headers)))
				table.Row(headers...)

				var keys = p.Config.Keys()
				sort.Strings(keys)

				// Sorted keys
				for _, key := range keys {
					if !all && strings.HasPrefix(key, "internal") {
						continue
					}

					var env = configs.KeyToEnv(prefix, key)
					var row = []string{key, env}
					if withData {
						var value, _ = p.Config.Get(key)
						var t = fmt.Sprintf("%T", value)
						row = append(row, t, fmt.Sprintf("%v", value))
					}

					table.Row(row...)
				}

				return table.End()
			},
		})

		return nil
	}
}
