package experiment

import (
	"github.com/pelletier/go-toml/v2"
)

var exp Experiment

// loads the read file into experiment and returns said experiment

func Loading(doc []byte) Experiment {

	err := toml.Unmarshal(doc, &exp)

	if err != nil {
		panic(err)
	}

	return exp
}
