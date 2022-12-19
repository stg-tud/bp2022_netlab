package outputgenerators

import "github.com/stg-tud/bp2022_netlab/internal/experiment"

const OUTPUT_FOLDER = "output"

type OutputGenerator interface {
	Generate(experiment.Experiment)
}
