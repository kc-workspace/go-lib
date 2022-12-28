package plugins

import (
	"github.com/kc-workspace/go-lib/commandline/flags"
	"github.com/kc-workspace/go-lib/mapper"
)

// SupportCluster will create clusters when --clusters is exist
func SupportCluster(p *PluginParameter) error {
	p.NewFlags(flags.Array{
		Name:    "clusters",
		Default: []string{""},
		Usage:   "setup output clusters",
		Action: func(data []string) mapper.Mapper {
			if len(data) > 0 {
				return mapper.New().Set("clusters", data)
			}
			return mapper.New()
		},
	})

	return nil
}
