package plugins

import "github.com/kc-workspace/go-lib/logger"

func New(l *logger.Logger) *Manager {
	return &Manager{
		plugins: make([]Plugin, 0),
		logger:  l.Extend("plugin"),
	}
}
