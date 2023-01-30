package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// A Network represents a network that nodes from NodeGroups can connect to
type Network struct {
	Name string
	Type networktypes.NetworkType
}

// NewNetwork returns a Network of the given NetworkType
func NewNetwork(name string, networkType networktypes.NetworkType) (Network, error) {
	if len(name) == 0 {
		return Network{}, errors.New("name of the Network must consist of at least on character")
	}
	return Network{
		Name: name,
		Type: networkType,
	}, nil
}

// NewDefaultNetwork returns a Network of the default NetworkType
func NewDefaultNetwork(name string) (Network, error) {
	return NewNetwork(name, networktypes.WirelessLAN{}.Default())
}
