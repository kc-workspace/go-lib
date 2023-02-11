package logger

import "github.com/kc-workspace/go-lib/utils"

func GetName(names []string) string {
	return utils.JoinString(":", names...)
}

func ShouldPrint(setting Level, level Level) bool {
	return setting >= level
}
