package simple

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// ComponentAliasAgent 表示组件的别名
type ComponentAliasAgent struct {
	targetName  string
	aliasName   string
	context     application.Context
	targetAgent application.ComponentAgent
}

func (inst *ComponentAliasAgent) getTarget() application.ComponentAgent {
	ta := inst.targetAgent
	if ta == nil {
		id := inst.targetName
		ta2, _ := inst.context.GetComponents().GetAgent(id)
		ta = ta2
	}
	return ta
}

func (inst *ComponentAliasAgent) GetInstance() lang.Object {
	return inst.getTarget().GetInstance()
}

func (inst *ComponentAliasAgent) GetInstanceRef() application.ComponentInstanceRef {
	return inst.getTarget().GetInstanceRef()
}

func (inst *ComponentAliasAgent) GetInfo() application.ComponentInfo {
	return inst.getTarget().GetInfo()
}

func (inst *ComponentAliasAgent) GetContext() application.Context {
	return inst.context
}

func (inst *ComponentAliasAgent) MakeChild(context application.Context) application.ComponentAgent {
	child := &ComponentAliasAgent{}
	if context == nil {
		context = inst.context
	}
	child.aliasName = inst.aliasName
	child.targetName = inst.targetName
	child.context = context
	child.targetAgent = nil
	return child
}
