package experiment

import (
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// A NodeGroup represents a group of dependent nodes sharing properties
// such as a MovementModel or network settings.
type NodeGroup struct {
	Prefix  string
	NoNodes int

	MovementModel movementpatterns.MovementPattern

	NetworkType networktypes.NetworkType
}

var defaultValues = NodeGroup{
	MovementModel: movementpatterns.RandomWaypoint{
		MinSpeed: 123,
		MaxSpeed: 456,
		MaxPause: 0,
	},

	NetworkType: networktypes.WirelessLAN{}.Default(),
}

// NewNodeGroup returns a new NodeGroup loaded with default values.
func NewNodeGroup(prefix string, noNodes int) NodeGroup {
	if len(prefix) == 0 {
		panic("NodeGroup prefix must contain at least one letter!")
	}
	out := defaultValues
	out.Prefix = prefix
	out.NoNodes = noNodes
	return out
}
