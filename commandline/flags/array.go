package flags

import (
	"flag"
	"fmt"

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
	Aliases []string
	Default []string
	Usage   string
	Action  func(data []string) mapper.Mapper
}

func (f Array) FlagName() string {
	return f.Name
}

func (f Array) FlagValueLabel() string {
	return fmt.Sprintf("<%v>", f.Default)
}

func (f Array) FlagAliases() []string {
	return f.Aliases
}

func (f Array) FlagUsage() string {
	return f.Usage
}

func (f Array) Parse(flag *flag.FlagSet) interface{} {
	var value = new(arrayValue)
	flag.Var(value, f.Name, f.Usage)
	for _, alias := range f.Aliases {
		flag.Var(value, alias, f.Usage)
	}

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
