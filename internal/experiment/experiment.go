package experiment

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

type Experiment struct {
	Name    string
	Runs    int
	Targets []Target

	Duration  int
	WorldSize customtypes.Area

	NodeGroups []NodeGroup
}
