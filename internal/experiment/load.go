package experiment

import (
	"os"

	"github.com/pelletier/go-toml/v2"
)

//loads the path string with toml file into experiment

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
	for i := 0; i < len(exp.NodeGroups); i++ {
		prefix := exp.NodeGroups[i].Prefix
		noNodes := exp.NodeGroups[i].NoNodes
		nodegroups = append(nodegroups, NewNodeGroup(prefix, noNodes))
	}
	exp.NodeGroups = nodegroups
	return exp
}
