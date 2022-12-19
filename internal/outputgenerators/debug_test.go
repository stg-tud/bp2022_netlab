package outputgenerators_test

import (
	"os"
	"testing"

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
		Name:       "Debug Output Test",
		Runs:       5,
		NodeGroups: nodegroups,
	}

	og := outputgenerators.Debug{}
	og.Generate(exp)

	expected, err := os.ReadFile("testdata/debug_out.toml")
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile("debug_out.toml")
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		t.Fatal("Output does not match expected output!")
	}

	os.Remove("debug_out.toml")

	return
}
