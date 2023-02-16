// Package outputgenerators holds generators for all supported output formats.
package outputgenerators

import (
	"embed"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// The name of the folder containing the template files for output generators.
const TemplatesFolder = "templates"

// Embed template files into binary.
//
//go:embed templates
var TemplatesFS embed.FS

type OutputGenerator interface {
	// Generate takes a Experiment and generates the output in the given format
	// and writes in the appropriate files or executes the correct function call.
	Generate(experiment.Experiment)
	String() string
}
