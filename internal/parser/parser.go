package parser

import (
	"os"

	logger "github.com/gookit/slog"
	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// Loads a file at the given path and tries to parse its contents to an experiment.Experiment
func LoadFromFile(path string) (experiment.Experiment, error) {
	logger.Info("Loading experiment from file")
	var output experiment.Experiment = experiment.Experiment{}

	buf, err := os.ReadFile(path)
	if err != nil {
		return output, err
	}

	return ParseText(buf)
}

// Parses a given input text to an experiment.Experiment
func ParseText(input []byte) (experiment.Experiment, error) {
	logger.Info("Loading experiment from string")
	var output experiment.Experiment = experiment.Experiment{}
	var tomlIn inputExperiment = inputExperiment{}

	err := toml.Unmarshal(input, &tomlIn)
	if err != nil {
		return output, err
	}

	output, err = parseGeneralExperiment(tomlIn)
	if err != nil {
		return output, err
	}

	networks, err := parseNetworks(tomlIn.Network)
	if err != nil {
		return output, err
	}
	output.Networks = networks

	nodeGroups, err := parseNodeGroups(tomlIn.NodeGroup, &output)
	if err != nil {
		return output, err
	}
	output.NodeGroups = nodeGroups

	eventGenerators, err := parseEventGenerators(tomlIn.EventGenerator)
	if err != nil {
		return output, err
	}
	output.EventGenerators = eventGenerators

	return output, nil
}
