package simple

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

type myPrototypeComAgent struct {
	context application.Context
	info    application.ComponentInfo
}

func (inst *myPrototypeComAgent) GetInstance() lang.Object {
	return inst.GetInstanceRef().GetInstance()
}

func (inst *myPrototypeComAgent) GetInstanceRef() application.ComponentInstanceRef {
	// TODO
	return nil
}

func (inst *myPrototypeComAgent) GetInfo() application.ComponentInfo {
	return inst.info
}

func (inst *myPrototypeComAgent) GetContext() application.Context {
	return inst.context
}

func (inst *myPrototypeComAgent) MakeChild(context application.Context) application.ComponentAgent {
	if context == nil {
		context = inst.context
	}
	return &myPrototypeComAgent{
		context: context,
		info:    inst.info,
	}
}
