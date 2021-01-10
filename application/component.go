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

type ComponentInfo struct {
	Name  string
	Class string
	Scope ComponentScope
}

// Component 是一个标记对象, 以该结构标记的struct是一个
type Component struct {
}

/////////////////////////////////

type ComponentInstance interface {
	Inject(context Context) error
	Init() error
	Destroy() error
	GetComponent() lang.Object
}

type ComponentFactory interface {
	NewInstance() ComponentInstance
}

type ComponentRegistration interface {
	ComponentFactory
	ComponentInstance
	GetInfo() ComponentInfo
	GetFactory() ComponentFactory
}

/////////////////////////////////

type ComponentHolder interface {
	GetInstance() ComponentInstance
	GetFactory() ComponentFactory
	GetInfo() ComponentInfo
	GetScope() ComponentScope
	MakeChild() ComponentHolder
}

// Components 接口表示一个组件的集合
type Components interface {
	GetComponent(name string) (lang.Object, error)
	GetComponentByClass(classSelector string) (lang.Object, error)
	GetComponentsByClass(classSelector string) []lang.Object
}
