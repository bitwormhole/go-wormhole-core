package application

import "github.com/bitwormhole/gss/collection"

// Context 表示一个通用的上下文对象
type Context interface {
	GetComponents() Components
	NewGetter() ContextGetter

	GetReleasePool() collection.ReleasePool

	GetArguments() collection.Arguments
	GetAttributes() collection.Attributes
	GetEnvironment() collection.Environment
	GetProperties() collection.Properties
	GetParameters() collection.Parameters
	GetResources() collection.Resources
}

// NodeContext 是一个抽象的上下文节点
type NodeContext interface {
	Context
	GetRoot() ProcessContext
	GetParent() NodeContext
	NewChild() FragmentContext
}

// ProcessContext 是app的全局上下文
type ProcessContext interface {
	NodeContext
	GetApplicationName() string
	GetApplicationVersion() string
	GetStartupTimestamp() int64
	GetShutdownTimestamp() int64
}

// FragmentContext 是一个局部的上下文
type FragmentContext interface {
	NodeContext
}
