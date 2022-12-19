package outputgenerators

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

type Debug struct{}

func (t Debug) Generate(exp experiment.Experiment) {
	b, err := toml.Marshal(exp)
	if err != nil {
		panic(err)
	}
	os.WriteFile("debug_out.toml", b, 0644)
}
