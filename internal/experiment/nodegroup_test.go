package experiment_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stretchr/testify/assert"
)

func TestNewNodeGroup(t *testing.T) {
	nodeGroup, err := experiment.NewNodeGroup("testing", 12)
	assert.NoError(t, err)
	assert.Equal(t, "testing", nodeGroup.Prefix)
	assert.EqualValues(t, 12, nodeGroup.NoNodes)

	nodeGroup, err = experiment.NewNodeGroup("", 41)
	assert.Error(t, err)
	assert.Equal(t, errors.New("prefix must consist of at least one character"), err)

	nodeGroup, err = experiment.NewNodeGroup("prefix", 0)
	assert.Error(t, err)
	assert.Equal(t, errors.New("NodeGroup must at least consist of one node"), err)
}
