package experiment

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

//loads the path string with toml file into experiment

func Loading(path string) Experiment {
	var exp Experiment

	doc, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}

	err := toml.Unmarshal(doc, &exp)
	if err != nil {
		panic(err)
	}
	return exp
}
