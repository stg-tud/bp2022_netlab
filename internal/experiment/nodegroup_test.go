package experiment_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestNewNodeGroup(t *testing.T) {
	nodeGroup, err := experiment.NewNodeGroup("testing", 12)
	if err != nil {
		t.Fatal("Error trying to generate a new NodeGroup", err)
	}
	if nodeGroup.Prefix != "testing" || nodeGroup.NoNodes != 12 {
		t.Fatal("NewNodeGroup() did not store the given parameters correctly!")
	}

	nodeGroup, err = experiment.NewNodeGroup("", 41)
	if err == nil {
		t.Fatal("NodeGroups without a prefix should not be allowed!")
	}

	nodeGroup, err = experiment.NewNodeGroup("prefix", 0)
	if err == nil {
		t.Fatal("NodeGroups without nodes should not be allowed!")
	}
}
