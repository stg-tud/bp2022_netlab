package outputgenerators_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestOutput(t *testing.T) {

	core := outputgenerators.CoreEmulab{}

	exp := experiment.Experiment{
		Name:     "e1",
		Duration: 245,
	}
	core.Generate(exp)
}
