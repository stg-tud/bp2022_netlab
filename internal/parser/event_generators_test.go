package parser_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestUnknownEventGeneratorType(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[EventGenerator]]
	Name = "EVG"
	Type = "Random"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing event generator 0: event generator type \"Random\" not found"), err)
}

func TestEventGeneratorDuplicateName(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[EventGenerator]]
	Name = "EVG"
	Type = "Burst"

	[[EventGenerator]]
	Name = "EVG"
	Type = "Event"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("an event generator with the name \"EVG\" already exists"), err)
}

func TestEventGeneratorDuplicatePrefix(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[EventGenerator]]
	Name = "EVG1"
	Type = "Burst"
	Prefix = "EVG1"

	[[EventGenerator]]
	Name = "EVG2"
	Type = "Event"
	Prefix = "EVG1"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing event generator 1: an event generator with the prefix \"EVG1\" already exists"), err)
}

func TestEventGeneratorPrefixFromName(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[EventGenerator]]
	Name = "EVG1"
	Type = "Burst"

	[[EventGenerator]]
	Name = "EVG2"
	Type = "Event"
	Prefix = "EVG2"
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)
	assert.Equal(t, "EVG1", exp.EventGenerators[0].Type.(eventgeneratortypes.MessageBurstGenerator).Prefix)
	assert.Equal(t, "EVG2", exp.EventGenerators[1].Type.(eventgeneratortypes.MessageEventGenerator).Prefix)
}
