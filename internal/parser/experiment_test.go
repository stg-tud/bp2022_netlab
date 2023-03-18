package parser_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestExperimentWithZeroRuns(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Runs = 0
	Duration = 123
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("experiment must have at least one run"), err)
}

func TestRandomSeedGeneration(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123
	`

	exp, err := parser.ParseText([]byte(toml))
	assert.NoError(t, err)
	assert.Greater(t, exp.RandomSeed, int64(0), "RandomSeed not generated correctly!")
}

func TestWrongTarget(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 123
	Targets = ["CORE", "netlab"]
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err) // With a unknown target, generation should succeed anyways

	assert.Equal(t, []experiment.Target{experiment.TargetCore}, exp.Targets)
}
