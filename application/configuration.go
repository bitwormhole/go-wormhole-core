package application

import "github.com/bitwormhole/go-wormhole-core/collection"

// Configuration 表示应用程序配置
type Configuration interface {
	GetBuilder() ConfigBuilder
	GetLoader() ProcessContextLoader
	GetComponents() []ComponentInfo
	GetResources() collection.Resources
}

// ConfigBuilder 表示应用程序配置
type ConfigBuilder interface {
	AddComponent(info ComponentInfo)
}
