package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// A NodeGroup represents a group of dependent nodes sharing properties
// such as a MovementModel or network settings.
type NodeGroup struct {
	Prefix          string
	NoNodes         uint
	DefaultPosition bool
	Position        customtypes.Position
	MovementModel   movementpatterns.MovementPattern

	NodesType NodeType

	Networks []*Network
}

var defaultValues = NodeGroup{
	MovementModel: movementpatterns.Static{},

	NodesType: NodeTypeRouter,

	Networks: []*Network{},
}

// NewNodeGroup returns a new NodeGroup loaded with default values.
func NewNodeGroup(prefix string, noNodes uint) (NodeGroup, error) {
	if len(prefix) == 0 {
		return NodeGroup{}, errors.New("prefix must consist of at least one character")
	}
	if noNodes <= 0 {
		return NodeGroup{}, errors.New("NodeGroup must at least consist of one node")
	}

	out := defaultValues
	out.Prefix = prefix
	out.NoNodes = noNodes
	return out, nil
}
