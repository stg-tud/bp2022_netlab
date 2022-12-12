package outputgenerators

import "github.com/stg-tud/bp2022_netlab/internal/experiment"

type Bonnmotion struct{}

func (t Bonnmotion) Generate(experiment.Experiment) string {
	return "Bonnmotion output"
}
