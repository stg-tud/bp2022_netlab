package outputgenerators

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

type Debug struct{}

func (t Debug) Generate(exp experiment.Experiment) string {
	b, err := toml.Marshal(exp)
	if err != nil {
		panic(err)
	}
	return "### DEBUG OUTPUT ###\n" + string(b)
}
