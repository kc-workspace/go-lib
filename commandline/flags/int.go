package flags

import (
	"flag"
	"fmt"

	"github.com/kc-workspace/go-lib/mapper"
)

type Int struct {
	Name    string
	Aliases []string
	Default int64
	Usage   string
	Action  func(data int64) mapper.Mapper
}

func (f Int) FlagName() string {
	return f.Name
}

func (f Int) FlagValueLabel() string {
	return fmt.Sprintf("<int:%v>", f.Default)
}

func (f Int) FlagAliases() []string {
	return f.Aliases
}

func (f Int) FlagUsage() string {
	return f.Usage
}

func (f Int) Parse(flag *flag.FlagSet) interface{} {
	var result = new(int64)
	flag.Int64Var(result, f.Name, f.Default, f.Usage)
	for _, alias := range f.Aliases {
		flag.Int64Var(result, alias, f.Default, f.Usage)
	}

	return result
}

func (f Int) Build(value interface{}) mapper.Mapper {
	return f.Action(*value.(*int64))
}
