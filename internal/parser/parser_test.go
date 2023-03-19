package parser_test

import (
	"io/fs"
	"testing"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseSimpleText(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Runs = 1
	Duration = 123

	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"

	[[NodeGroup]]
	Prefix = "users"
	NoNodes = 7
	Networks = ["Wifi"]
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)

	assert.Equal(t, "Testing Experiment", exp.Name)
	assert.EqualValues(t, 1, exp.Runs)
	assert.EqualValues(t, 123, exp.Duration)

	assert.Len(t, exp.Networks, 1)
	assert.Len(t, exp.NodeGroups, 1)

	assert.Equal(t, "Wifi", exp.Networks[0].Name)
	assert.IsType(t, networktypes.WirelessLAN{}, exp.Networks[0].Type)

	assert.Equal(t, "users", exp.NodeGroups[0].Prefix)
	assert.EqualValues(t, 7, exp.NodeGroups[0].NoNodes)
	assert.EqualValues(t, &exp.Networks[0], exp.NodeGroups[0].Networks[0])
}

func TestParseNonExistingFile(t *testing.T) {
	_, err := parser.LoadFromFile("testdata/does-not-exist.toml")
	assert.Error(t, err)
	assert.IsType(t, &fs.PathError{}, err)
}

func TestParseWrongFormattedFile(t *testing.T) {
	_, err := parser.LoadFromFile("testdata/wrong-formatted.toml")
	assert.Error(t, err)
	assert.IsType(t, &toml.DecodeError{}, err)
}

func TestParseSimpleFile(t *testing.T) {
	exp, err := parser.LoadFromFile("testdata/simple.toml")

	assert.NoError(t, err)

	assert.Equal(t, "Testing Experiment", exp.Name)
	assert.EqualValues(t, 1, exp.Runs)
	assert.EqualValues(t, 123, exp.Duration)

	assert.Len(t, exp.Networks, 1)
	assert.Len(t, exp.NodeGroups, 1)

	assert.Equal(t, "Wifi", exp.Networks[0].Name)
	assert.IsType(t, networktypes.WirelessLAN{}, exp.Networks[0].Type)

	assert.Equal(t, "users", exp.NodeGroups[0].Prefix)
	assert.EqualValues(t, 7, exp.NodeGroups[0].NoNodes)
	assert.EqualValues(t, &exp.Networks[0], exp.NodeGroups[0].Networks[0])
}

