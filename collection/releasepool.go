package collection

import (
	"github.com/bitwormhole/gss/lang"
)

type ReleasePool interface {
	Release()
	Push(target lang.Disposable)
}
