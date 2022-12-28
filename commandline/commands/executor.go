package commands

import (
	"github.com/kc-workspace/go-lib/caches"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type ExecutorParameter struct {
	Name   string
	Meta   *models.Metadata
	Config mapper.Mapper
	Cache  *caches.Service
	Logger *logger.Logger
	Args   []string
}

type Executor func(p *ExecutorParameter) error
