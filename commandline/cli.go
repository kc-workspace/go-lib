package commandline

import (
	"fmt"

	"github.com/kc-workspace/go-lib/caches"
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/commandline/plugins"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type Commandline struct {
	Metadata *models.Metadata
	flags    *flags.Manager // global flags
	commands *commands.Manager
	hooks    *hooks.Manager
	plugins  *plugins.Manager

	cache  *caches.Service
	logger *logger.Logger
}

func (c *Commandline) Flag(flag flags.Flag) *Commandline {
	c.flags.Add(flag)
	return c
}

func (c *Commandline) Command(cmd *commands.Command) *Commandline {
	c.commands.Add(cmd)
	return c
}

func (c *Commandline) Hook(t hooks.Type, hook hooks.Hook) *Commandline {
	c.hooks.Add(t, hook)
	return c
}

func (c *Commandline) Plugin(plugin plugins.Plugin) *Commandline {
	c.plugins.Add(plugin)
	return c
}

func (c *Commandline) Start(args []string) error {
	if c.Metadata.Name == "" {
		c.Metadata.Name = args[0]
	}

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

	option, parsed, err := c.flags.Build(c.Metadata.Name, args[1:])
	if err != nil {
		return err
	}
	for _, value := range option {
		config = mapper.Merger(config).Add(value).Merge()
	}
	if err = c.hooks.Start(hooks.AFTER_FLAG, config); err != nil {
		return err
	}

	// get command object and command args
	var cmd, final = c.commands.Get(parsed, config)
	if len(final) > 0 {
		config.Set("internal.command", cmd.Name).Set("internal.args", final)
	}

	// special hooks for help command
	if err = c.hooks.Start(
		hooks.INTERNAL_HELP,
		config.
			Set("internal.command", c.commands).
			Set("internal.flag", c.flags),
	); err != nil {
		return err
	}

	// Delete special keys for internal-help only
	config.
		Del("internal.command").
		Del("internal.flag")

	if err = c.hooks.Start(hooks.BEFORE_COMMAND, config); err != nil {
		return err
	}

	c.logger.Debug("executing command %s with %v", cmd.Name, final)
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

func (c *Commandline) String() string {
	return fmt.Sprintf("%s commandline {plugin=%d, command=%d, flag=%d}",
		c.Metadata.Name,
		c.plugins.Size(),
		c.commands.Size(),
		c.flags.Size(),
	)
}
