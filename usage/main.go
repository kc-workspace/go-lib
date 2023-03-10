package main

import (
	"log"
	"os"

	"github.com/kc-workspace/go-lib/commandline"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/commandline/plugins"
	"github.com/kc-workspace/go-lib/logger"
)

func main() {
	var err = commandline.New(&models.Metadata{}).
		Plugin(plugins.SupportVersion()).
		Plugin(plugins.SupportHelp()).
		Plugin(plugins.SupportConfig([]string{})).
		Plugin(plugins.SupportVar()).
		Plugin(plugins.SupportDotEnv(false)).
		Plugin(plugins.SupportLogLevel(logger.WARN)).
		Start(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
