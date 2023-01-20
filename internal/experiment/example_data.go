package experiment

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// GetExampleExperiment returns a Experiment loaded with example values.
func GetExampleExperiment() Experiment {
	var nodegroups []NodeGroup
	ng, _ := NewNodeGroup("n", 1)
	nodegroups = append(nodegroups, ng)
	ng, _ = NewNodeGroup("p", 29)
	nodegroups = append(nodegroups, ng)

	ng, _ = NewNodeGroup("x", 17)
	ng.MovementModel = movementpatterns.Static{}
	nodegroups = append(nodegroups, ng)

	var ExampleExperiment = Experiment{
		Name:    "Example Experiment",
		Runs:    1,
		Targets: []Target{TargetTheOne, TargetCore},

		RandomSeed: GenerateRandomSeed(),

		Duration: 120,
		WorldSize: customtypes.Area{
			Height: 800,
			Width:  800,
		},

		NodeGroups: nodegroups,
	}
	return ExampleExperiment
}
