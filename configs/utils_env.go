package configs

import (
	"fmt"
	"path"
	"strings"
)

func BuildEnvPrefix(prefix string) (output string) {
	if prefix != "" {
		// If prefix is empty string, path.Base will convert to '.'
		output = path.Base(prefix)
	}

	output = strings.ToUpper(output)
	output = strings.ReplaceAll(output, ".", "__")
	output = strings.ReplaceAll(output, "-", "_")

	return
}

func IsEnvKey(prefix, k string) bool {
	return strings.HasPrefix(k, prefix+"_")
}

func IsCEnvKey(prefix, k string) bool {
	return strings.HasPrefix(k, prefix+CUSTOM_PREFIX+"_")
}

func EnvToKey(prefix, c string) (string, bool) {
	if (!IsCEnvKey(prefix, c) && !IsEnvKey(prefix, c)) || strings.Contains(c, "___") {
		return "", false
	}

	var output string = ""
	if IsCEnvKey(prefix, c) {
		output = strings.Replace(c, prefix+CUSTOM_PREFIX+"_", "", 1)
	} else if IsEnvKey(prefix, c) {
		output = strings.Replace(c, prefix+"_", "", 1)
	}

	output = strings.ReplaceAll(output, "__", ".")
	output = strings.ReplaceAll(output, "_", "-")
	if IsCEnvKey(prefix, c) {
		output = "_." + output
	}

	return strings.ToLower(output), true
}

// NOTE: ou should not use _ as a key, since we cannot parse underscroll back after it created environment variable
// Please use - (dash) instead
func KeyToEnv(prefix, key string) string {
	if key[0] == '_' {
		var result = fmt.Sprintf("%s-%s", prefix+CUSTOM_PREFIX, key[2:])
		return strings.ToUpper(
			strings.ReplaceAll(strings.ReplaceAll(result, ".", "__"), "-", "_"),
		)
	} else {
		var result = fmt.Sprintf("%s-%s", prefix, key)
		return strings.ToUpper(
			strings.ReplaceAll(strings.ReplaceAll(result, ".", "__"), "-", "_"),
		)
	}
}
