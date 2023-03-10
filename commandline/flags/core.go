package flags

import "github.com/kc-workspace/go-lib/logger"

func New(l *logger.Logger, flags ...Flag) *Manager {
	var keys = make([]string, 0)
	var m = make(map[string]Flag)
	for _, flag := range flags {
		keys = append(keys, flag.FlagName())
		m[flag.FlagName()] = flag
	}

	return &Manager{
		keys:  keys,
		flags: m,

		logger: l.Extend("flag"),
	}
}
