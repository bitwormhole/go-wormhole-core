package application

// ComponentRegistration 表示组件的注册信息
type ComponentRegistration struct {
	ID     string
	Class  string
	Scope  ComponentScope
	Source func() ComponentInstanceRef
}

// Configuration 表示应用程序配置
type Configuration struct {
	registrations []*ComponentRegistration
}

func (inst *Configuration) Component(item *ComponentRegistration) {

	list := inst.registrations

	if list == nil {
		list = make([]*ComponentRegistration, 0)
	}

	list = append(list, item)
	inst.registrations = list
}
