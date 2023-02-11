package commands

import "github.com/kc-workspace/go-lib/logger"

func New(l *logger.Logger) *Manager {
	return &Manager{
		keys:     make([]string, 0),
		aliases:  make([]string, 0),
		commands: make(map[string]*Command),
		logger:   l.Extend("command"),
	}
}
