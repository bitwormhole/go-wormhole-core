package simple

import "github.com/bitwormhole/go-wormhole-core/application"

type RuntimeComInfo struct {
	id      string
	class   string
	aliases []string
	scope   application.ComponentScope
	factory application.ComponentFactory
}

func (inst *RuntimeComInfo) GetID() string {
	return inst.id
}

func (inst *RuntimeComInfo) GetClass() string {
	return inst.class
}

func (inst *RuntimeComInfo) GetAliases() []string {
	return inst.aliases
}

func (inst *RuntimeComInfo) GetScope() application.ComponentScope {
	return inst.scope
}

func (inst *RuntimeComInfo) GetFactory() application.ComponentFactory {
	return inst.factory
}

func (inst *RuntimeComInfo) xxx() {

}
