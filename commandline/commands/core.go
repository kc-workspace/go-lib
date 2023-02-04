package commands

import "github.com/kc-workspace/go-lib/logger"

func New() *Manager {
	return &Manager{
		keys:     make([]string, 0),
		commands: make(map[string]*Command),
		logger:   logger.Get("commandline", "command"),
	}
}
