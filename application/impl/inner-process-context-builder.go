package impl

import "github.com/bitwormhole/go-wormhole-core/application"

// innerProcessContext
type innerProcessContext struct {
	baseContext

	appName    string
	appVersion string
}

func (inst *innerProcessContext) GetParent() application.NodeContext {
	return nil
}

func (inst *innerProcessContext) GetRoot() application.ProcessContext {
	return inst
}

func (inst *innerProcessContext) NewChild() application.FragmentContext {
	return nil
}

func (inst *innerProcessContext) GetApplicationName() string {
	return inst.appName
}

func (inst *innerProcessContext) GetApplicationVersion() string {
	return inst.appVersion
}

func (inst *innerProcessContext) GetStartupTimestamp() int64 {
	return 0
}

func (inst *innerProcessContext) GetShutdownTimestamp() int64 {
	return 0
}

// ProcessContextBuilder 是 ProcessContext 实例的创建者
type ProcessContextBuilder struct {
	config *application.Configuration
}

// Create 方法用于创建进程上下文
func (inst *ProcessContextBuilder) Create() (application.ProcessContext, error) {

	// 构造容器中的各个集合

	container := &innerContextContainer{}
	pc := &innerProcessContext{}

	container.componentTable = make(map[string]ComponentHolder)
	container.contextFacade = pc
	container.arguments = nil
	container.attributes = nil
	container.components = &innerComponents{container: container}
	container.environment = nil
	container.parameters = nil
	container.properties = nil
	container.releasePool = nil
	container.resources = nil

	pc.baseContext.container = container

	// 从configuration中加载各个组件
	inst.loadComponents(container)

	return pc, nil
}

func (inst *ProcessContextBuilder) loadComponent(reg application.ComponentRegistration) ComponentHolder {
	info := reg.GetInfo()
	switch info.Scope {
	case ScopeContext:
		return createContextScopeComponentHolder(reg)
	case ScopePrototype:
		return createPrototypeScopeComponentHolder(reg)
	case ScopeSingleton:
	default:
		break
	}
	return createSingletonScopeComponentHolder(reg)
}

func (inst *ProcessContextBuilder) prepareComponentRegistration(cr application.ComponentRegistration, index int) application.ComponentRegistration {
	cr2 := &innerComponentRegistration{}
	cr2.init(cr, index)
	return cr2
}

func (inst *ProcessContextBuilder) loadComponents(container *innerContextContainer) {

	src := inst.config.registrations
	dst := container.componentTable

	for index := range src {
		item := src[index]
		item = inst.prepareComponentRegistration(item, index)
		holder := inst.loadComponent(item)
		info := holder.GetInfo()
		dst[info.Name] = holder
	}
}

// Init 方法用于初始化Builder对象
func (inst *ProcessContextBuilder) Init(cfg *Configuration) {
	inst.config = cfg
}
