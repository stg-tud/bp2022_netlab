// Package outputgenerators holds generators for all supported output formats.
package outputgenerators

import "github.com/stg-tud/bp2022_netlab/internal/experiment"

const OUTPUT_FOLDER = "output"

type OutputGenerator interface {
	// Generate takes a Experiment and generates the output in the given format
	// and writes in the appropriate files or executes the correct function call.
	Generate(experiment.Experiment)
}
