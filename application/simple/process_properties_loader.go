package simple

import (
	"fmt"

	"github.com/bitwormhole/go-wormhole-core/collection"
)

type processPropertiesLoader struct {
	building *myProcessContextBuilding
}

func (inst *processPropertiesLoader) load() error {

	key := "application.profiles.active"

	inst.building.propertiesFromRes1 = inst.loadFromRes("application.properties")
	inst.remix()

	master := inst.building.propertiesMaster
	profile := master[key]

	if profile != "" {
		name := "application-" + profile + ".properties"
		inst.building.propertiesFromRes2 = inst.loadFromRes(name)
		inst.remix()
	}

	master = inst.building.propertiesMaster
	profile = master[key]

	fmt.Println(key, "=", profile)

	return nil
}

func (inst *processPropertiesLoader) copyProps(src map[string]string, dst map[string]string) {

	if src == nil {
		return
	}

	for key := range src {
		val := src[key]
		dst[key] = val
	}
}

func (inst *processPropertiesLoader) remix() error {

	table := make(map[string]string)
	src := inst.building

	inst.copyProps(src.propertiesFromRes1, table)
	inst.copyProps(src.propertiesFromRes2, table)
	inst.copyProps(src.propertiesFromArgs, table)

	inst.building.propertiesMaster = table

	return nil
}

func (inst *processPropertiesLoader) loadFromRes(path string) map[string]string {

	table := make(map[string]string)
	res := inst.building.resources
	text, err := res.GetText(path)
	if err != nil {
		return table
	}

	props := &collection.SimpleProperties{}
	props2, err := collection.ParseProperties(text, props)
	if err != nil {
		return table
	}

	table = props2.Export(table)
	return table
}