func TestParseComplexFile(t *testing.T) {
	exp, err := parser.LoadFromFile("testdata/complex.toml")

	assert.NoError(t, err)

	assert.Equal(t, "Complex Experiment", exp.Name)
	assert.EqualValues(t, 5, exp.Runs)
	assert.EqualValues(t, 456789, exp.Duration)
	assert.EqualValues(t, 17, exp.Warmup)
	assert.EqualValues(t, 1337, exp.RandomSeed)
	assert.EqualValues(t, 2000, exp.WorldSize.Height)
	assert.EqualValues(t, 3000, exp.WorldSize.Width)
	assert.EqualValues(t, true, exp.ExternalMovement.Active)
	assert.EqualValues(t, "three_nodes.pos", exp.ExternalMovement.FileName)

	assert.Equal(t, []experiment.Target{experiment.TargetCore, experiment.TargetCoreEmulab, experiment.TargetTheOne}, exp.Targets)

	assert.Len(t, exp.Networks, 6)
	assert.Len(t, exp.NodeGroups, 3)
	assert.Len(t, exp.EventGenerators, 2)

	assert.Equal(t, "Wifi 2.4GHz", exp.Networks[0].Name)
	assert.IsType(t, exp.Networks[0].Type, networktypes.WirelessLAN{})

	assert.Equal(t, "Wifi 5GHz", exp.Networks[1].Name)
	assert.IsType(t, exp.Networks[1].Type, networktypes.WirelessLAN{})
	assert.True(t, assert.ObjectsAreEqual(networktypes.WirelessLAN{
		Bandwidth:   1300000000,
		Range:       80,
		Jitter:      2,
		Delay:       1000,
		Loss:        0.1,
		Promiscuous: true,
	}, exp.Networks[1].Type))

	assert.Equal(t, "LTE", exp.Networks[2].Name)
	assert.IsType(t, networktypes.Wireless{}, exp.Networks[2].Type)
	assert.True(t, assert.ObjectsAreEqual(networktypes.Wireless{
		Movement:       true,
		Bandwidth:      250000000,
		Range:          600,
		Jitter:         5,
		Delay:          2000,
		LossInitial:    0.1,
		LossFactor:     3.0,
		LossStartRange: 100.0,
	}, exp.Networks[2].Type))

	assert.Equal(t, "Hubbed", exp.Networks[3].Name)
	assert.IsType(t, networktypes.Hub{}, exp.Networks[3].Type)

	assert.Equal(t, "Switched", exp.Networks[4].Name)
	assert.IsType(t, networktypes.Switch{}, exp.Networks[4].Type)

	assert.Equal(t, "EMANE", exp.Networks[5].Name)
	assert.IsType(t, networktypes.Emane{}, exp.Networks[5].Type)

	assert.Equal(t, "laptops", exp.NodeGroups[0].Prefix)
	assert.EqualValues(t, 7, exp.NodeGroups[0].NoNodes)
	assert.EqualValues(t, experiment.NodeTypePC, exp.NodeGroups[0].NodesType)
	assert.EqualValues(t, []*experiment.Network{&exp.Networks[0], &exp.Networks[1], &exp.Networks[3], &exp.Networks[4]}, exp.NodeGroups[0].Networks)
	assert.IsType(t, movementpatterns.Static{}, exp.NodeGroups[0].MovementModel)

	assert.Equal(t, "IoT", exp.NodeGroups[1].Prefix)
	assert.EqualValues(t, 12, exp.NodeGroups[1].NoNodes)
	assert.EqualValues(t, experiment.NodeTypeRouter, exp.NodeGroups[1].NodesType)
	assert.EqualValues(t, []*experiment.Network{&exp.Networks[0], &exp.Networks[5]}, exp.NodeGroups[1].Networks)

	assert.Equal(t, "smartphones", exp.NodeGroups[2].Prefix)
	assert.EqualValues(t, 12, exp.NodeGroups[2].NoNodes)
	assert.EqualValues(t, experiment.NodeTypePC, exp.NodeGroups[2].NodesType)
	assert.EqualValues(t, []*experiment.Network{&exp.Networks[0], &exp.Networks[1], &exp.Networks[2]}, exp.NodeGroups[2].Networks)
	assert.IsType(t, movementpatterns.RandomWaypoint{}, exp.NodeGroups[2].MovementModel)
	assert.True(t, assert.ObjectsAreEqual(movementpatterns.RandomWaypoint{
		MinSpeed: 5,
		MaxSpeed: 7,
		MaxPause: 2,
	}, exp.NodeGroups[2].MovementModel))

	assert.Equal(t, "EVG1", exp.EventGenerators[0].Name)
	assert.IsType(t, eventgeneratortypes.MessageBurstGenerator{}, exp.EventGenerators[0].Type)
	assert.True(t, assert.ObjectsAreEqual(eventgeneratortypes.MessageBurstGenerator{
		Prefix: "i",
		Interval: customtypes.Interval{
			From: 3,
			To:   4,
		},
		Size: customtypes.Interval{
			From: 17,
			To:   35,
		},
		Hosts: customtypes.Interval{
			From: 9,
			To:   12,
		},
		ToHosts: customtypes.Interval{
			From: 0,
			To:   9,
		},
	}, exp.EventGenerators[0].Type))

	assert.Equal(t, "EVG2", exp.EventGenerators[1].Name)
	assert.IsType(t, eventgeneratortypes.MessageEventGenerator{}, exp.EventGenerators[1].Type)
	assert.Equal(t, "ii", exp.EventGenerators[1].Type.(eventgeneratortypes.MessageEventGenerator).Prefix)
}
