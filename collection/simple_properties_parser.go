package collection

type simplePropertiesParser struct{}

// ParseProperties 函数把参数 text 解析为属性表，存入dest中。
func ParseProperties(text string, dest Properties) error {

	parser := &simplePropertiesParser{}
	parser.todo()
	return nil
}

func (inst *simplePropertiesParser) todo() {}
