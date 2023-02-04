package plugins

import (
	"fmt"

	"github.com/kc-workspace/go-lib/logger"
)

type Manager struct {
	plugins []Plugin
	logger  *logger.Logger
}

func (m *Manager) Add(plugin Plugin) *Manager {
	m.plugins = append(m.plugins, plugin)
	return m
}

func (m *Manager) Size() int {
	return len(m.plugins)
}

func (m *Manager) Build(parameters *PluginParameter) error {
	m.logger.Debug("building %d plugins", m.Size())

	for _, plugin := range m.plugins {
		var err = plugin(parameters)
		if err != nil {
			return err
		}
	}
	return nil
}

func (m *Manager) String() string {
	return fmt.Sprintf("Manager %d plugins", len(m.plugins))
}
