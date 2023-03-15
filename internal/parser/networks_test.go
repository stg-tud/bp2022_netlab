package parser_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestUnknownNetworkType(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[Network]]
	Name = "Wifi"
	Type = "bluetooth"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: network type \"bluetooth\" not found"), err)
}

func TestDuplicateNetworkNames(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123

	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"

	[[Network]]
	Name = "Wifi"
	Type = "wireless"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("a network with the name \"Wifi\" already exists"), err)
}
