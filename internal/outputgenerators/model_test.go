package outputgenerators_test

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// The relative name of the folder containing the expected test outputs and additional testing data.
const TestDataFolder = "testdata"

// GetTestingExperiment returns an Experiment loaded with values used for unit tests.
func GetTestingExperiment() experiment.Experiment {
	var nodegroups []experiment.NodeGroup
	ng, _ := experiment.NewNodeGroup("a", 1)
	nodegroups = append(nodegroups, ng)
	ng, _ = experiment.NewNodeGroup("b", 2)
	nodegroups = append(nodegroups, ng)
	ng, _ = experiment.NewNodeGroup("c", 3)
	nodegroups = append(nodegroups, ng)
	ng, _ = experiment.NewNodeGroup("d", 4)
	nodegroups = append(nodegroups, ng)
	ng, _ = experiment.NewNodeGroup("e", 5)
	nodegroups = append(nodegroups, ng)
	ng, _ = experiment.NewNodeGroup("f", 6)
	nodegroups = append(nodegroups, ng)
	ng, _ = experiment.NewNodeGroup("g", 7)
	nodegroups = append(nodegroups, ng)

	nodegroups[2].MovementModel = movementpatterns.Static{}

	nodegroups[4].MovementModel = movementpatterns.Static{}

	nodegroups[5].MovementModel = movementpatterns.RandomWaypoint{
		MinSpeed: 1,
		MaxSpeed: 2,
		MaxPause: 17,
	}

	exp := experiment.Experiment{
		Name:    "Testing Experiment",
		Runs:    5,
		Targets: []experiment.Target{experiment.TargetCore, experiment.TargetTheOne},

		RandomSeed: 1673916419715,

		Duration: 123456,
		WorldSize: customtypes.Area{
			Height: 170,
			Width:  240,
		},

		NodeGroups: nodegroups,
	}

	return exp
}
