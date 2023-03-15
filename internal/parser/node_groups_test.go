package parser_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
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
