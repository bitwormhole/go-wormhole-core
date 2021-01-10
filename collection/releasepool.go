package collection

import (
	"github.com/bitwormhole/go-wormhole-core/lang"
)

type ReleasePool interface {
	Release()
	Push(target lang.Disposable)
}
