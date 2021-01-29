package application

import "github.com/bitwormhole/go-wormhole-core/lang"

// ComponentScope 枚举表示组件的作用域
type ComponentScope uint32

const (
	ScopeMin       ComponentScope = 0 // 最小
	ScopeSingleton ComponentScope = 1
	ScopeContext   ComponentScope = 2
	ScopePrototype ComponentScope = 3
	ScopeMax       ComponentScope = 4 // 最大
)

// ComponentAgent 一个具体的组件的代理
type ComponentAgent interface {
	IsAlias() bool
	GetInstance() lang.Object
	GetInstanceHolder() ComponentInstanceHolder
	GetFactory() ComponentFactory
	GetInfo() ComponentInfo
	GetContext() Context
	MakeChild(context Context) ComponentAgent
}

// ComponentInstanceHolder 一个具体的组件的实例托管者
type ComponentInstanceHolder interface {
	GetInstance() lang.Object

	GetFactory() ComponentFactory
	GetContext() Context

	Inject() error
	Init() error
	Destroy() error
}

// ComponentInstanceRef 一个具体的组件的实例的引用
type ComponentInstanceRef interface {
	GetInstance() lang.Object
	Inject(context Context) error
	Init() error
	Destroy() error
}

// ComponentProvider 一个组件的实例创建者
type ComponentProvider interface {
	NewInstance() ComponentInstanceRef
}

// ComponentFactory 一个组件的工厂
type ComponentFactory interface {
	GetInfo() ComponentInfo
	NewAgent(context Context) ComponentAgent
	NewInstance(context Context) ComponentInstanceHolder
}

// ComponentInfo 一个组件的配置
type ComponentInfo interface {
	GetID() string
	GetClass() string
	GetScope() ComponentScope
	GetProvider() ComponentProvider
}

// Components 接口表示一个组件的集合
type Components interface {
	GetComponent(name string) (lang.Object, error)
	GetComponentByClass(classSelector string) (lang.Object, error)
	GetComponentsByClass(classSelector string) []lang.Object
	////
	GetAgent(name string) (ComponentAgent, error)
	SetAgent(name string, value ComponentAgent)
	Clear()
	Export(map[string]ComponentAgent) map[string]ComponentAgent
	Import(map[string]ComponentAgent)
}
