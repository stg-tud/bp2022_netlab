package main

import (
	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/logging"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
)

func main() {
	logging.Init(true)
	logger.Info("Starting")

	exampleExperiment, _ := parser.LoadFromFile("internal/experiment/testdata/load_test.toml")

	logger.Info("Using random seed", exampleExperiment.RandomSeed)

	outputGenerators := []outputgenerators.OutputGenerator{
		outputgenerators.Core{},
		outputgenerators.Bonnmotion{},
		outputgenerators.Debug{},
		outputgenerators.TheOne{},
		outputgenerators.CoreEmulab{},
	}

	for _, outputGenerator := range outputGenerators {
		outputGenerator.Generate(exampleExperiment)
	}

	logger.Info("Finished")
}
