// Package outputgenerators holds generators for all supported output formats.
package outputgenerators

import "github.com/stg-tud/bp2022_netlab/internal/experiment"

// The name of the folder all output files from generators should be put in.
const OutputFolder = "output"

type OutputGenerator interface {
	// Generate takes a Experiment and generates the output in the given format
	// and writes in the appropriate files or executes the correct function call.
	Generate(experiment.Experiment)
}
