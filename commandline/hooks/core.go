package hooks

import "github.com/kc-workspace/go-lib/logger"

func New(l *logger.Logger) *Manager {
	return &Manager{
		hooks:  make(map[Type][]Hook),
		logger: l.Extend("hooks"),
	}
}
