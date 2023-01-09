package experiment

import (
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// A NodeGroup represents a group of dependent nodes sharing properties
// such as a MovementModel or network settings.
type NodeGroup struct {
	Prefix  string
	NoNodes int

	MovementModel movementpatterns.MovementPattern

	IPv4Net  string
	IPv4Mask int
	IPv6Net  string
	IPv6Mask int

	NetworkType NetworkType
	Bandwidth   int
	Range       int
	Jitter      int
	Delay       int
	Error       int
	Promiscuous int
}

type NetworkType int

const (
	WIRELESS_LAN NetworkType = iota
)

func (nt NetworkType) String() string {
	switch nt {
	case WIRELESS_LAN:
		return "WIRELESS_LAN"

	default:
		return ""
	}
}

var defaultValues = NodeGroup{
	MovementModel: movementpatterns.RandomWaypoint{
		MinSpeed: 123,
		MaxSpeed: 456,
		MaxPause: 0,
	},

	IPv4Net:  "10.0.0.0",
	IPv4Mask: 24,
	IPv6Net:  "2001::",
	IPv6Mask: 120,

	NetworkType: WIRELESS_LAN,
	Range:       180,
	Bandwidth:   54000000,
	Jitter:      0,
	Delay:       20000,
	Error:       0,
	Promiscuous: 0,
}

// NewNodeGroup returns a new NodeGroup loaded with default values.
func NewNodeGroup(prefix string, noNodes int) NodeGroup {
	out := defaultValues
	out.Prefix = prefix
	out.NoNodes = noNodes
	return out
}
