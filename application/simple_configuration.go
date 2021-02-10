package application

import "github.com/bitwormhole/go-wormhole-core/lang"

////////////////////////////////////////////////////////////////////////////////

// SimpleConfiguration 提供一个简易的 Configuration 实现
type SimpleConfiguration struct {
	// implements Configuration
	components []ComponentInfo
}

func (inst *SimpleConfiguration) getComList(create bool) []ComponentInfo {
	list := inst.components
	if (create) && (list == nil) {
		list = make([]ComponentInfo, 0)
		inst.components = list
	}
	return list
}

// Components 返回组件的注册信息
func (inst *SimpleConfiguration) Components() []ComponentInfo {
	return inst.getComList(true)
}

// AddComponent 注册一个组件
func (inst *SimpleConfiguration) AddComponent(info ComponentInfo) {
	list := inst.getComList(true)
	inst.components = append(list, info)
}

////////////////////////////////////////////////////////////////////////////////

// SimpleCom 提供一个简易的 ComponentInfo 实现
type SimpleCom struct {
	// implements ComponentInfo
	ID      string
	Class   string
	Scope   ComponentScope
	Aliases []string

	OnNew     func() lang.Object
	OnInit    func(obj lang.Object) error
	OnDestroy func(obj lang.Object) error
	OnInject  func(obj lang.Object, context Context) error
}

// GetAliases 获取组件的别名
func (inst *SimpleCom) GetAliases() []string {
	return inst.Aliases
}

// GetID 获取组件的ID
func (inst *SimpleCom) GetID() string {
	return inst.ID
}

// GetClass 获取组件的类
func (inst *SimpleCom) GetClass() string {
	return inst.Class
}

// GetScope 获取组件的作用域
func (inst *SimpleCom) GetScope() ComponentScope {
	return 0
}

// GetFactory 获取组件的工厂
func (inst *SimpleCom) GetFactory() ComponentFactory {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
