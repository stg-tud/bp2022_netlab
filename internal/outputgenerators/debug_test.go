package outputgenerators_test

import (
	"os"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

const EXPECTED_OUTPUT = "Test"

func TestDebug(t *testing.T) {
	var nodegroups []experiment.NodeGroup
	nodegroups = append(nodegroups, experiment.NewNodeGroup("a", 1))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("b", 2))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("c", 3))
	nodegroups = append(nodegroups, experiment.NewNodeGroup("d", 4))

	exp := experiment.Experiment{
		Name:    "Debug Output Test",
		Runs:    5,
		Targets: []experiment.Target{experiment.TARGET_CORE, experiment.TARGET_THEONE},

		Duration: 123456,
		WorldSize: customtypes.Area{
			Height: 170,
			Width:  240,
		},

		NodeGroups: nodegroups,
	}

	t.Cleanup(func() {
		os.RemoveAll(outputgenerators.OUTPUT_FOLDER)
	})

	og := outputgenerators.Debug{}
	og.Generate(exp)

}
