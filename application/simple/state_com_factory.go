package simple

import (
	"errors"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

////////////////////////////////////////////////////////////////////////////////

type myStateComponentCloser struct {
	target application.ComponentInstanceRef
}

func (inst *myStateComponentCloser) Dispose() error {
	ptr := inst.target
	inst.target = nil
	if ptr == nil {
		return nil
	}
	return ptr.Destroy()
}

////////////////////////////////////////////////////////////////////////////////

// myStateComponentRef
type myStateComponentRef struct {
	context    application.Context
	target     application.ComponentInstanceRef
	hasInit    bool
	hasInject  bool
	hasDestroy bool
}

func (inst *myStateComponentRef) Inject(ctx application.Context) error {
	if inst.hasInject {
		return nil
	}
	inst.hasInject = true
	inst.context = ctx
	return inst.target.Inject(ctx)
}

func (inst *myStateComponentRef) Init() error {
	if inst.hasInit {
		return nil
	}
	if !inst.hasInject {
		return errors.New("init without inject")
	}
	inst.hasInit = true
	err := inst.target.Init()
	if err != nil {
		return err
	}
	closer := &myStateComponentCloser{target: inst.target}
	inst.context.GetReleasePool().Push(closer)
	return nil
}

func (inst *myStateComponentRef) Destroy() error {
	// NOP
	return nil
}

func (inst *myStateComponentRef) GetInstance() lang.Object {
	return inst.target.GetInstance()
}

////////////////////////////////////////////////////////////////////////////////

// myStateComponentFactory
type myStateComponentFactory struct {
	target application.ComponentFactory
}

func (inst *myStateComponentFactory) NewInstance() application.ComponentInstanceRef {
	ref := inst.target.NewInstance()
	return &myStateComponentRef{target: ref}
}

////////////////////////////////////////////////////////////////////////////////
// EOF
