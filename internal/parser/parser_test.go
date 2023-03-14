package parser_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestParseText(t *testing.T) {
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

	experiment, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)

	assert.Equal(t, experiment.Name, "Testing Experiment")
	assert.EqualValues(t, experiment.Runs, 1)
	assert.EqualValues(t, experiment.Duration, 123)

	assert.Len(t, experiment.Networks, 1)
	assert.Len(t, experiment.NodeGroups, 1)

	assert.Equal(t, experiment.Networks[0].Name, "Wifi")
	assert.IsType(t, networktypes.WirelessLAN{}, experiment.Networks[0].Type)

	assert.Equal(t, experiment.NodeGroups[0].Prefix, "users")
	assert.EqualValues(t, experiment.NodeGroups[0].NoNodes, 7)
	assert.EqualValues(t, experiment.NodeGroups[0].Networks[0], &experiment.Networks[0])

	outputgenerators.Core{}.Generate(experiment)
}
