package parser_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestUnknownNodesType(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[NodeGroup]]
	Prefix = "users"
	NodesType = "Laptop"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing node group 0: node type \"Laptop\" not found"), err)
}

func TestDuplicateNodeGroupPrefixes(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[NodeGroup]]
	Prefix = "users"
	NoNodes = 7

	[[NodeGroup]]
	Prefix = "users"
	NoNodes = 3
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("a node group with the prefix \"users\" already exists"), err)
}

func TestAssignedNetworkDoesNotExist(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[NodeGroup]]
	Prefix = "users"
	NoNodes = 7
	Networks = ["Wifi"]
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing node group 0: network \"Wifi\" not found"), err)
}

func TestAssignedToSameNetworkMultipleTimes(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"

	[[NodeGroup]]
	Prefix = "users"
	NoNodes = 7
	Networks = ["Wifi", "Wifi"]
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)
	assert.Equal(t, []*experiment.Network{&exp.Networks[0]}, exp.NodeGroups[0].Networks)
}

func TestMovementModels(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[NodeGroup]]
	Prefix = "static"
	NoNodes = 7
	MovementModel = "static"

	[[NodeGroup]]
	Prefix = "rwp"
	NoNodes = 7
	MovementModel = "random waypoint"
	MinSpeed = 3
	MaxSpeed = 4
	MaxPause = 5

	[[NodeGroup]]
	Prefix = "slaw"
	NoNodes = 7
	MovementModel = "slaw"
	NumberOfWaypoints = 21
	MinPause = 3
	MaxPause = 4
	LevyExponent = 22.2
	HurstParameter = 33.3
	DistanceWeight = 19.8
	ClusteringRange = 7.8
	ClusterRatio = 8
	WaypointRatio = 12

	[[NodeGroup]]
	Prefix = "swim"
	NoNodes = 7
	MovementModel = "swim"
	Radius = 19.1
	CellDistanceWeight = 32.1
	NodeSpeedMultiplier = 7.6
	WaitingTimeExponent = 0.9
	WaitingTimeUpperBound = 22.2

	[[NodeGroup]]
	Prefix = "smooth"
	NoNodes = 7
	MovementModel = "smooth"
	Range = 8
	Clusters = 9
	Alpha = 99.9
	MinFlight = 7
	MaxFlight = 12
	Beta = 11.1
	MinPause = 3
	MaxPause = 4
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)

	assert.IsType(t, movementpatterns.Static{}, exp.NodeGroups[0].MovementModel)
	assert.True(t, assert.ObjectsAreEqual(movementpatterns.Static{}, exp.NodeGroups[0].MovementModel))

	assert.IsType(t, movementpatterns.RandomWaypoint{}, exp.NodeGroups[1].MovementModel)
	assert.True(t, assert.ObjectsAreEqual(movementpatterns.RandomWaypoint{
		MinSpeed: 3,
		MaxSpeed: 4,
		MaxPause: 5,
	}, exp.NodeGroups[1].MovementModel))

	assert.IsType(t, movementpatterns.SLAW{}, exp.NodeGroups[2].MovementModel)
	assert.True(t, assert.ObjectsAreEqual(movementpatterns.SLAW{
		NumberOfWaypoints: 21,
		MinPause:          3,
		MaxPause:          4,
		LevyExponent:      22.2,
		HurstParameter:    33.3,
		DistanceWeight:    19.8,
		ClusteringRange:   7.8,
		ClusterRatio:      8,
		WaypointRatio:     12,
	}, exp.NodeGroups[2].MovementModel))

	assert.IsType(t, movementpatterns.SWIM{}, exp.NodeGroups[3].MovementModel)
	assert.True(t, assert.ObjectsAreEqual(movementpatterns.SWIM{
		Radius:                19.1,
		CellDistanceWeight:    32.1,
		NodeSpeedMultiplier:   7.6,
		WaitingTimeExponent:   0.9,
		WaitingTimeUpperBound: 22.2,
	}, exp.NodeGroups[3].MovementModel))

	assert.IsType(t, movementpatterns.SMOOTH{}, exp.NodeGroups[4].MovementModel)
	assert.True(t, assert.ObjectsAreEqual(movementpatterns.SMOOTH{
		Range:     8,
		Clusters:  9,
		Alpha:     99.9,
		MinFlight: 7,
		MaxFlight: 12,
		Beta:      11.1,
		MinPause:  3,
		MaxPause:  4,
	}, exp.NodeGroups[4].MovementModel))
}

func TestUnknownMovementModel(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[NodeGroup]]
	Prefix = "users"
	NoNodes = 7
	MovementModel = "circular"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing node group 0: movement pattern \"circular\" not found"), err)
}
