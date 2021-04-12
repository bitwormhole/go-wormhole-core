package config

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/config"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// Config config the app
func Config(cfg application.ConfigBuilder) {

	cfg.AddComponent(&config.ComInfo{
		ID:    "x1",
		Class: "x",

		OnInject: nil,
	})

	cfg.AddComponent(&config.ComInfo{
		ID:    "car1",
		Class: "car",
		Scope: application.ScopeSingleton,

		OnNew: func() lang.Object {
			return &Car{}
		},

		OnInit: func(obj lang.Object) error {
			car := obj.(Car)
			return car.start()
		},

		OnInject: func(obj lang.Object, context application.RuntimeContext) error {

			car := obj.(Car)
			getter := context.NewGetter(nil)
			ec := getter.ErrorCollector()

			car.driver = getter.GetComponent("driver").(*Driver)
			ec.AddErrorIfNil(car.driver, "")

			return ec.Result()
		},
	})

}
