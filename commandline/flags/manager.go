package flags

import (
	"flag"
	"fmt"

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

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
