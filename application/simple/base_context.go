package simple

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// BaseContext 提供一个基本的上下文实现
type BaseContext struct {
	components application.Components

	pool        collection.ReleasePool
	arguments   collection.Arguments
	attributes  collection.Attributes
	environment collection.Environment
	properties  collection.Properties
	parameters  collection.Parameters
	resources   collection.Resources
}

// GetComponents 获取这个上下文的组件列表
func (inst *BaseContext) GetComponents() application.Components {
	return inst.components
}

// GetEnvironment 获取这个上下文的环境变量
func (inst *BaseContext) GetEnvironment() collection.Environment {
	return inst.environment
}

// NewGetter 取...集合
func (inst *BaseContext) NewGetter(ec lang.ErrorCollector) application.ContextGetter {
	return nil
}

// GetReleasePool 取...集合
func (inst *BaseContext) GetReleasePool() collection.ReleasePool {
	return inst.pool
}

// GetArguments 取...集合
func (inst *BaseContext) GetArguments() collection.Arguments {
	return inst.arguments
}

// GetAttributes 取...集合
func (inst *BaseContext) GetAttributes() collection.Attributes {
	return inst.attributes
}

// GetProperties 取...集合
func (inst *BaseContext) GetProperties() collection.Properties {
	return inst.properties
}

// GetParameters 取...集合
func (inst *BaseContext) GetParameters() collection.Parameters {
	return inst.parameters
}

// GetResources 取...集合
func (inst *BaseContext) GetResources() collection.Resources {
	return inst.resources
}
