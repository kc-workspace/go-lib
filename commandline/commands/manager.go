package commands

import (
	"fmt"
	"strings"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type Manager struct {
	keys     []string
	aliases  []string
	commands map[string]*Command
	logger   *logger.Logger
}

func (m *Manager) Add(cmd *Command) {
	if _, exist := m.commands[cmd.Name]; exist {
		// Do not duplicate commands
		return
	}

	m.keys = append(m.keys, cmd.Name)
	// add command name to command mapping
	m.commands[cmd.Name] = cmd
	// add command alias to command mapping
	for _, alias := range cmd.Aliases {
		m.aliases = append(m.aliases, alias)
		m.commands[alias] = cmd
	}
}

func (m *Manager) Size() int {
	return len(m.keys)
}

func (m *Manager) Get(args []string, config mapper.Mapper) (*Command, []string) {
	m.logger.Debug("finding command from %d commands list", m.Size())

	var commands = append([]string{}, m.keys...)
	commands = append(commands, m.aliases...)
	var name, parsed = getName(commands, args, config)
	m.logger.Debug("found command %s with %v", name, parsed)

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

func (m *Manager) Help(level int) string {
	var tab = "  "
	var size = 10
	for _, key := range m.keys {
		var length = len(key)
		if length > size {
			size = length
		}
	}

	for i := 0; i < level; i++ {
		tab = tab + tab
	}

	var format = fmt.Sprintf(`%s- %%-%ds :%%s%%s`+"\n", tab, size)

	var builder = strings.Builder{}
	for _, key := range m.keys {
		var cmd = m.commands[key]

		var alias = ""
		if len(cmd.Aliases) > 0 {
			alias = fmt.Sprintf(" %v", cmd.Aliases)
		}

		var usage = " <no-description>"
		if cmd.Usage != "" {
			usage = fmt.Sprintf(" %s", cmd.Usage)
		}

		builder.WriteString(
			fmt.Sprintf(format, cmd.Name, alias, usage),
		)

		if cmd.Flags != nil {
			builder.WriteString(
				cmd.Flags.Help(1),
			)
		}
	}

	return builder.String()
}

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
