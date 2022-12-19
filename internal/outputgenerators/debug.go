package outputgenerators

import (
	"fmt"
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
	os.Mkdir(OUTPUT_FOLDER, 0755)
	os.WriteFile(fmt.Sprintf("%s/debug_out.toml", OUTPUT_FOLDER), b, 0644)
}
