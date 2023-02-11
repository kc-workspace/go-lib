package logger

import "github.com/kc-workspace/go-lib/utils"

func GetName(names []string) string {
	return utils.JoinString(":", names...)
}
