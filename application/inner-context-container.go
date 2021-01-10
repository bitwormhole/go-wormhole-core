package application

import (
	"github.com/bitwormhole/go-wormhole-core/collection"
)

// innerContextContainer 定义了一个通用的 Context 的内部结构
type innerContextContainer struct {
	componentTable map[string]ComponentHolder

	components  Components
	arguments   collection.Arguments
	attributes  collection.Attributes
	environment collection.Environment
	properties  collection.Properties
	parameters  collection.Parameters
	resources   collection.Resources
	releasePool collection.ReleasePool

	contextFacade Context
}
