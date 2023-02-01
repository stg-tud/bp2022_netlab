package outputgenerators

import (
	"os"
	"path/filepath"

	logger "github.com/gookit/slog"
	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
)

// Debug output generator dumps the experiment config as TOML for debug purposes.
type Debug struct{}

// The name of the file where the debug output should be dumped to
const DebugOutputFile = "debug_out.toml"

// Generate outputs the given Experiment as TOML to the file debug_out.toml
func (Debug) Generate(exp experiment.Experiment) {
	logger.Info("Generating debug output")
	b, err := toml.Marshal(exp)
	if err != nil {
		logger.Error("Could not marshal Experiment to TOML!", err)
		return
	}
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	outputFilePath := filepath.Join(outputFolder, DebugOutputFile)
	if !folderstructure.MayCreatePath(outputFilePath) {
		logger.Error("Not allowed to write output file!")
		return
	}
	logger.Tracef("Writing file \"%s\"", outputFilePath)
	os.WriteFile(outputFilePath, b, 0644)
	logger.Trace("Finished generation")
}
