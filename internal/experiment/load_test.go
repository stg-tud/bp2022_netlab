package experiment_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

func TestLoad(t *testing.T) {

	actual, err := experiment.LoadFromFile("testdata/load_test.toml")
	if err != nil {
		t.Fatal(err)
	}

	if actual.Duration != 123456 {
		t.Fatal("Wrong duration")
	}
	expect, err := experiment.NewDefaultEventGenerator("MessageEventGenerator")
	if err != nil {
		t.Fatal(err)
	}
	if actual.EventGenerators[0] != expect {
		t.Fatal("Wrong eventgenerators[0]")
	}
	if actual.Name != "Testing Experiment" {
		t.Fatal("Wrong experiment name")
	}
	if actual.RandomSeed != 1673916419715 {
		t.Fatal("Wrong Randomseed")
	}
	if actual.NodeGroups[0].MovementModel.String() != "Static" ||
		actual.NodeGroups[2].MovementModel.String() != "SLAW: Self-similar Least Action Walk" ||
		actual.NodeGroups[3].MovementModel.String() != "Random Waypoint" {
		t.Fatal("Wrong movementModel")
	}
	if actual.NodeGroups[4].NodesType.String() != "PC" {
		t.Fatal("Wrong nodetype")
	}
	if len(actual.NodeGroups[1].Networks) != 2 {
		t.Fatal("Wrong networks")
	}
	expected := movementpatterns.RandomWaypoint{
		MinSpeed: 1,
		MaxSpeed: 2,
		MaxPause: 17,
	}
	if actual.NodeGroups[5].MovementModel != expected {
		t.Fatal("Wrong custom movementModel")
	}
	net, _ := experiment.NewNetwork("wireless_lan", networktypes.WirelessLAN{}.Default())
	if actual.Networks[0] != net {
		t.Fatal("Wrong network at [0]")
	}
	changedWifi := networktypes.WirelessLAN{}.Default()
	changedWifi.Bandwidth = 17
	changedWifi.Promiscuous = true
	net, _ = experiment.NewNetwork("changed_wifi", changedWifi)
	if actual.Networks[2] != net {
		t.Fatal("Wrong network at [2]")
	}
	if actual.NodeGroups[6].NodesType.String() != "Router" {
		t.Fatal("Wrong default network for Nodegroup")
	}

}
