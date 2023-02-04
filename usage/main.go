package main

import (
	"log"
	"os"

	"github.com/kc-workspace/go-lib/commandline"
	"github.com/kc-workspace/go-lib/commandline/models"
	"github.com/kc-workspace/go-lib/commandline/plugins"
)

func main() {
	var err = commandline.New(&models.Metadata{}).
		Plugin(plugins.SupportVersion()).
		Plugin(plugins.SupportHelp()).
		Plugin(plugins.SupportConfig("gol", []string{})).
		Plugin(plugins.SupportVar()).
		Plugin(plugins.SupportDotEnv()).
		Plugin(plugins.SupportLogLevel()).
		Start(os.Args)

	if err != nil {
		log.Fatal(err)
	}
}
