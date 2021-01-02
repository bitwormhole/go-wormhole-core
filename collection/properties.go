package collection

// Properties 接口表示对属性列表的引用。
type Properties interface {
	GetProperty(name string) (string, error)
}
