package experiment

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// GetExampleExperiment returns a Experiment loaded with example values.
func GetExampleExperiment() Experiment {
	var nodegroups []NodeGroup
	nodegroups = append(nodegroups, NewNodeGroup("n", 1))

	ng2 := NewNodeGroup("p", 29)
	ng2.NetworkType = networktypes.Switch{}.Default()
	nodegroups = append(nodegroups, ng2)

	ng3 := NewNodeGroup("x", 17)
	ng3.MovementModel = movementpatterns.Static{}
	ng3.NetworkType = networktypes.Switch{}.Default()
	ng3.NodesType = NODE_TYPE_PC
	nodegroups = append(nodegroups, ng3)

	var ExampleExperiment = Experiment{
		Name:    "Example Experiment",
		Runs:    1,
		Targets: []Target{TARGET_THEONE, TARGET_CORE},

		Duration: 120,
		WorldSize: customtypes.Area{
			Height: 800,
			Width:  1000,
		},

		NodeGroups: nodegroups,
	}
	return ExampleExperiment
}
