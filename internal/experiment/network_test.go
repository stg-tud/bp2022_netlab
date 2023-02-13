package experiment_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

func TestNetworkWithoutName(t *testing.T) {
	_, error1 := experiment.NewDefaultNetwork("")
	_, error2 := experiment.NewNetwork("", networktypes.Emane{})
	if error1 == nil || error2 == nil {
		t.Fatal("Networks without name should not be allowed!")
	}
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
		if err != nil {
			t.Fatalf("Error creating new '%s' network: %s", networkType.String(), err)
		}
		if network_under_test.Name != networkName {
			t.Fatalf("Network has wrong name '%s', expected '%s'!", network_under_test.Name, networkName)
		}
		if network_under_test.Type != networkType {
			t.Fatalf("Network has wrong type '%s', expected '%s'!", network_under_test.Type, networkType)
		}
	}
}
