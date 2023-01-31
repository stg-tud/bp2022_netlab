package experiment

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// GetExampleExperiment returns a Experiment loaded with example values.
func GetExampleExperiment() Experiment {
	var networks []Network
	net, _ := NewNetwork("switched", networktypes.Switch{}.Default())
	networks = append(networks, net)
	net, _ = NewNetwork("wifi", networktypes.WirelessLAN{}.Default())
	networks = append(networks, net)
	net, _ = NewNetwork("bluetooth", networktypes.Wireless{}.Default())
	networks = append(networks, net)

	var nodegroups []NodeGroup
	ng, _ := NewNodeGroup("n", 1)
	ng.Networks = append(ng.Networks, &networks[0])
	nodegroups = append(nodegroups, ng)

	ng, _ = NewNodeGroup("p", 29)
	ng.Networks = append(ng.Networks, &networks[1])
	ng.Networks = append(ng.Networks, &networks[2])
	nodegroups = append(nodegroups, ng)

	ng, _ = NewNodeGroup("x", 17)
	ng.MovementModel = movementpatterns.Static{}
	ng.Networks = append(ng.Networks, &networks[2])
	ng.NodesType = NodeTypePC
	nodegroups = append(nodegroups, ng)

	var ExampleExperiment = Experiment{
		Name:    "Example Experiment",
		Runs:    1,
		Targets: []Target{TargetTheOne, TargetCore},

		RandomSeed: GenerateRandomSeed(),

		Duration: 120,
		WorldSize: customtypes.Area{
			Height: 800,
			Width:  1000,
		},

		Networks:   networks,
		NodeGroups: nodegroups,
	}
	return ExampleExperiment
}
