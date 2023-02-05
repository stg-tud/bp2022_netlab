package experiment

import (
	"fmt"
	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type ExpConf struct {
	Name string
	Runs uint
	Networks []NT
	RandomSeed int64
	Duration uint
	WorldSize customtypes.Area
	NodeGroups []Nodes
	Targets []Target

}
type NT struct{
	Name string
	Bandwidth int
	Range int
	Jitter int
	Delay int
	Loss float32
	Promiscuous bool
	Movement bool
	
	LossInitial float32
	LossFactor float32
	LossStartRange float32
}

type Nodes struct{
	Prefix string
	NoNodes uint
	MovementModel Movement
	NodesType NodeType
	Networks []NT
}
type Movement struct{
	Model string
	MinSpeed int
	MaxSpeed int
	MaxPause int
}
// Loads the path string with toml file into experiment
func LoadFromFile(file string) Experiment {
	var exp ExpConf
	experiment:= Experiment{
		Networks: []Network{
			
		},
	}
	buf, e := os.ReadFile(file)
	if e != nil {
		panic(e)
	}
	
	err := toml.Unmarshal(buf, &exp)
	if err != nil {
		fmt.Println(err)
	}

	
	typ:=networktypes.Emane{

	}
	experiment.Networks[0].Type=typ
	fmt.Println(experiment.Networks[1].Name)

	return experiment

}
