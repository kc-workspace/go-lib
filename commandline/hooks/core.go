package hooks

import "github.com/kc-workspace/go-lib/logger"

func New() *Manager {
	return &Manager{
		hooks:  make(map[Type][]Hook),
		logger: logger.Get("commandline", "hooks"),
	}
}
