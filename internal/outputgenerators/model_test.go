package outputgenerators_test

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

const TESTDATA_FOLDER = "testdata"

func GetTestingExperiment() experiment.Experiment {
	var nodegroups []experiment.NodeGroup
	nodegroups = append(nodegroups, experiment.NewNodeGroup("a", 1))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("b", 2))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("c", 3))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("d", 4))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("e", 5))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("f", 6))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("g", 7))

	nodegroups[2].NetworkType = networktypes.Switch{}.Default()
	nodegroups[2].MovementModel = movementpatterns.Static{}

	changedWifi := networktypes.WirelessLAN{}.Default()
	changedWifi.Bandwidth = 17
	changedWifi.Promiscuous = true
	nodegroups[3].NetworkType = changedWifi

	nodegroups[4].NodesType = experiment.NODE_TYPE_PC
	nodegroups[4].NetworkType = networktypes.Hub{}.Default()
	nodegroups[4].MovementModel = movementpatterns.Static{}

	nodegroups[5].NetworkType = networktypes.Emane{}.Default()
	nodegroups[5].MovementModel = movementpatterns.RandomWaypoint{
		MinSpeed: 1,
		MaxSpeed: 2,
		MaxPause: 17,
	}

	nodegroups[6].NetworkType = networktypes.Wireless{}.Default()

	exp := experiment.Experiment{
		Name:    "Testing Experiment",
		Runs:    5,
		Targets: []experiment.Target{experiment.TARGET_CORE, experiment.TARGET_THEONE},

		Duration: 123456,
		WorldSize: customtypes.Area{
			Height: 170,
			Width:  240,
		},

		NodeGroups: nodegroups,
	}

	return exp
}
