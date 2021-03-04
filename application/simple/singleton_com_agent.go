package simple

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

type mySingletonComAgent struct {
	info    application.ComponentInfo
	context application.Context
	comRef  application.ComponentInstanceRef
}

func (inst *mySingletonComAgent) GetInstance() lang.Object {
	return inst.GetInstanceRef().GetInstance()
}

func (inst *mySingletonComAgent) GetInstanceRef() application.ComponentInstanceRef {
	ref := inst.comRef
	if ref == nil {
		ref = inst.info.GetFactory().NewInstance()
		inst.comRef = ref
	}
	return ref
}

func (inst *mySingletonComAgent) GetInfo() application.ComponentInfo {
	return inst.info
}

func (inst *mySingletonComAgent) GetContext() application.Context {
	return inst.context
}

func (inst *mySingletonComAgent) MakeChild(context application.Context) application.ComponentAgent {
	return inst
}
