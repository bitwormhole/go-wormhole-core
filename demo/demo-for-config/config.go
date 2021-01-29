package config

import "github.com/bitwormhole/go-wormhole-core/application"

// Config config the app
func Config(config *application.Configuration) {

	config.Component(&application.ComponentRegistration{
		ID:     "car1",
		Class:  "car",
		Scope:  application.ScopeSingleton,
		Source: newCarRef,
	})

}
