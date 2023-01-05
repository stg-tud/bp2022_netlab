package main

import (
	"fmt"
	"github.com/stg-tud/bp2022_netlab/internal/load"
)

func main() {

	/**
	og := outputgenerators.Debug{}
	exampleExperiment := experiment.GetExampleExperiment()
	og.Generate(exampleExperiment)
	**/
	
	load.Loading()
	
	fmt.Println(load.GetExperiment().Duration)

	
}
