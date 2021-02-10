package application

import "github.com/bitwormhole/go-wormhole-core/lang"

// ComponentScope 枚举表示组件的作用域
type ComponentScope uint32

const (
	// ScopeMin 是作用域的最小值
	ScopeMin ComponentScope = 0 // 最小

	// ScopeSingleton 表示单例模式
	ScopeSingleton ComponentScope = 1

	// ScopeContext 表示上下文模式
	ScopeContext ComponentScope = 2

	// ScopePrototype 表示原型模式
	ScopePrototype ComponentScope = 3

	// ScopeMax 是作用域的最大值
	ScopeMax ComponentScope = 4 // 最大
)

// ComponentInstanceRef 一个具体的组件的实例的引用
type ComponentInstanceRef interface {
	GetInstance() lang.Object
	Inject(context Context) error
	Init() error
	Destroy() error
}

// ComponentFactory 一个组件的工厂
type ComponentFactory interface {
	NewInstance() ComponentInstanceRef
}

// ComponentInfo 一个组件的配置
type ComponentInfo interface {
	GetID() string
	GetClass() string
	GetAliases() []string
	GetScope() ComponentScope
	GetFactory() ComponentFactory
}

// ComponentAgent 一个具体的组件的代理
type ComponentAgent interface {
	GetInstance() lang.Object
	GetInstanceRef() ComponentInstanceRef
	GetInfo() ComponentInfo
	GetContext() Context
	MakeChild(context Context) ComponentAgent
}

// Components 接口表示一个组件的集合
type Components interface {
	GetComponent(name string) (lang.Object, error)
	GetComponentByClass(classSelector string) (lang.Object, error)
	GetComponentsByClass(classSelector string) []lang.Object
	////
	GetAgent(name string) (ComponentAgent, error)
	SetAgent(name string, agent ComponentAgent)
	Clear()
	Export(map[string]ComponentAgent) map[string]ComponentAgent
	Import(map[string]ComponentAgent)
}
