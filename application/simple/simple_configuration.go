package simple

import (
	"embed"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

////////////////////////////////////////////////////////////////////////////////

// AppConfig 提供一个简易的 Configuration 实现
type AppConfig struct {
	// implements Configuration
	components []application.ComponentInfo
	resources  collection.Resources
}

func (inst *AppConfig) getComList(create bool) []application.ComponentInfo {
	list := inst.components
	if (create) && (list == nil) {
		list = make([]application.ComponentInfo, 0)
		inst.components = list
	}
	return list
}

// GetComponents 返回组件的注册信息
func (inst *AppConfig) GetComponents() []application.ComponentInfo {
	return inst.getComList(true)
}

// AddComponent 注册一个组件
func (inst *AppConfig) AddComponent(info application.ComponentInfo) {
	list := inst.getComList(true)
	inst.components = append(list, info)
}

// GetLoader 返回加载器
func (inst *AppConfig) GetLoader() application.ProcessContextLoader {
	return &myProcessContextBuilder{}
}

// GetBuilder 返回构建器
func (inst *AppConfig) GetBuilder() application.ConfigBuilder {
	return inst
}

// SetResources 用于配置上下文的资源文件夹
func (inst *AppConfig) SetResources(fs *embed.FS, prefix string) {
	inst.resources = &simpleEmbedResFS{
		fs:     fs,
		prefix: prefix,
	}
}

// GetResources 用于获取上下文的资源文件夹
func (inst *AppConfig) GetResources() collection.Resources {
	return inst.resources
}

////////////////////////////////////////////////////////////////////////////////

// ComInfo 提供一个简易的 ComponentInfo 实现
type ComInfo struct {
	// implements ComponentInfo
	ID      string
	Class   string
	Scope   application.ComponentScope
	Aliases []string

	OnNew     func() lang.Object
	OnInit    func(obj lang.Object) error
	OnDestroy func(obj lang.Object) error
	OnInject  func(obj lang.Object, context application.Context) error
}

// GetAliases 获取组件的别名
func (inst *ComInfo) GetAliases() []string {
	return inst.Aliases
}

// GetID 获取组件的ID
func (inst *ComInfo) GetID() string {
	return inst.ID
}

// GetClass 获取组件的类
func (inst *ComInfo) GetClass() string {
	return inst.Class
}

// GetScope 获取组件的作用域
func (inst *ComInfo) GetScope() application.ComponentScope {
	return inst.Scope
}

// GetFactory 获取组件的工厂
func (inst *ComInfo) GetFactory() application.ComponentFactory {
	return nil
}

////////////////////////////////////////////////////////////////////////////////
// EOF
