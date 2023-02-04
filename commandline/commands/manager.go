package commands

import (
	"fmt"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type Manager struct {
	keys     []string
	commands map[string]*Command
	logger   *logger.Logger
}

func (m *Manager) Add(cmd *Command) {
	m.keys = append(m.keys, cmd.Name)
	m.commands[cmd.Name] = cmd
}

func (m *Manager) Size() int {
	return len(m.keys)
}

func (m *Manager) Get(args []string, config mapper.Mapper) (*Command, []string) {
	m.logger.Debug("finding command from %d commands list", m.Size())

	var name, parsed = getName(m.keys, args, config)
	m.logger.Debug("parsed command %s: %v", name, parsed)

	var cmd = m.commands[name]
	if cmd == nil {
		return &Command{
			Name: DEFAULT,
			Executor: func(parameters *ExecutorParameter) error {
				parameters.Logger.ErrorString("You didn't specify command name to run")
				parameters.Logger.ErrorString("Supported commands: %v", m.keys)
				return nil
			},
		}, parsed
	}

	return cmd, parsed
}

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
