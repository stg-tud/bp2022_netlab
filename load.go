package load

import (
	
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

var exp experiment.Experiment 

//
func Loading() {

	
	doc, err1 := os.ReadFile("test/format.toml")
	if err1 != nil {
		panic(err1)
	}

	
	err := toml.Unmarshal(doc,&exp)

	if err != nil {
		panic(err)
	}

}

func GetExperiment() (experiment.Experiment){
	return exp
}
