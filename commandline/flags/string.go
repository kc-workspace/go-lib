package flags

import (
	"flag"

	"github.com/kc-workspace/go-lib/mapper"
)

type String struct {
	Name    string
	Default string
	Usage   string
	Action  func(data string) mapper.Mapper
}

func (f String) FlagName() string {
	return f.Name
}

func (f String) Parse(flag *flag.FlagSet) interface{} {
	var result = new(string)
	flag.StringVar(result, f.Name, f.Default, f.Usage)
	return result
}

func (f String) Build(value interface{}) mapper.Mapper {
	return f.Action(*value.(*string))
}
