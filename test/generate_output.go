package main

import (
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func main() {
	og := outputgenerators.Debug{}
	exampleExperiment := experiment.GetExampleExperiment()
	og.Generate(exampleExperiment)
}
