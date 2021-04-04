package application

import (
	"github.com/bitwormhole/go-wormhole-core/collection"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// Context 表示一个通用的上下文对象
type Context interface {
	GetComponents() Components
	NewGetter(ec lang.ErrorCollector) ContextGetter

	GetReleasePool() collection.ReleasePool

	GetArguments() collection.Arguments
	GetAttributes() collection.Attributes
	GetEnvironment() collection.Environment
	GetProperties() collection.Properties
	GetParameters() collection.Parameters
	GetResources() collection.Resources
}

// RuntimeContext 是app的全局上下文
type RuntimeContext interface {
	Context

	GetApplicationName() string
	GetApplicationVersion() string
	GetStartupTimestamp() int64
	GetShutdownTimestamp() int64

	NewChild() RuntimeContext
}
