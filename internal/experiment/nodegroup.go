package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// A NodeGroup represents a group of dependent nodes sharing properties
// such as a MovementModel or network settings.
type NodeGroup struct {
	Prefix  string
	NoNodes uint

	MovementModel movementpatterns.MovementPattern

	NodesType NodeType

	NetworkType networktypes.NetworkType
}

var defaultValues = NodeGroup{
	MovementModel: movementpatterns.RandomWaypoint{
		MinSpeed: 123,
		MaxSpeed: 456,
		MaxPause: 0,
	},

	NodesType: NODE_TYPE_ROUTER,

	NetworkType: networktypes.WirelessLAN{}.Default(),
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
