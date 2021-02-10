package application

// Run 函数启动一个应用实例，返回应用上下文
func Run(config Configuration) (ProcessContext, error) {
	// TODO
	//	builder := &ProcessContextBuilder{}
	//	builder.Init(config)
	//	return builder.Create()
	return nil, nil
}

// Exit 函数用于退出应用
func Exit(context ProcessContext) int32 {
	context.GetReleasePool().Release()
	return 0
}
