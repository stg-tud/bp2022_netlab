// Package experiment holds the data structure needed to represent a experiment.
package experiment

// A Experiment is a fixed set of parameters to run a simulation.
type Experiment struct {
	Name       string
	Runs       int
	NodeGroups []NodeGroup
}
