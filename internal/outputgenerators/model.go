// Package outputgenerators holds generators for all supported output formats.
package outputgenerators

import (
	"path/filepath"
	"runtime"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// The name of the folder all output files from generators should be put in.
const OutputFolder = "output"

// The name of the folder containing the template files for output generators.
const TemplatesFolder = "templates"

type OutputGenerator interface {
	// Generate takes a Experiment and generates the output in the given format
	// and writes in the appropriate files or executes the correct function call.
	Generate(experiment.Experiment)
}

// GetTemplatesFolder returns the absolute path of the folder containing the template files.
func GetTemplatesFolder() string {
	_, callerPath, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(callerPath), TemplatesFolder)
}
