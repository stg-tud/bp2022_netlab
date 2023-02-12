package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// A Network represents a network that nodes from NodeGroups can connect to
type Network struct {
	Name           string
	Type           networktypes.NetworkType
	NrofInterfaces uint
}

// NewNetwork returns a Network of the given NetworkType
func NewNetwork(name string, networkType networktypes.NetworkType, nrofInterfaces uint) (Network, error) {
	if len(name) == 0 {
		return Network{}, errors.New("name of the Network must consist of at least on character")
	}
	return Network{
		Name:           name,
		Type:           networkType,
		NrofInterfaces: nrofInterfaces,
	}, nil
}

// NewDefaultNetwork returns a Network of the default NetworkType
func NewDefaultNetwork(name string) (Network, error) {
	return NewNetwork(name, networktypes.WirelessLAN{}.Default(), 1)
}
