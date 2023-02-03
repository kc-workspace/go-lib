package plugins

import (
	"fmt"

	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/configs"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportVar will add --var "<name>=<value>" for assign data to variables.*
func SupportVar(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "var",
		Default: []string{},
		Usage:   "add data to variables config",
		Action: func(data []string) mapper.Mapper {
			var m = mapper.New()
			for _, d := range data {
				var key, value, ok = configs.ParseOverride(d)
				if ok {
					m.Set(fmt.Sprintf("variables.%s", key), value)
				}
			}
			return m
		},
	})
	return nil
}
