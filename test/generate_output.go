package main

import (
	"fmt"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func main() {
	og := outputgenerators.Debug{}
	fmt.Println(og.Generate(experiment.ExampleExperiment))
}
