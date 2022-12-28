package configs

import (
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

func New(name string, config mapper.Mapper) *Builder {
	return &Builder{
		name:     name,
		config:   config,
		override: mapper.New(),
		strategy: mapper.New(),

		logger: logger.Get("config", "builder"),
	}
}
