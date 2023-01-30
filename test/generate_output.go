package main

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func main() {
	// og := outputgenerators.Bonnmotion{}
	og2 := outputgenerators.Core{}
	// exampleExperiment := experiment.GetExampleExperiment()

	// exampleExperiment := experiment.Experiment{
	// 	Name:     "Test",
	// 	Runs:     1,
	// 	Targets:  []experiment.Target{experiment.TargetCore},
	// 	Duration: 300,
	// 	WorldSize: customtypes.Area{
	// 		Height: 750,
	// 		Width:  1000,
	// 	},
	// 	RandomSeed: experiment.GenerateRandomSeed(),
	// 	NodeGroups: []experiment.NodeGroup{
	// 		{
	// 			Prefix:  "n",
	// 			NoNodes: 4,
	// 			MovementModel: movementpatterns.RandomWaypoint{
	// 				MinSpeed: 5,
	// 				MaxSpeed: 10,
	// 				MaxPause: 2,
	// 			},
	// 			NodesType:   experiment.NodeTypeRouter,
	// 			NetworkType: networktypes.WirelessLAN{}.Default(),
	// 		},
	// 		{
	// 			Prefix:  "p",
	// 			NoNodes: 5,
	// 			MovementModel: movementpatterns.RandomWaypoint{
	// 				MinSpeed: 50,
	// 				MaxSpeed: 100,
	// 				MaxPause: 1,
	// 			},
	// 			NodesType:   experiment.NodeTypeRouter,
	// 			NetworkType: networktypes.WirelessLAN{}.Default(),
	// 		},
	// 	},
	// }

	networks := []experiment.Network{
		{
			Name: "wifi",
			Type: networktypes.WirelessLAN{}.Default(),
		},
		{
			Name: "switch",
			Type: networktypes.Switch{}.Default(),
		},
	}
	exampleExperiment := experiment.Experiment{
		Name:     "Test",
		Runs:     1,
		Targets:  []experiment.Target{experiment.TargetCore},
		Duration: 300,
		WorldSize: customtypes.Area{
			Height: 750,
			Width:  1000,
		},
		RandomSeed: experiment.GenerateRandomSeed(),
		Networks:   networks,
		NodeGroups: []experiment.NodeGroup{
			{
				Prefix:        "n",
				NoNodes:       4,
				MovementModel: movementpatterns.Static{},
				NodesType:     experiment.NodeTypeRouter,
				Networks:      []*experiment.Network{&networks[0]},
			},
			{
				Prefix:        "p",
				NoNodes:       2,
				MovementModel: movementpatterns.Static{},
				NodesType:     experiment.NodeTypeRouter,
				Networks:      []*experiment.Network{&networks[0], &networks[1]},
			},
		},
	}

	// og.Generate(exampleExperiment)
	og2.Generate(exampleExperiment)
}
