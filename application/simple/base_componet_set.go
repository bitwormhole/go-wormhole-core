package simple

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// BaseComponentSet 是默认的组件集合 ,实现了 Components 接口
type BaseComponentSet struct {
	context application.Context
	table   map[string]application.ComponentAgent
}

func (inst *BaseComponentSet) Clear() {}

func (inst *BaseComponentSet) Export(dst map[string]application.ComponentAgent) map[string]application.ComponentAgent {
	return nil
}

func (inst *BaseComponentSet) Import(map[string]application.ComponentAgent) {

}

func (inst *BaseComponentSet) GetAgent(name string) (application.ComponentAgent, error) {
	return nil, nil
}

func (inst *BaseComponentSet) GetComponent(name string) (lang.Object, error) {
	return nil, nil
}

func (inst *BaseComponentSet) GetComponentByClass(classSelector string) (lang.Object, error) {
	return nil, nil
}

func (inst *BaseComponentSet) GetComponentsByClass(classSelector string) []lang.Object {
	return nil
}

func (inst *BaseComponentSet) SetAgent(name string, agent application.ComponentAgent) {

}

func NewComponentSet(context application.Context) application.Components {
	cs := &BaseComponentSet{}
	cs.context = context
	return cs
}
