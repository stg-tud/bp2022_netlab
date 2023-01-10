package experiment

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

// loads the path string with toml file into experiment and returns experiment

func Loading(path string) Experiment {
	var exp Experiment
	doc, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}

	err := toml.Unmarshal(doc, &exp)

	if err != nil {
		panic(err)
	}
	var nodegroups []NodeGroup
	for i, _ := range exp.NodeGroups {
		nodegroups = append(nodegroups, NewNodeGroup(exp.NodeGroups[i].Prefix, exp.NodeGroups[i].NoNodes))
	}
	exp.NodeGroups = nodegroups
	return exp
}
