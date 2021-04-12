package application

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration) (RuntimeContext, error) {
	return config.GetLoader().Load(config)
}

// Exit 函数用于退出应用
func Exit(context RuntimeContext) int32 {
	context.GetReleasePool().Release()
	return 0
}
