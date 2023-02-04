package flags

import (
	"flag"
	"fmt"
	"strings"

	"github.com/kc-workspace/go-lib/logger"
	"github.com/kc-workspace/go-lib/mapper"
)

type Manager struct {
	keys   []string
	flags  map[string]Flag
	logger *logger.Logger
}

func (m *Manager) Add(flag Flag) {
	m.keys = append(m.keys, flag.FlagName())
	m.flags[flag.FlagName()] = flag
}

func (m *Manager) Size() int {
	return len(m.keys)
}

func (m *Manager) Build(name string, args []string) (map[string]mapper.Mapper, []string, error) {
	m.logger.Debug("building %d flags", m.Size())

	var option = mapper.New()
	// parse option flags
	var flagSet = flag.NewFlagSet(name, flag.ExitOnError)
	for key, builder := range m.flags {
		option.Set(key, builder.Parse(flagSet))
	}

	flagSet.Usage = func() {
		// Default flag when custom --help not provided
		m.logger.Log(fmt.Sprintf(
			"Usage of %s\n%s",
			name,
			m.Help(0),
		))
	}

	m.logger.Debug("parsing flagset")
	err := flagSet.Parse(args)
	if err != nil {
		return make(map[string]mapper.Mapper), args, err
	}

	// build result from flags
	var arguments = flagSet.Args()
	var result = make(map[string]mapper.Mapper)
	for key, value := range option {
		m.logger.Debug("building flag %s='%v'", key, value)
		var mapper = m.flags[key].Build(value)
		if !mapper.IsEmpty() {
			result[key] = mapper
		}
	}

	return result, arguments, err
}

func (m *Manager) Help(level int) string {
	const tab = "  "
	const newline = "\n"

	var indent = tab
	var maxNameSize = 12
	var maxLabelSize = 8
	for _, key := range m.keys {
		var flag = m.flags[key]

		var nameSize = len(flag.FlagName())
		if maxNameSize < nameSize {
			maxNameSize = nameSize
		}
		var labelSize = len(flag.FlagValueLabel())
		if maxLabelSize < labelSize {
			maxLabelSize = labelSize
		}
	}

	for i := 0; i < level; i++ {
		indent = indent + tab
	}

	var format = fmt.Sprintf(
		`%s--%%-%ds %%-%ds %s%s%%s%%s%s`,
		indent,
		maxNameSize,
		maxLabelSize,
		newline,
		indent+tab,
		newline,
	)

	var builder = strings.Builder{}
	for _, key := range m.keys {
		var flag = m.flags[key]

		var alias = ""
		var flagAliases = flag.FlagAliases()
		if len(flagAliases) > 0 {
			var aliases = make([]string, len(flagAliases))
			for i, alias := range flagAliases {
				if len(alias) == 1 {
					aliases[i] = fmt.Sprintf("-%s", alias)
				} else {
					aliases[i] = fmt.Sprintf("--%s", alias)
				}
			}

			alias = fmt.Sprintf(" %v", aliases)
		}

		var usage = " <no-description>"
		if flag.FlagUsage() != "" {
			usage = fmt.Sprintf(" %s", flag.FlagUsage())
		}

		builder.WriteString(fmt.Sprintf(
			format,
			flag.FlagName(),
			flag.FlagValueLabel(),
			alias,
			usage,
		))
	}

	return builder.String()
}

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
