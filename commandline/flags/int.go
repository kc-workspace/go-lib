package flags

import (
	"flag"

	"github.com/kc-workspace/go-lib/mapper"
)

type Int struct {
	Name    string
	Default int64
	Usage   string
	Action  func(data int64) mapper.Mapper
}

func (f Int) FlagName() string {
	return f.Name
}

func (f Int) Parse(flag *flag.FlagSet) interface{} {
	var result = new(int64)
	flag.Int64Var(result, f.Name, f.Default, f.Usage)
	return result
}

func (f Int) Build(value interface{}) mapper.Mapper {
	return f.Action(*value.(*int64))
}
