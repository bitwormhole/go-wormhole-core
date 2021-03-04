package config

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/application/simple"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// Config config the app
func Config(config application.ConfigBuilder) {

	config.AddComponent(&simple.ComInfo{
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

		OnInject: func(obj lang.Object, context application.Context) error {

			car := obj.(Car)
			getter := context.NewGetter(nil)
			ec := getter.ErrorCollector()

			car.driver = getter.GetComponent("driver").(*Driver)
			ec.AddErrorIfNil(car.driver, "")

			return ec.Result()
		},
	})

}
