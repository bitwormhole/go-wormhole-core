package runtime

import (
	"github.com/bitwormhole/go-wormhole-core/application"
	"github.com/bitwormhole/go-wormhole-core/collection"
)

type ContextCore struct {
	Context    application.Context
	Components application.Components

	ReleasePool collection.ReleasePool

	Arguments   collection.Arguments
	Attributes  collection.Attributes
	Environment collection.Environment
	Properties  collection.Properties
	Parameters  collection.Parameters
	Resources   collection.Resources

	ComponentAgentTable map[string]application.ComponentAgent
}
