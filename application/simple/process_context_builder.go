package simple

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
)

////////////////////////////////////////////////////////////////////////////////

// myProcessContext 是默认的进程上下文
type myProcessContext struct {
	BaseContext
}

func (inst *myProcessContext) GetApplicationVersion() string {
	return ""
}

func (inst *myProcessContext) GetApplicationName() string {
	return ""
}

func (inst *myProcessContext) GetStartupTimestamp() int64 {
	return 0
}

func (inst *myProcessContext) GetShutdownTimestamp() int64 {
	return 0
}

func (inst *myProcessContext) GetRoot() application.ProcessContext {
	return inst
}

func (inst *myProcessContext) GetParent() application.NodeContext {
	return nil
}

func (inst *myProcessContext) NewChild() application.FragmentContext {
	return nil
}

////////////////////////////////////////////////////////////////////////////////

type myProcessContextBuilding struct {
	comInfoList        []application.ComponentInfo
	comAgentTable      map[string]application.ComponentAgent
	context            *myProcessContext
	comIndexGen        int32
	arguments          []string
	resources          collection.Resources
	propertiesFromArgs map[string]string
	propertiesFromRes1 map[string]string
	propertiesFromRes2 map[string]string
	propertiesMaster   map[string]string
}

////////////////////////////////////////////////////////////////////////////////

// ProcessContextBuilder 用于创建 ProcessContext
type myProcessContextBuilder struct {
	building *myProcessContextBuilding
}

// Init 用配置初始化这个   ProcessContextBuilder
func (inst *myProcessContextBuilder) Init(config application.Configuration) {

	pc := &myProcessContext{}
	building := &myProcessContextBuilding{}
	comSet := &BaseComponentSet{}
	comTable := make(map[string]application.ComponentAgent)
	res := config.GetResources()

	comSet.context = pc
	comSet.table = comTable

	pc.components = comSet
	pc.attributes = nil
	pc.arguments = nil
	pc.resources = res

	building.context = pc
	building.comInfoList = nil
	building.comAgentTable = comTable
	building.resources = res

	inst.building = building
}

// Create 创建一个新的进程上下文
func (inst *myProcessContextBuilder) Create() (application.ProcessContext, error) {

	comLoader := &myComponentLoader{building: inst.building}

	err := comLoader.load()
	if err != nil {
		return nil, err
	}
	return inst.building.context, nil
}

func (inst *myProcessContextBuilder) loadProperties() {

	loader := &processPropertiesLoader{building: inst.building}
	err := loader.load()
	if err != nil {
		panic(err.Error())
	}
}

func (inst *myProcessContextBuilder) loadArguments() {
	loader := &processArgumentsLoader{building: inst.building}
	err := loader.load()
	if err != nil {
		panic(err.Error())
	}
}

func (inst *myProcessContextBuilder) loadResources() {
	// NOP
}

func (inst *myProcessContextBuilder) loadComponents() {
	loader := &myComponentLoader{building: inst.building}
	err := loader.load()
	if err != nil {
		panic(err.Error())
	}
}

// Load
func (inst *myProcessContextBuilder) Load(config application.Configuration) (application.ProcessContext, error) {

	inst.Init(config)

	inst.loadResources()
	inst.loadArguments()
	inst.loadProperties()
	inst.loadComponents()

	return inst.Create()
}

////////////////////////////////////////////////////////////////////////////////
// EOF
