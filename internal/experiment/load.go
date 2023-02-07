package experiment

import (
	

	"os"

	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"

	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type ExpConf struct {
	Name       string
	Runs       uint
	Networks   []NT
	RandomSeed int64
	Duration   uint
	WorldSize  customtypes.Area
	NodeGroups []Nodes
	Targets    []Target
}
type NT struct {
	Name        string
	Bandwidth   int
	Range       int
	Jitter      int
	Delay       int
	Loss        float32
	Promiscuous bool
	Movement    bool

	LossInitial    float32
	LossFactor     float32
	LossStartRange float32
}

type Nodes struct {
	Prefix        string
	NoNodes       uint
	MovementModel Movement
	NodesType     NodeType
	Networks      []*NT
}
type Movement struct {
	Model    string
	MinSpeed int
	MaxSpeed int
	MaxPause int
}

// Loads the path string with toml file into experiment
func LoadFromFile(file string) Experiment {
	var conf ExpConf

	buf, e := os.ReadFile(file)
	if e != nil {
		panic(e)
	}

	err := toml.Unmarshal(buf, &conf)
	if err != nil {
		panic(err)
	}
	// actual experiment
	exp := Experiment{}
	exp.Targets = conf.Targets
	exp.Duration = conf.Duration
	exp.Name = conf.Name
	exp.RandomSeed = conf.RandomSeed
	exp.WorldSize = conf.WorldSize

	// network slices
	nets := conf.Networks
	for i := 0; i < len(nets); i++ {

		netType := setDefaultNet(nets[i].Name, nets[i])

		net, _ := NewNetwork(nets[i].Name, netType)

		exp.Networks = append(exp.Networks, net)

	}
	// nodegroups slices
	nodes := conf.NodeGroups
	for i := 0; i < len(nodes); i++ {

		node, _ := NewNodeGroup(nodes[i].Prefix, nodes[i].NoNodes)
		node.NodesType = nodes[i].NodesType
		exp.NodeGroups = append(exp.NodeGroups, node)
		nets := nodes[i].Networks

		//networks of nodegroups
		for k := 0; k < len(nets); k++ {
			netType := setDefaultNet(nets[k].Name, *nets[k])

			net, _ := NewNetwork(nets[k].Name, netType)

			exp.NodeGroups[i].Networks = append(exp.NodeGroups[i].Networks, &net)
		}

		//movementmodel of nodegroups
		model := nodes[i].MovementModel.Model

		if model == "Static" {
			exp.NodeGroups[i].MovementModel = movementpatterns.Static{}
		}

		if model == "custom" {
			exp.NodeGroups[i].MovementModel = movementpatterns.RandomWaypoint{
				MinSpeed: nodes[i].MovementModel.MinSpeed,
				MaxSpeed: nodes[i].MovementModel.MaxSpeed,
				MaxPause: nodes[i].MovementModel.MaxPause,
			}
		}

	}

	
	return exp

}

// return a networktype with the given name and sets them to deafault/ custom
func setDefaultNet(s string, net NT) (networkType networktypes.NetworkType) {

	switch {
	case s == "wireless_lan":
		wirelesslan := networktypes.WirelessLAN{}.Default()
		if net.Bandwidth != 0 {
			wirelesslan.Bandwidth = net.Bandwidth
		}
		if net.Range != 0 {
			wirelesslan.Range = net.Range
		}
		if net.Delay != 0 {
			wirelesslan.Bandwidth = net.Delay
		}
		if net.Loss != 0 {
			wirelesslan.Loss = net.Loss
		}
		if net.Jitter != 0 {
			wirelesslan.Jitter = net.Jitter
		}
		if net.Promiscuous == true {
			wirelesslan.Promiscuous = true
		}

		return wirelesslan
	case s == "wireless":
		wireless := networktypes.Wireless{}.Default()
		if net.Bandwidth != 0 {
			wireless.Bandwidth = net.Bandwidth
		}
		if net.Range != 0 {
			wireless.Range = net.Range
		}
		if net.Delay != 0 {
			wireless.Bandwidth = net.Delay
		}
		if net.LossInitial != 0 {
			wireless.LossInitial = net.Loss
		}
		if net.LossFactor != 0 {
			wireless.LossFactor = net.Loss
		}
		if net.LossStartRange != 0 {
			wireless.LossStartRange = net.Loss
		}
		if net.Jitter != 0 {
			wireless.Jitter = net.Jitter
		}
		if net.Movement == false {
			wireless.Movement = false
		}

		return wireless
	case s == "emane":
		return networktypes.Emane{}.Default()
	case s == "hub":
		return networktypes.Hub{}.Default()
	case s == "switch":
		return networktypes.Switch{}.Default()
	default:
		return networktypes.WirelessLAN{}.Default()
	}
}
