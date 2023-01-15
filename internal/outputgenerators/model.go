// Package outputgenerators holds generators for all supported output formats.
package outputgenerators

import (
	"path/filepath"
	"runtime"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

const OUTPUT_FOLDER = "output"
const TEMPLATES_FOLDER = "templates"

type OutputGenerator interface {
	// Generate takes a Experiment and generates the output in the given format
	// and writes in the appropriate files or executes the correct function call.
	Generate(experiment.Experiment)
}

func GetTemplatesFolder() string {
	_, callerPath, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(callerPath), TEMPLATES_FOLDER)
}
