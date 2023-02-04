package hooks

import (
	"fmt"
	"strings"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type Manager struct {
	hooks  map[Type][]Hook
	logger *logger.Logger
}

func (m *Manager) Add(t Type, hook Hook) {
	m.hooks[t] = append(m.hooks[t], hook)
}

func (m *Manager) Size(t Type) int {
	return len(m.hooks[t])
}

func (m *Manager) Start(t Type, config mapper.Mapper) error {
	m.logger.Debug("starting %d %s hooks", m.Size(t), t)

	for _, hook := range m.hooks[t] {
		var err = hook(config)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) String() string {
	var str strings.Builder
	for key, value := range m.hooks {
		str.WriteString(fmt.Sprintf("%s: %d jobs\n", key, len(value)))
	}
	return str.String()
}
