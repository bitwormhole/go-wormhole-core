package runtime

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

type componentLoading struct {
	holder   application.ComponentHolder
	instance application.ComponentInstance
}

type componentLoader struct {
	core *creationContextCore
}

////////////////////////////////////////////////////////////////////////////////
// impl componentLoader

func (inst *componentLoader) loadComponent(holder application.ComponentHolder) (application.ComponentInstance, error) {
	return nil, nil
}

func (inst *componentLoader) loadComponents(holders []application.ComponentHolder) ([]application.ComponentInstance, error) {
	return nil, nil
}

func (inst *componentLoader) toTargetArray(src []application.ComponentInstance) []lang.Object {
	if src == nil {
		return make([]lang.Object, 0)
	}
	size := len(src)
	dst := make([]lang.Object, size)
	for i := 0; i < size; i++ {
		dst[i] = src[i].Get()
	}
	return dst
}
