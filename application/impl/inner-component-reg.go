package impl

import (
	"fmt"
	"sort"
	"strings"

	"github.com/bitwormhole/go-wormhole-core/application"
)

type innerComponentRegistration struct {
	application.ComponentRegistration

	id      string
	class   string
	scope   application.ComponentScope
	factory application.ComponentFactory
}

func (inst *innerComponentRegistration) normalizeClassString(class string) string {

	var b strings.Builder
	class = strings.ReplaceAll(class, " ", ",")
	array := strings.Split(class, ",")
	sort.Strings(array)

	for index := range array {
		item := array[index]
		item = strings.TrimSpace(item)
		if len(item) == 0 {
			continue
		}
		b.WriteString(" .")
		b.WriteString(item)
		b.WriteString(" ")
	}

	return b.String()
}

func (inst *innerComponentRegistration) init(src application.ComponentRegistration, index int) {

	info := src.GetInfo()
	id := info.Name
	class := info.Class
	scope := info.Scope

	// class
	class = inst.normalizeClassString(class)

	// id (aka.name)
	id = strings.TrimSpace(id)
	if len(id) == 0 {
		id = fmt.Sprint("#.wh-component-", index, " ", class)
	} else if id[0] == '#' {
		// NOP
	} else {
		id = "#" + id
	}

	// scope
	if (scope <= application.ScopeMin) || (application.ScopeMax <= scope) {
		scope = application.ScopeSingleton
	}

	// fill inst
	inst.factory = src.GetFactory()
	inst.id = id
	inst.class = class
	inst.scope = scope
}
