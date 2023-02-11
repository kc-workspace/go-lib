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

func New(metadata *models.Metadata) *Commandline {
	return NewCustom(
		metadata,
		caches.Global,
		logger.DefaultManager.New("commandline"),
	)
}

func NewCustom(
	metadata *models.Metadata,
	cache *caches.Service,
	logger *logger.Logger,
) *Commandline {
	return &Commandline{
		Metadata: metadata,
		flags:    flags.New(logger),
		commands: commands.New(logger),
		hooks:    hooks.New(logger),
		plugins:  plugins.New(logger),

		cache:  cache,
		logger: logger,
	}
}
