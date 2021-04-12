package config

import (
	"testing"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"
)

func TestDemo(t *testing.T) {

	config := &config.AppConfig{}
	Config(config)

	context, err := application.Run(config)
	if err != nil {
		t.Error(err)
		return
	}

	code := application.Exit(context)
	t.Logf("exit with code: %d", code)
}
