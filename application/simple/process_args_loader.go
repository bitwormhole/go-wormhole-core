package simple

import (
	"os"
	"strings"
)

type processArgumentsLoader struct {
	building *myProcessContextBuilding
}

func (inst *processArgumentsLoader) load() error {

	props := make(map[string]string)
	args := os.Args
	args1 := args[1:]

	for index := range args1 {
		arg := args1[index]
		inst.parseArgument(arg, props)
	}

	inst.building.arguments = args1
	inst.building.propertiesFromArgs = props

	return nil
}

func (inst *processArgumentsLoader) parseArgument(text string, table map[string]string) {

	prefix := "--"
	if strings.HasPrefix(text, prefix) {
		text = strings.TrimPrefix(text, prefix)
	} else {
		return
	}

	index := strings.IndexRune(text, '=')
	if index < 1 {
		return
	}

	array := strings.SplitN(text, "=", 2)
	key := array[0]
	val := array[1]

	table[key] = val
}
