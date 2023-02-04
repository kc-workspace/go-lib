package commandline

import (
	"github.com/kc-workspace/go-lib/caches"
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/commandline/plugins"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type cli struct {
	Metadata *models.Metadata
	flags    *flags.Manager // global flags
	commands *commands.Manager
	hooks    *hooks.Manager
	plugins  *plugins.Manager

	cache  *caches.Service
	logger *logger.Logger
}

func (c *cli) Flag(flag flags.Flag) *cli {
	c.flags.Add(flag)
	return c
}

func (c *cli) Command(cmd *commands.Command) *cli {
	c.commands.Add(cmd)
	return c
}

func (c *cli) Hook(t hooks.Type, hook hooks.Hook) *cli {
	c.hooks.Add(t, hook)
	return c
}

func (c *cli) Plugin(plugin plugins.Plugin) *cli {
	c.plugins.Add(plugin)
	return c
}

func (c *cli) Start(args []string) error {
	c.logger.Debug("starting %s command", c.Metadata.Name)

	var config = mapper.New()
	config.Set("internal.meta", c.Metadata.ToMapper())

	if err := c.hooks.Start(hooks.BEFORE_PLUGIN, config); err != nil {
		return err
	}

	var err = c.plugins.Build(&plugins.PluginParameter{
		Metadata:   *c.Metadata,
		NewCommand: c.commands.Add,
		NewFlags:   c.flags.Add,
		NewHook:    c.hooks.Add,
		Config:     config,
		Logger:     c.logger,
	})
	if err == nil {
		err = c.hooks.Start(hooks.AFTER_PLUGIN, config)
	}
	if err != nil {
		return err
	}

	var name string
	if c.Metadata != nil && c.Metadata.Name != "" {
		name = c.Metadata.Name
	} else {
		name = args[0]
	}

	option, parsed, err := c.flags.Build(name, args[1:])
	if err != nil {
		return err
	}
	for _, value := range option {
		config = mapper.Merger(config).Add(value).Merge()
	}
	if err = c.hooks.Start(hooks.AFTER_FLAG, config); err != nil {
		return err
	}

	var cmd, final = c.commands.Get(parsed, config)
	if len(final) > 0 {
		config.Set("internal.command", cmd.Name).Set("internal.args", final)
	}

	if err = c.hooks.Start(hooks.BEFORE_COMMAND, config); err != nil {
		return err
	}

	var cmderr = cmd.Start(&commands.ExecutorParameter{
		Name:   cmd.Name,
		Meta:   c.Metadata,
		Args:   final,
		Config: config,

		Cache:  c.cache,
		Logger: logger.Get("command", cmd.Name),
	})

	config.Set("internal.error", cmderr)
	if err = c.hooks.Start(hooks.AFTER_COMMAND, config); err != nil {
		return err
	}

	return cmderr
}
