// Package experiment holds the data structure needed to represent a experiment.
package experiment

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

// A Experiment is a fixed set of parameters to run a simulation.
type Experiment struct {
	Name    string
	Runs    uint
	Targets []Target

	Duration  uint
	WorldSize customtypes.Area

	NodeGroups []NodeGroup
}
