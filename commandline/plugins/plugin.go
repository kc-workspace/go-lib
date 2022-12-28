package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type PluginParameter struct {
	Metadata models.Metadata

	NewCommand commands.Creator
	NewFlags   flags.Creator
	NewHook    hooks.Creator

	Config mapper.Mapper
	Logger *logger.Logger
}

type Plugin func(p *PluginParameter) error
