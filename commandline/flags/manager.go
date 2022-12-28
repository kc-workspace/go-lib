package flags

import (
	"flag"
	"fmt"

	"github.com/kc-workspace/go-lib/mapper"
)

type Manager struct {
	keys  []string
	flags map[string]Flag
}

func (f *Manager) Add(flag Flag) {
	f.keys = append(f.keys, flag.FlagName())
	f.flags[flag.FlagName()] = flag
}

func (f *Manager) Build(name string, args []string) (map[string]mapper.Mapper, []string, error) {
	var option = mapper.New()

	// parse option flags
	var flagSet = flag.NewFlagSet(name, flag.ExitOnError)
	for key, builder := range f.flags {
		option.Set(key, builder.Parse(flagSet))
	}
	err := flagSet.Parse(args)
	if err != nil {
		return make(map[string]mapper.Mapper), args, err
	}

	// build result from flags
	var arguments = flagSet.Args()
	var result = make(map[string]mapper.Mapper)
	for key, value := range option {
		var mapper = f.flags[key].Build(value)
		if !mapper.IsEmpty() {
			result[key] = mapper
		}
	}

	return result, arguments, err
}

func (m *Manager) String() string {
	return fmt.Sprintf("%v", m.keys)
}
