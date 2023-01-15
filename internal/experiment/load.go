package experiment

import (
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// Loads the path string with toml file into experiment
func LoadFromFile(path string) Experiment {
	var exp Experiment

	doc, e := os.ReadFile(path)
	if e != nil {
		panic(e)
	}

	err := toml.Unmarshal(doc, &exp)
	if err != nil {
		panic(err)
	}

	for i := 0; i < len(exp.NodeGroups); i++ {

		if exp.NodeGroups[i].IPv4Net == "" {
			exp.NodeGroups[i].IPv4Net = defaultValues.IPv4Net
		}
		if exp.NodeGroups[i].IPv4Mask == 0 {
			exp.NodeGroups[i].IPv4Mask = defaultValues.IPv4Mask
		}
		if exp.NodeGroups[i].IPv6Net == "" {
			exp.NodeGroups[i].IPv6Net = defaultValues.IPv6Net
		}
		if exp.NodeGroups[i].IPv6Mask == 0 {
			exp.NodeGroups[i].IPv6Mask = defaultValues.IPv6Mask
		}
		if exp.NodeGroups[i].NetworkType == "" {
			exp.NodeGroups[i].NetworkType = defaultValues.NetworkType
		}
		if exp.NodeGroups[i].Range == 0 {
			exp.NodeGroups[i].Range = defaultValues.Range
		}
		if exp.NodeGroups[i].Bandwidth == 0 {
			exp.NodeGroups[i].Bandwidth = defaultValues.Bandwidth
		}
		if exp.NodeGroups[i].Jitter == 0 {
			exp.NodeGroups[i].Jitter = defaultValues.Jitter
		}
		if exp.NodeGroups[i].Delay == 0 {
			exp.NodeGroups[i].Delay = defaultValues.Delay
		}
		if exp.NodeGroups[i].Error == 0 {
			exp.NodeGroups[i].Error = defaultValues.Error
		}
		if exp.NodeGroups[i].Promiscuous == 0 {
			exp.NodeGroups[i].Promiscuous = defaultValues.Promiscuous
		}
		var nilMovement = movementpatterns.RandomWaypoint{}
		if exp.NodeGroups[i].MovementModel == nilMovement {
			exp.NodeGroups[i].MovementModel = defaultValues.MovementModel
		}
	}

	return exp
}
