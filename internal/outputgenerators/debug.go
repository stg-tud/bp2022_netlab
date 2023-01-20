package outputgenerators

import (
	"os"
	"path"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// Debug output generator dumps the experiment config as TOML for debug purposes.
type Debug struct{}

// The name of the file where the debug output should be dumped to
const DebugOutputFile = "debug_out.toml"

// Generate outputs the given Experiment as TOML to the file debug_out.toml
func (Debug) Generate(exp experiment.Experiment) {
	b, err := toml.Marshal(exp)
	if err != nil {
		panic(err)
	}
	os.Mkdir(OutputFolder, 0755)
	os.WriteFile(path.Join(OutputFolder, DebugOutputFile), b, 0644)
}
