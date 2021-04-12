package runtime

import (
	"errors"

	"github.com/bitwormhole/go-wormhole-core/application"
)

type componentFinder struct {
	table map[string]application.ComponentHolder
}

func (inst *componentFinder) findHolderById(id string) (application.ComponentHolder, error) {
	holder := inst.table[id]
	if holder == nil {
		return nil, errors.New("no component with id:" + id)
	}
	return holder, nil
}

func (inst *componentFinder) findHolderByTypeName(name string) (application.ComponentHolder, error) {
	list := inst.findHoldersByTypeName(name)
	size := len(list)
	if size == 1 {
		return list[0], nil
	} else if size > 1 {
		return nil, errors.New("there are several components, more then one, with type of " + name)
	}
	return nil, errors.New("there is no component, with type of " + name)
}

func (inst *componentFinder) findHoldersByTypeName(name string) []application.ComponentHolder {
	list := make([]application.ComponentHolder, 0)
	table := inst.table
	for key := range table {
		holder := table[key]
		info := holder.GetInfo()
		if key != info.GetID() {
			continue
		}
		if info.IsTypeOf(name) {
			list = append(list, holder)
		}
	}
	return list
}
