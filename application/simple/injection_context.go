package simple

import (
	"errors"
	"fmt"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

type myInjectionTarget struct {
	agent application.ComponentAgent
	ref   application.ComponentInstanceRef

	hasInject  bool
	hasInit    bool
	hasDestroy bool
}

////////////////////////////////////////////////////////////////////////////////

type myInjectionContextCore struct {
	target application.Context
	buffer map[string]myInjectionTarget
}

func (inst *myInjectionContextCore) init(ctx application.Context) {

	inst.target = ctx
	inst.buffer = make(map[string]myInjectionTarget)

}

func (inst *myInjectionContextCore) GetArguments() collection.Arguments {
	return inst.target.GetArguments()
}

func (inst *myInjectionContextCore) GetAttributes() collection.Attributes {
	return inst.target.GetAttributes()
}

func (inst *myInjectionContextCore) GetParameters() collection.Parameters {
	return inst.target.GetParameters()
}

func (inst *myInjectionContextCore) GetProperties() collection.Properties {
	return inst.target.GetProperties()
}

func (inst *myInjectionContextCore) GetEnvironment() collection.Environment {
	return inst.target.GetEnvironment()
}

func (inst *myInjectionContextCore) GetReleasePool() collection.ReleasePool {
	return inst.target.GetReleasePool()
}

func (inst *myInjectionContextCore) GetResources() collection.Resources {
	return inst.target.GetResources()
}

func (inst *myInjectionContextCore) NewGetter(ec lang.ErrorCollector) application.ContextGetter {
	return application.NewGetter(inst)
}

func (inst *myInjectionContextCore) GetComponents() application.Components {
	return inst
}

func (inst *myInjectionContextCore) Clear() {

}

func (inst *myInjectionContextCore) GetComponentByClass(selector string) (lang.Object, error) {
	list := inst.GetComponentsByClass(selector)
	size := len(list)
	if size == 1 {
		item := list[0]
		if item != nil {
			return item, nil
		}
	}
	if size > 1 {
		msg := fmt.Sprint("there are multiple components with class: ", selector)
		return nil, errors.New(msg)
	}
	msg := fmt.Sprint("no component with class: ", selector)
	return nil, errors.New(msg)
}

func (inst *myInjectionContextCore) GetComponentsByClass(selector string) []lang.Object {
	return nil
}

func (inst *myInjectionContextCore) GetComponent(name string) (lang.Object, error) {
	return nil, nil
}

func (inst *myInjectionContextCore) GetAgent(name string) (application.ComponentAgent, error) {
	// NOP
	return nil, nil
}

func (inst *myInjectionContextCore) SetAgent(name string, agent application.ComponentAgent) {
	// NOP
}

func (inst *myInjectionContextCore) Import(table map[string]application.ComponentAgent) {
	// NOP
}

func (inst *myInjectionContextCore) Export(table map[string]application.ComponentAgent) map[string]application.ComponentAgent {
	// NOP
	return table
}

////////////////////////////////////////////////////////////////////////////////

type myInjectionContextFacade struct {
	current    application.Context
	configtime application.Context
	runtime    application.Context
}

func (inst *myInjectionContextFacade) init(runtimeContext application.Context) error {
	configTime := &myInjectionContextCore{}
	configTime.init(runtimeContext)

	inst.configtime = configTime
	inst.current = configTime
	inst.runtime = runtimeContext
	return nil
}

func (inst *myInjectionContextFacade) start() error {
	inst.current = inst.runtime
	inst.configtime = nil
	return nil
}

func (inst *myInjectionContextFacade) stop() error {
	return nil
}

func (inst *myInjectionContextFacade) GetArguments() collection.Arguments {
	return inst.current.GetArguments()
}

func (inst *myInjectionContextFacade) GetAttributes() collection.Attributes {
	return inst.current.GetAttributes()
}

func (inst *myInjectionContextFacade) GetParameters() collection.Parameters {
	return inst.current.GetParameters()
}

func (inst *myInjectionContextFacade) GetProperties() collection.Properties {
	return inst.current.GetProperties()
}

func (inst *myInjectionContextFacade) GetEnvironment() collection.Environment {
	return inst.current.GetEnvironment()
}

func (inst *myInjectionContextFacade) GetComponents() application.Components {
	return inst.current.GetComponents()
}

func (inst *myInjectionContextFacade) GetReleasePool() collection.ReleasePool {
	return inst.current.GetReleasePool()
}

func (inst *myInjectionContextFacade) GetResources() collection.Resources {
	return inst.current.GetResources()
}

func (inst *myInjectionContextFacade) NewGetter(ec lang.ErrorCollector) application.ContextGetter {
	return inst.current.NewGetter(ec)
}

////////////////////////////////////////////////////////////////////////////////
// EOF
