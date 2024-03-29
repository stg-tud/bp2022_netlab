package outputgenerators_test

import (
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// The relative name of the folder containing the expected test outputs and additional testing data.
const TestDataFolder = "testdata"

// GetTestingExperiment returns an Experiment loaded with values used for unit tests.
func GetTestingExperiment() experiment.Experiment {
	var networks []experiment.Network
	net, _ := experiment.NewNetwork("wireless_lan", networktypes.WirelessLAN{}.Default())
	networks = append(networks, net)
	net, _ = experiment.NewNetwork("switch", networktypes.Switch{}.Default())
	networks = append(networks, net)

	changedWifi := networktypes.WirelessLAN{}.Default()
	changedWifi.Bandwidth = 17
	changedWifi.Promiscuous = true
	net, _ = experiment.NewNetwork("changed_wifi", changedWifi)
	networks = append(networks, net)

	net, _ = experiment.NewNetwork("hub", networktypes.Hub{}.Default())
	networks = append(networks, net)
	net, _ = experiment.NewNetwork("emane", networktypes.Emane{}.Default())
	networks = append(networks, net)
	net, _ = experiment.NewNetwork("wireless", networktypes.Wireless{}.Default())
	networks = append(networks, net)

	var eventgenerators []experiment.EventGenerator
	evg, _ := experiment.NewEventGenerator("MessageEventGenerator", eventgeneratortypes.MessageEventGenerator{}.Default())
	eventgenerators = append(eventgenerators, evg)
	changedBurst := eventgeneratortypes.MessageBurstGenerator{}.Default()
	changedBurst.Interval = customtypes.Interval{From: 25, To: 35}
	evg, _ = experiment.NewEventGenerator("MessageBurstGenerator", changedBurst)
	eventgenerators = append(eventgenerators, evg)

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

	nodegroups[0].Networks = []*experiment.Network{&networks[0]}
	nodegroups[0].PredefinedPosition = true
	nodegroups[0].Position = customtypes.Position{
		X: 40,
		Y: 55,
	}

	nodegroups[1].Networks = []*experiment.Network{&networks[0], &networks[1]}

	nodegroups[2].Networks = []*experiment.Network{&networks[1]}
	nodegroups[2].MovementModel = movementpatterns.RandomWaypoint{}.Default()

	nodegroups[3].Networks = []*experiment.Network{&networks[2]}

	nodegroups[4].NodesType = experiment.NodeTypePC
	nodegroups[4].Networks = []*experiment.Network{&networks[3]}

	nodegroups[5].Networks = []*experiment.Network{&networks[4]}
	nodegroups[5].MovementModel = movementpatterns.RandomWaypoint{
		MinSpeed: 123,
		MaxSpeed: 456,
		MaxPause: 789,
	}

	nodegroups[6].Networks = []*experiment.Network{&networks[5]}

	exp := experiment.Experiment{
		Name:    "Testing Experiment",
		Runs:    5,
		Targets: []experiment.Target{experiment.TargetCore, experiment.TargetTheOne},

		RandomSeed: 1673916419715,
		Warmup:     5,
		Duration:   43,
		Automator:  "three_nodes.pos",
		WorldSize: customtypes.Area{
			Height: 170,
			Width:  240,
		},

		Networks:        networks,
		NodeGroups:      nodegroups,
		EventGenerators: eventgenerators,
	}

	return exp
}
