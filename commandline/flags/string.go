package flags

import (
	"flag"
	"fmt"

	"github.com/kc-workspace/go-lib/mapper"
)

type String struct {
	Name    string
	Aliases []string
	Default string
	Usage   string
	Action  func(data string) mapper.Mapper
}

func (f String) FlagName() string {
	return f.Name
}

func (f String) FlagValueLabel() string {
	return fmt.Sprintf("<str:%v>", f.Default)
}

func (f String) FlagAliases() []string {
	return f.Aliases
}

func (f String) FlagUsage() string {
	return f.Usage
}

func (f String) Parse(flag *flag.FlagSet) interface{} {
	var result = new(string)
	flag.StringVar(result, f.Name, f.Default, f.Usage)
	for _, alias := range f.Aliases {
		flag.StringVar(result, alias, f.Default, f.Usage)
	}

	return result
}

func (f String) Build(value interface{}) mapper.Mapper {
	return f.Action(*value.(*string))
}
