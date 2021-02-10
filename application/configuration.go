package application

// Configuration 表示应用程序配置
type Configuration interface {
	Components() []ComponentInfo
}
