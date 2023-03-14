package parser

import (
	"errors"
	"strings"

	logger "github.com/gookit/slog"

	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

type inputExperiment struct {
	Name       any `default:"Experiment"`
	Runs       any `default:"1"`
	RandomSeed uint64
	Duration   any `required:"true"`
	WorldSize  inputWorldSize
	Targets    []string
	Warmup     uint

	Network        []inputNetwork
	NodeGroup      []inputNodeGroup
	EventGenerator []inputEventGenerator
}

type intermediateExperiment struct {
	Name       string
	Runs       uint
	RandomSeed int64
	Warmup     uint
	Duration   uint
}

type inputWorldSize struct {
	Height any `default:"300"`
	Width  any `default:"300"`
}

func parseGeneralExperiment(input inputExperiment) (experiment.Experiment, error) {
	var output experiment.Experiment

	intermediate, err := fillDefaults[inputExperiment, intermediateExperiment](input)
	if err != nil {
		return output, err
	}

	output.Name = intermediate.Name
	output.Runs = intermediate.Runs
	output.Duration = intermediate.Duration
	output.Warmup = intermediate.Warmup
	output.Targets = parseTargets(input.Targets)
	output.RandomSeed = intermediate.RandomSeed
	if intermediate.RandomSeed == 0 {
		logger.Info("No random seed given. Using a generated one.")
		output.RandomSeed = experiment.GenerateRandomSeed()
	}

	worldSize, err := fillDefaults[inputWorldSize, customtypes.Area](input.WorldSize)
	if err != nil {
		return output, err
	}
	output.WorldSize = worldSize

	if intermediate.Runs == 0 {
		return output, errors.New("experiment must have at least one run")
	}

	return output, nil
}

func parseTargets(input []string) []experiment.Target {
	var output []experiment.Target
	for _, targetString := range input {
		switch strings.ToLower(targetString) {
		case "core", "coreemu", "core-emu":
			output = append(output, experiment.TargetCore)
		case "coreemulab", "coreemu-lab", "core-emulab", "core-emu-lab", "clab":
			output = append(output, experiment.TargetCoreEmulab)
		case "the one", "theone", "one":
			output = append(output, experiment.TargetTheOne)
		default:
			logger.Warnf("Unknown target \"%s\". Please check your config.", targetString)
		}
	}
	return output
}
