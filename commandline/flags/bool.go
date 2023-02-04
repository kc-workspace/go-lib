package flags

import (
	"flag"
	"fmt"

	"github.com/kc-workspace/go-lib/mapper"
)

type Bool struct {
	Name    string
	Aliases []string
	Default bool
	Usage   string
	Action  func(data bool) mapper.Mapper
}

func (f Bool) FlagName() string {
	return f.Name
}

func (f Bool) FlagValueLabel() string {
	return fmt.Sprintf("<bool:%v>", f.Default)
}

func (f Bool) FlagAliases() []string {
	return f.Aliases
}

func (f Bool) FlagUsage() string {
	return f.Usage
}

func (f Bool) Parse(flag *flag.FlagSet) interface{} {
	var result = new(bool)
	flag.BoolVar(result, f.Name, f.Default, f.Usage)
	for _, alias := range f.Aliases {
		flag.BoolVar(result, alias, f.Default, f.Usage)
	}

	return result
}

func (f Bool) Build(value interface{}) mapper.Mapper {
	return f.Action(*value.(*bool))
}
