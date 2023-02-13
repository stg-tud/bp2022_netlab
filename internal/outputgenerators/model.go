// Package outputgenerators holds generators for all supported output formats.
package outputgenerators

import (
	"path/filepath"
	"runtime"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// The name of the folder containing the template files for output generators.
const TemplatesFolder = "templates"

type OutputGenerator interface {
	// Generate takes a Experiment and generates the output in the given format
	// and writes in the appropriate files or executes the correct function call.
	Generate(experiment.Experiment)
	String() string
}

// GetTemplatesFolder returns the absolute path of the folder containing the template files.
func GetTemplatesFolder() string {
	_, callerPath, _, _ := runtime.Caller(0)
	logger.Trace("Getting template folder from caller path:", callerPath)
	return filepath.Join(filepath.Dir(callerPath), TemplatesFolder)
}
