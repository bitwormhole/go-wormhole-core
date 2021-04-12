package collection

import "github.com/bitwormhole/go-wormhole-core/lang"

type Attributes interface {
	GetAttribute(name string) (lang.Object, error)

	Import(map[string]lang.Object)
	Export(map[string]lang.Object) map[string]lang.Object
}
