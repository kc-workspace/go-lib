package commandline

import (
	"github.com/kc-workspace/go-lib/caches"
	"github.com/kc-workspace/go-lib/commandline/commands"
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/commandline/hooks"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/commandline/plugins"
	"github.com/kc-workspace/go-lib/logger"
)

func New(metadata *models.Metadata) *cli {
	return &cli{
		Metadata: metadata,
		flags:    flags.New(),
		commands: commands.New(),
		hooks:    hooks.New(),
		plugins:  plugins.New(),

		cache:  caches.Global,
		logger: logger.Get("commandline"),
	}
}

func NewCustom(
	metadata *models.Metadata,
	cache *caches.Service,
	logger *logger.Logger,
) *cli {
	return &cli{
		Metadata: metadata,
		flags:    flags.New(),
		commands: commands.New(),
		hooks:    hooks.New(),
		plugins:  plugins.New(),

		cache:  cache,
		logger: logger,
	}
}
