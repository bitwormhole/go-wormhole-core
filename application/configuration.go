package application

// Configuration 表示应用程序配置
type Configuration struct {
	registrations []ComponentRegistration
}

func (inst *Configuration) tryInit() {
	if inst.registrations == nil {
		inst.registrations = make([]ComponentRegistration, 0)
	}
}

// Component 方法用于注册组件
func (inst *Configuration) Component(source ComponentRegistration) {
	inst.tryInit()
	inst.registrations = append(inst.registrations, source)
}
