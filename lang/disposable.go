package lang

// Disposable 接口用于释放对象持有的资源
type Disposable interface {
	Dispose() error
}
