package flags

import "github.com/kc-workspace/go-lib/logger"

func New(flags ...Flag) *Manager {
	var m = make(map[string]Flag)
	for _, flag := range flags {
		m[flag.FlagName()] = flag
	}

	return &Manager{
		keys:  make([]string, 0),
		flags: m,

		logger: logger.Get("commandline", "flag"),
	}
}
