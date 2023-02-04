package flags

import (
	"flag"

	"github.com/kc-workspace/go-lib/mapper"
)

type Flag interface {
	FlagName() string
	FlagValueLabel() string
	FlagAliases() []string
	FlagUsage() string

	Parse(flag *flag.FlagSet) interface{}
	Build(value interface{}) mapper.Mapper
}
