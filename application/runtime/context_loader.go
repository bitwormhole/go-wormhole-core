package runtime

import (
	"errors"

	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/lang"
)

// RuntimeContextLoader 运行时上下文加载器
type RuntimeContextLoader struct {
	comInfoList []application.ComponentInfo
	comTable    map[string]application.ComponentHolder
	context     application.RuntimeContext
	config      application.Configuration
	todoInit    application.CreationContext
}

// Load 方法根据传入的配置加载运行时上下文
func (inst *RuntimeContextLoader) Load(config application.Configuration) (application.RuntimeContext, error) {

	inst.config = config
	inst.comTable = make(map[string]application.ComponentHolder)
	inst.comInfoList = nil
	inst.context = nil

	tc := &lang.TryChain{}

	tc.Try(func() error {
		return nil
	}).Try(func() error {
		return inst.createRuntimeContext()
	}).Try(func() error {
		return inst.prepareComInfoList()
	}).Try(func() error {
		return inst.doCreateComponents()
	}).Try(func() error {
		return inst.doInjectComponents()
	}).Try(func() error {
		return inst.doInitComponents()
	}).Try(func() error {
		return nil
	})

	err := tc.Result()
	ctx := inst.context
	if err != nil {
		ctx = nil
	}
	return ctx, err
}

func (inst *RuntimeContextLoader) createRuntimeContext() error {

	core := createRuntimeContextCore()

	core.appName = ""
	core.appVersion = ""
	core.time1 = 0
	core.time2 = 0
	core.uri = ""
	core.resources = inst.config.GetResources()

	inst.context = core.context
	return nil
}

func (inst *RuntimeContextLoader) prepareComInfoList() error {
	src := inst.config.GetComponents()
	dst := make([]application.ComponentInfo, 0)
	preprocessor := &componentInfoPreprocessor{}
	for index := range src {
		info := src[index]
		info, err := preprocessor.prepare(info, index)
		if err != nil {
			return err
		}
		dst = append(dst, info)
	}
	inst.comInfoList = dst
	return nil
}

func (inst *RuntimeContextLoader) doCreateComponents() error {

	// 根据 info 创建 对应的 holder

	ctx := inst.context
	src := inst.comInfoList
	dst := make(map[string]application.ComponentHolder)

	for index := range src {
		info := src[index]
		scope := info.GetScope()
		var holder application.ComponentHolder
		if scope == application.ScopeSingleton {
			holder = &SingletonComponentHolder{context: ctx, info: info}
		} else if scope == application.ScopePrototype {
			holder = &PrototypeComponentHolder{context: ctx, info: info}
		} else if scope == application.ScopeContext {
			continue
		} else {
			continue
		}
		err := inst.putComHolderToTable(dst, holder)
		if err != nil {
			return err
		}
	}

	// 导入到 context 里
	com_set := ctx.GetComponents()
	com_set.Import(dst)
	inst.comTable = com_set.Export(nil)

	return nil
}

func (inst *RuntimeContextLoader) putComHolderToTable(table map[string]application.ComponentHolder, holder application.ComponentHolder) error {

	info := holder.GetInfo()
	id := info.GetID()
	aliases := info.GetAliases()

	id_in_list := false
	for index := range aliases {
		name := aliases[index]
		if name == id {
			id_in_list = true
			break
		}
	}
	if !id_in_list {
		aliases = append(aliases, id)
	}

	for index := range aliases {
		name := aliases[index]
		older := table[name]
		if older != nil {
			return errors.New("the ID (alias) of component is duplicate:" + name)
		}
		table[name] = holder
	}

	return nil
}

func (inst *RuntimeContextLoader) doInjectComponents() error {

	scopeWant := application.ScopeSingleton

	cc := inst.context.OpenCreationContext(scopeWant)
	context := cc.GetContext()
	components := context.GetComponents()
	table := components.Export(nil)

	for name := range table {
		holder := table[name]
		info := holder.GetInfo()
		id := info.GetID()
		scope := info.GetScope()
		if (id == name) && (scope == scopeWant) {
			_, err := components.GetComponent(name)
			if err != nil {
				return err
			}
		}
	}

	inst.todoInit = cc
	return nil
}

func (inst *RuntimeContextLoader) doInitComponents() error {
	cc := inst.todoInit
	if cc == nil {
		return errors.New("no CreationContext")
	}
	return cc.Close()
}
