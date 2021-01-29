package config

import (
	"testing"

	"github.com/bitwormhole/go-wormhole-core/application"
)

func TestDemo(t *testing.T) {

	config := &application.Configuration{}
	Config(config)

	context, err := application.Run(config)
	if err != nil {
		t.Error(err)
		return
	}

	code := application.Exit(context)
	t.Logf("exit with code: %d", code)
}
