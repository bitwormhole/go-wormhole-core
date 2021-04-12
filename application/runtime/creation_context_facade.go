package runtime

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type creationContextFacade struct {
	// impl CreationContext & RuntimeContext
	core *creationContextCore
}

type creationComponentsFacade struct {
	core *creationContextCore
}

////////////////////////////////////////////////////////////////////////////////
// impl creationContextFacade

func (inst *creationContextFacade) getParentRC() application.RuntimeContext {
	return inst.core.parent.context
}

func (inst *creationContextFacade) GetURI() string {
	return inst.getParentRC().GetURI()
}

func (inst *creationContextFacade) GetApplicationName() string {
	return inst.getParentRC().GetApplicationName()
}

func (inst *creationContextFacade) GetApplicationVersion() string {
	return inst.getParentRC().GetApplicationVersion()
}

func (inst *creationContextFacade) GetStartupTimestamp() int64 {
	return inst.getParentRC().GetStartupTimestamp()
}

func (inst *creationContextFacade) GetShutdownTimestamp() int64 {
	return inst.getParentRC().GetShutdownTimestamp()
}

func (inst *creationContextFacade) NewChild() application.RuntimeContext {
	return inst.getParentRC().NewChild()
}

func (inst *creationContextFacade) GetComponents() application.Components {
	return inst.core.components
}

func (inst *creationContextFacade) NewGetter(ec lang.ErrorCollector) application.ContextGetter {
	// todo ...
	return nil
}

func (inst *creationContextFacade) GetReleasePool() collection.ReleasePool {
	return inst.core.pool
}

func (inst *creationContextFacade) GetArguments() collection.Arguments {
	return inst.getParentRC().GetArguments()
}

func (inst *creationContextFacade) GetAttributes() collection.Attributes {
	return inst.getParentRC().GetAttributes()
}

func (inst *creationContextFacade) GetEnvironment() collection.Environment {
	return inst.getParentRC().GetEnvironment()
}

func (inst *creationContextFacade) GetProperties() collection.Properties {
	return inst.getParentRC().GetProperties()
}

func (inst *creationContextFacade) GetParameters() collection.Parameters {
	return inst.getParentRC().GetParameters()
}

func (inst *creationContextFacade) GetResources() collection.Resources {
	return inst.getParentRC().GetResources()
}

func (inst *creationContextFacade) OpenCreationContext(scope application.ComponentScope) application.CreationContext {
	return inst
}

func (inst *creationContextFacade) GetErrorHandler() lang.ErrorHandler {
	return inst.getParentRC().GetErrorHandler()
}

func (inst *creationContextFacade) SetErrorHandler(h lang.ErrorHandler) {
	inst.getParentRC().SetErrorHandler(h)
}

// as CreationContext

func (inst *creationContextFacade) GetScope() application.ComponentScope {
	return inst.core.scope
}

func (inst *creationContextFacade) GetContext() application.RuntimeContext {
	return inst.core.proxy
}

func (inst *creationContextFacade) Close() error {
	// todo:  close CC
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// impl creationComponentsFacade

func (inst *creationComponentsFacade) GetComponent(name string) (lang.Object, error) {
	holder, err := inst.core.finder.findHolderById(name)
	if err != nil {
		return nil, err
	}
	instance, err := inst.core.loader.loadComponent(holder)
	if err != nil {
		return nil, err
	}
	return instance.Get(), nil
}

func (inst *creationComponentsFacade) GetComponentByClass(classSelector string) (lang.Object, error) {
	holder, err := inst.core.finder.findHolderByTypeName(classSelector)
	if err != nil {
		return nil, err
	}
	instance, err := inst.core.loader.loadComponent(holder)
	if err != nil {
		return nil, err
	}
	return instance.Get(), nil
}

func (inst *creationComponentsFacade) GetComponentsByClass(classSelector string) []lang.Object {
	holders := inst.core.finder.findHoldersByTypeName(classSelector)
	instances, err := inst.core.loader.loadComponents(holders)
	if err != nil {
		inst.core.proxy.current.GetErrorHandler().OnError(err)
		return make([]lang.Object, 0)
	}
	results := inst.core.loader.toTargetArray(instances)
	return results
}

func (inst *creationComponentsFacade) Export(table map[string]application.ComponentHolder) map[string]application.ComponentHolder {
	// NOP
	if table == nil {
		table = make(map[string]application.ComponentHolder)
	}
	return table
}

func (inst *creationComponentsFacade) Import(map[string]application.ComponentHolder) {
	// NOP
}

////////////////////////////////////////////////////////////////////////////////
