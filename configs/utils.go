package configs

import (
	"strings"

	"github.com/kc-workspace/go-lib/mapper"
)

func ParseOverride(str string) (key, value string, ok bool) {
	ok = false
	var arr = strings.Split(str, "=")
	if len(arr) != 2 {
		return
	}

	ok = true
	key = arr[0]
	value = arr[1]
	return
}

func BuildClusterConfig(cluster string, config mapper.Mapper) mapper.Mapper {
	if cluster == "" {
		return config
	}

	// We check cluster as case-insensitive
	// TODO: move this code to when we write cluster key in map, it will reduce significant number of calculate
	var underscroll = config.Mi("_")
	var c1 = underscroll.Mi(cluster)
	var c2 = underscroll.Mi(strings.ToLower(cluster))
	var c3 = underscroll.Mi(strings.ToUpper(cluster))

	return mapper.Merger(config).Add(c1).Add(c2).Add(c3).Merge()
}
