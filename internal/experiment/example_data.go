package experiment

// GetExampleExperiment returns a Experiment loaded with example values.
func GetExampleExperiment() Experiment {
	var nodegroups []NodeGroup
	nodegroups = append(nodegroups, NewNodeGroup("n", 1))
	nodegroups = append(nodegroups, NewNodeGroup("p", 2))

	var ExampleExperiment = Experiment{
		Name:       "Example Experiment",
		Runs:       1,
		NodeGroups: nodegroups,
	}
	return ExampleExperiment
}
