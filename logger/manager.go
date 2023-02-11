package logger

import (
	"github.com/kc-workspace/go-lib/utils"
)

type Manager struct {
	names   []string
	level   Level
	printer *Printer
}

func (m *Manager) SetLevel(level Level) *Manager {
	m.level = level
	return m
}

func (m *Manager) New(names ...string) *Logger {
	return NewLogger(utils.CloneArray(m.names, names...), m.level, m.printer)
}

func (m *Manager) NewTable(size uint) *Table {
	return NewTable(size, m.level, m.printer)
}
