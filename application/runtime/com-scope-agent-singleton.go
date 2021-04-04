package runtime

import "github.com/bitwormhole/go-wormhole-core/application"

type SingletonComponentAgent struct {
	context application.Context
	ref     application.ComponentInstanceRef
	info    application.ComponentInfo
}
