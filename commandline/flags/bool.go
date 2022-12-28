package flags

import (
	"flag"

	"github.com/kc-workspace/go-lib/mapper"
)

type Bool struct {
	Name    string
	Default bool
	Usage   string
	Action  func(data bool) mapper.Mapper
}

func (f Bool) FlagName() string {
	return f.Name
}

func (f Bool) Parse(flag *flag.FlagSet) interface{} {
	var result = new(bool)
	flag.BoolVar(result, f.Name, f.Default, f.Usage)
	return result
}

func (f Bool) Build(value interface{}) mapper.Mapper {
	return f.Action(*value.(*bool))
}
