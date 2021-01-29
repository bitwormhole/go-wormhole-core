package config

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

func newCarRef() application.ComponentInstanceRef {
	inst := &Car{}
	return &carRef{car: inst}
}

type carRef struct {
	car *Car
}

func (inst *carRef) GetInstance() lang.Object {
	return inst.car
}

func (inst *carRef) Inject(context application.Context) error {
	return nil
}

func (inst *carRef) Init() error {
	return inst.car.start()
}

func (inst *carRef) Destroy() error {
	return inst.car.stop()
}
