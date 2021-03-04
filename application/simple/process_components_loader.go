package simple

import (
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/go-wormhole-core/application"
)

type myComponentLoader struct {
	building *myProcessContextBuilding
}

func (inst *myComponentLoader) load() error {

	ctx := &myInjectionContextFacade{}
	ctx.init(inst.building.context)

	err := inst.doComNew(ctx)
	if err != nil {
		return err
	}

	err = inst.doComInject(ctx)
	if err != nil {
		return err
	}

	err = inst.doComStart(ctx)
	if err != nil {
		return err
	}

	return ctx.start()
}

func (inst *myComponentLoader) addComponentAndAliases(id string, agent application.ComponentAgent) error {
	// add main agent
	err := inst.addComponentAgent(id, agent)
	if err != nil {
		return err
	}
	// add aliases agent
	aliases := agent.GetInfo().GetAliases()
	if aliases == nil {
		return nil
	}
	for index := range aliases {
		alias := aliases[index]
		err := inst.addComponentAlias(alias, id, agent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *myComponentLoader) addComponentAlias(aliasName string, targetName string, targetAgent application.ComponentAgent) error {
	ctx := targetAgent.GetContext()
	aliasAgent := &ComponentAliasAgent{
		aliasName:   aliasName,
		targetName:  targetName,
		targetAgent: targetAgent,
		context:     ctx,
	}
	return inst.addComponentAgent(aliasName, aliasAgent)
}

func (inst *myComponentLoader) addComponentAgent(name string, agent application.ComponentAgent) error {
	table := inst.building.comAgentTable
	old := table[name]
	if old != nil {
		return errors.New("the component id is conflict:" + name)
	}
	table[name] = agent
	return nil
}

func (inst *myComponentLoader) normalizeClassText(classText string) string {

	// example: "[.abc][.xyz]"

	old := " "
	new := ","
	classText = strings.ReplaceAll(classText, old, new)
	array := strings.Split(classText, new)
	builder := &strings.Builder{}

	sort.Strings(array)

	for index := range array {
		item := array[index]
		if item == "" {
			continue
		}
		builder.WriteString("[.")
		builder.WriteString(item)
		builder.WriteRune(']')
	}

	return builder.String()
}

func (inst *myComponentLoader) prepareComponent(info application.ComponentInfo) (application.ComponentAgent, error) {

	inst.building.comIndexGen++
	comIndex := (inst.building.comIndexGen)

	scope := info.GetScope()
	aliases := info.GetAliases()
	id := info.GetID()
	clazz := info.GetClass()
	factory := info.GetFactory()

	//scope
	if scope <= application.ScopeMin || scope >= application.ScopeMax {
		scope = application.ScopeSingleton
	}

	// aliases
	if aliases == nil {
		aliases = []string{}
	}

	// id + class
	clazz = inst.normalizeClassText(clazz)
	if id == "" {
		id = fmt.Sprint("com_", comIndex, clazz)
	}

	// factory
	factory = &myStateComponentFactory{target: factory}

	// wrap info
	info = &RuntimeComInfo{
		id:      id,
		class:   clazz,
		scope:   scope,
		factory: factory,
		aliases: aliases,
	}

	// agent
	context := inst.building.context
	var agent application.ComponentAgent
	if scope == application.ScopePrototype {
		agent = &myPrototypeComAgent{
			info:    info,
			context: context,
		}
	} else {
		agent = &mySingletonComAgent{
			info:    info,
			context: context,
		}
	}
	return agent, nil
}

func (inst *myComponentLoader) doComNew(context application.Context) error {
	list := inst.building.comInfoList
	if list == nil {
		return nil
	}
	for index := range list {
		info := list[index]
		agent, err := inst.prepareComponent(info)
		if err != nil {
			return err
		}
		info = agent.GetInfo()
		id := info.GetID()
		err = inst.addComponentAndAliases(id, agent)
		if err != nil {
			return err
		}
	}
	return nil
}

func (inst *myComponentLoader) doComInject(context application.Context) error {

	coms := context.GetComponents()
	all := coms.Export(nil)

	for key := range all {
		agent := all[key]
		scope := agent.GetInfo().GetScope()
		if scope == application.ScopeSingleton {
			ref := agent.GetInstanceRef()
			err := ref.Inject(context)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func (inst *myComponentLoader) doComStart(context application.Context) error {
	return nil
}
