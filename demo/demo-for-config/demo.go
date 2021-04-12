package config

import (
	"os"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"
)

func Demo() int {

	config := &config.AppConfig{}
	Config(config)

	context, err := application.Run(config, os.Args)
	if err != nil {
		panic(err)
	}

	code := application.Exit(context)
	return code
}
