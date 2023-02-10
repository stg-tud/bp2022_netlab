package main

import (
	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/logging"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func main() {
	logging.Init()
	logger.Info("Starting")

	exampleExperiment := experiment.GetExampleExperiment()

	logger.Info("Using random seed", exampleExperiment.RandomSeed)

	outputGenerators := []outputgenerators.OutputGenerator{
		outputgenerators.Core{},
		outputgenerators.Bonnmotion{},
		outputgenerators.Debug{},
		outputgenerators.Theone{},
	}

	for _, outputGenerator := range outputGenerators {
		outputGenerator.Generate(exampleExperiment)
	}

	logger.Info("Finished")
}
