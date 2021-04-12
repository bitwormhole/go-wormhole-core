package runtime

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

////////////////////////////////////////////////////////////////////////////////
// struct

type runtimeContextFacade struct {
	core *runtimeContextCore
}

type runtimeComponentsFacade struct {
	core *runtimeContextCore
}

////////////////////////////////////////////////////////////////////////////////
// impl runtimeContextFacade

func (inst *runtimeContextFacade) GetComponents() application.Components {
	return inst.core.components
}

func (inst *runtimeContextFacade) GetReleasePool() collection.ReleasePool {
	return inst.core.releasePool
}

func (inst *runtimeContextFacade) GetArguments() collection.Arguments {
	return inst.core.arguments
}

func (inst *runtimeContextFacade) GetAttributes() collection.Attributes {
	return inst.core.attributes
}

func (inst *runtimeContextFacade) GetEnvironment() collection.Environment {
	return inst.core.environment
}

func (inst *runtimeContextFacade) GetProperties() collection.Properties {
	return inst.core.properties
}

func (inst *runtimeContextFacade) GetParameters() collection.Parameters {
	return inst.core.parameters
}

func (inst *runtimeContextFacade) GetResources() collection.Resources {
	return inst.core.resources
}

func (inst *runtimeContextFacade) GetApplicationName() string {
	return inst.core.appName
}

func (inst *runtimeContextFacade) GetApplicationVersion() string {
	return inst.core.appVersion
}

func (inst *runtimeContextFacade) GetStartupTimestamp() int64 {
	return inst.core.time1
}

func (inst *runtimeContextFacade) GetShutdownTimestamp() int64 {
	return inst.core.time2
}

func (inst *runtimeContextFacade) GetURI() string {
	return inst.core.uri
}

func (inst *runtimeContextFacade) GetErrorHandler() lang.ErrorHandler {
	h := inst.core.errorHandler
	if h == nil {
		h = lang.DefaultErrorHandler()
	}
	return h
}

func (inst *runtimeContextFacade) SetErrorHandler(h lang.ErrorHandler) {
	inst.core.errorHandler = h
}

func (inst *runtimeContextFacade) OpenCreationContext(scope application.ComponentScope) application.CreationContext {
	ccc := createCreationContextCore(inst.core)
	ccc.scope = scope
	if scope == application.ScopeSingleton {
		ccc.pool = inst.core.releasePool
	} else {
		ccc.pool = collection.CreateReleasePool()
	}
	return ccc.facade
}

func (inst *runtimeContextFacade) NewChild() application.RuntimeContext {
	ctx, _ := CreateRuntimeContext(inst)
	return ctx
}

func (inst *runtimeContextFacade) NewGetter(ec lang.ErrorCollector) application.ContextGetter {
	getter := &innerContextGetter{}
	getter.init(inst, ec)
	return getter
}

////////////////////////////////////////////////////////////////////////////////
// impl runtimeComponentsFacade

func (inst *runtimeComponentsFacade) Import(map[string]application.ComponentHolder) {

}

func (inst *runtimeComponentsFacade) Export(table map[string]application.ComponentHolder) map[string]application.ComponentHolder {
	return table
}

func (inst *runtimeComponentsFacade) GetComponent(name string) (lang.Object, error) {
	return nil, nil
}

func (inst *runtimeComponentsFacade) GetComponentByClass(classSelector string) (lang.Object, error) {
	return nil, nil
}

func (inst *runtimeComponentsFacade) GetComponentsByClass(classSelector string) []lang.Object {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
