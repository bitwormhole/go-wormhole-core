package config

import (
	"github.com/bitwormhole/go-wormhole-core/application"
)

// Car class
type Car struct {
	id      string
	context application.Context

	driver *Driver
	engine *Engine

	wheelFrontLeft  *Wheel
	wheelFrontRight *Wheel
	wheelBackLeft   *Wheel
	wheelBackRight  *Wheel
}

func (inst *Car) start() error {
	return nil
}

func (inst *Car) stop() error {
	return nil
}
