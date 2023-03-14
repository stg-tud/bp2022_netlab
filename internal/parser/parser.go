package parser

import (
	"os"

	logger "github.com/gookit/slog"
	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func LoadFromFile(path string) (experiment.Experiment, error) {
	logger.Info("Loading experiment from file")

	var output experiment.Experiment
	var tomlIn inputExperiment = inputExperiment{}

	buf, err := os.ReadFile(path)
	if err != nil {
		return output, err
	}
	err = toml.Unmarshal(buf, &tomlIn)
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

	logger.Info("Finished loading from file")
	return output, nil
}
