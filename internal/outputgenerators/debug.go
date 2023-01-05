package outputgenerators

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// Debug output generator dumps the experiment config as TOML for debug purposes.
type Debug struct{}

// Generate outputs the given Experiment as TOML to the file debug_out.toml
func (t Debug) Generate(exp experiment.Experiment) {
	b, err := toml.Marshal(exp)
	if err != nil {
		panic(err)
	}
	os.Mkdir(OUTPUT_FOLDER, 0755)
	os.WriteFile(fmt.Sprintf("%s/debug_out.toml", OUTPUT_FOLDER), b, 0644)
}
