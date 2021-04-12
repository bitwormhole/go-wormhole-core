package application

import "github.com/bitwormhole/go-wormhole-core/collection"

// Configuration 表示应用程序配置
type Configuration interface {
	GetBuilder() ConfigBuilder
	GetLoader() ContextLoader
	GetComponents() []ComponentInfo
	GetResources() collection.Resources
}

//  ContextLoader 用于加载进程上下文
type ContextLoader interface {
	Load(config Configuration) (RuntimeContext, error)
}

// ConfigBuilder 表示应用程序配置
type ConfigBuilder interface {
	AddComponent(info ComponentInfo)
}
