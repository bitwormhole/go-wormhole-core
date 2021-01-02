package application

// innerFragmentContext
type innerFragmentContext struct {
	baseContext
}

func (inst *innerFragmentContext) GetParent() NodeContext {
	return nil
}

func (inst *innerFragmentContext) GetRoot() ProcessContext {
	return nil
}

func (inst *innerFragmentContext) NewChild() FragmentContext {
	return nil
}

//  FragmentContextBuilder
type FragmentContextBuilder struct{}

func (inst *FragmentContextBuilder) Create() FragmentContext {
	return &innerFragmentContext{}
}

func (inst *FragmentContextBuilder) Init(parent NodeContext) {

}
