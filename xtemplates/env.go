package xtemplates

import (
	"fmt"
	"text/template"

	"github.com/kc-workspace/go-lib/mapper"
)

func byCluster(config mapper.Mapper, clusterName string, configPath string) (interface{}, error) {
	var path = fmt.Sprintf("_.%s.%s", clusterName, configPath)
	if config.Has(path) {
		return config.Get(path)
	}

	// fallback to get default value
	return config.Get(configPath)
}

var envFuncs template.FuncMap = map[string]interface{}{
	"byCluster": byCluster,
}
