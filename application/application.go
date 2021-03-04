package application

// ProcessContextLoader 用于加载进程上下文
type ProcessContextLoader interface {
	Load(config Configuration) (ProcessContext, error)
}

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration) (ProcessContext, error) {
	return config.GetLoader().Load(config)
}

// Exit 函数用于退出应用
func Exit(context ProcessContext) int32 {
	context.GetReleasePool().Release()
	return 0
}
