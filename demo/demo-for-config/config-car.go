package config

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

func newCarRef() application.ComponentInstance {
	inst := &Car{}
	return &carRef{car: inst}
}

type carRef struct {
	car *Car
}

func (inst *carRef) Get() lang.Object {
	return inst.car
}

func (inst *carRef) IsLoaded() bool {
	return false
}

func (inst *carRef) Inject(context application.RuntimeContext) error {
	return nil
}

func (inst *carRef) Init() error {
	return inst.car.start()
}

func (inst *carRef) Destroy() error {
	return inst.car.stop()
}
