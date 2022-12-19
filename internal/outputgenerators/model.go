package outputgenerators

import "github.com/stg-tud/bp2022_netlab/internal/experiment"

type OutputGenerator interface {
	Generate(experiment.Experiment)
}
