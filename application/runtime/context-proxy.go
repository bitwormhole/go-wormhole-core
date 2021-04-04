package runtime

import "github.com/bitwormhole/go-wormhole-core/application"

type ContextProxy struct {
	current  application.Context
	building application.Context
	running  application.Context
}
