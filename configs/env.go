package configs

import (
	"github.com/kc-workspace/go-lib/mapper"
)

// Return map of string, client will decide how to parse string data
func ParseConfigFromEnv(prefix string, environments []string) (mapper.Mapper, error) {
	var result = mapper.New()
	for _, env := range environments {
		if k, v, ok := ParseOverride(env); ok {
			if key, ok := EnvToKey(prefix, k); ok {
				result.Set(key, v)
			}
		}
	}

	return result, nil
}
