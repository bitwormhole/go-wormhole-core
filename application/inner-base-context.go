package application

import "github.com/bitwormhole/go-wormhole-core/collection"

type baseContext struct {
	container *innerContextContainer
}

func (inst *baseContext) GetComponents() Components {
	return inst.container.components
}

func (inst *baseContext) NewGetter() ContextGetter {
	var ctx Context = inst
	return &innerContextGetter{
		context: ctx,
	}
}

func (inst *baseContext) GetReleasePool() collection.ReleasePool {
	return inst.container.releasePool
}

func (inst *baseContext) GetArguments() collection.Arguments {
	return inst.container.arguments
}

func (inst *baseContext) GetAttributes() collection.Attributes {
	return inst.container.attributes
}

func (inst *baseContext) GetEnvironment() collection.Environment {
	return inst.container.environment
}

func (inst *baseContext) GetProperties() collection.Properties {
	return inst.container.properties
}

func (inst *baseContext) GetParameters() collection.Parameters {
	return inst.container.parameters
}

func (inst *baseContext) GetResources() collection.Resources {
	return inst.container.resources
}
