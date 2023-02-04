package plugins

import "github.com/kc-workspace/go-lib/logger"

func New() *Manager {
	return &Manager{
		plugins: make([]Plugin, 0),
		logger:  logger.Get("commandline", "plugin"),
	}
}
