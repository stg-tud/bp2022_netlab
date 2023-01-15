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

	nodegroups[2].NetworkType = networktypes.Switch{}.Default()
	nodegroups[2].MovementModel = movementpatterns.Static{}

	changedWifi := networktypes.WirelessLAN{}.Default()
	changedWifi.Bandwidth = 17
	changedWifi.Promiscuous = true
	nodegroups[3].NetworkType = changedWifi

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
