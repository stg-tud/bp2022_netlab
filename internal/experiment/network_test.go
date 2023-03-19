package experiment_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
	"github.com/stretchr/testify/assert"
)

func TestNetworkWithoutName(t *testing.T) {
	_, err := experiment.NewDefaultNetwork("")
	assert.Error(t, err)
	assert.Equal(t, errors.New("name of the Network must consist of at least on character"), err)
	_, err = experiment.NewNetwork("", networktypes.Emane{})
	assert.Error(t, err)
	assert.Equal(t, errors.New("name of the Network must consist of at least on character"), err)
}

func TestNewNetwork(t *testing.T) {
	networkTypes := []networktypes.NetworkType{
		networktypes.Wireless{}.Default(),
		networktypes.WirelessLAN{}.Default(),
		networktypes.Emane{}.Default(),
		networktypes.Switch{}.Default(),
		networktypes.Hub{}.Default(),
	}
	for _, networkType := range networkTypes {
		networkName := fmt.Sprintf("network_under_test_%s", strings.ToLower(networkType.String()))
		network_under_test, err := experiment.NewNetwork(networkName, networkType)
		assert.NoError(t, err)
		assert.Equal(t, networkName, network_under_test.Name)
		assert.Equal(t, networkType, network_under_test.Type)
	}
}
