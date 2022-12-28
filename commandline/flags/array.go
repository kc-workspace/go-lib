package flags

import (
	"flag"

	"github.com/kc-workspace/go-lib/mapper"
)

type arrayValue []string

func (i *arrayValue) String() string {
	return "array"
}

func (i *arrayValue) Set(value string) error {
	*i = append(*i, value)
	return nil
}

type Array struct {
	Name    string
	Default []string
	Usage   string
	Action  func(data []string) mapper.Mapper
}

func (f Array) FlagName() string {
	return f.Name
}

func (f Array) Parse(flag *flag.FlagSet) interface{} {
	var value = new(arrayValue)
	flag.Var(value, f.Name, f.Usage)
	return value
}

func (f Array) Build(value interface{}) mapper.Mapper {
	var data = f.Default
	var arr = value.(*arrayValue)
	if *arr != nil {
		data = *arr
	}
	return f.Action(data)
}
