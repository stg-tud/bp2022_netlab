package experiment

import (
	"os"

	logger "github.com/gookit/slog"
	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"

	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type expConf struct {
	Name       string
	Runs       uint
	Networks   []network
	RandomSeed int64
	Duration   uint
	WorldSize  customtypes.Area
	NodeGroups []node
	Targets    []Target
}
type network struct {
	Name        string
	Type        string
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

type node struct {
	Prefix        string
	NoNodes       uint
	MovementModel Movement
	NodesType     NodeType
	Networks      []uint
}
type Movement struct {
	Model    string
	MinSpeed int
	MaxSpeed int
	MaxPause int
}

// parse toml file into experiment struct
func LoadFromFile(file string) Experiment {
	logger.Info("Generate experiment")
	var conf expConf

	buf, e := os.ReadFile(file)
	if e != nil {
		logger.Error("could not find toml file")
	}

	err := toml.Unmarshal(buf, &conf)
	if err != nil {
		logger.Error("Error parsing toml into struct")
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

		netType := setDefaultNet(nets[i].Type, nets[i])

		net, err := NewNetwork(nets[i].Name, netType)
		if err != nil {
			logger.Error("Error creating new Networks")
		}
		exp.Networks = append(exp.Networks, net)

	}
	// nodegroups slices
	nodes := conf.NodeGroups
	for i := 0; i < len(nodes); i++ {
		node, err := NewNodeGroup(nodes[i].Prefix, nodes[i].NoNodes)
		if err != nil {
			logger.Error("Error creating new Nodegroups")
		}
		node.NodesType = nodes[i].NodesType
		exp.NodeGroups = append(exp.NodeGroups, node)
		nets := nodes[i].Networks

		//networks of nodegroups
		for k := 0; k < len(nets); k++ {
			exp.NodeGroups[i].Networks = append(exp.NodeGroups[i].Networks, &exp.Networks[nets[k]])
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
	logger.Trace("Finished generation")
	return exp

}

// return a networktype with the given name and sets them to deafault/ custom
func setDefaultNet(name string, net network) (networkType networktypes.NetworkType) {

	switch name {
	case "wireless_lan":
		wirelesslan := networktypes.WirelessLAN{}.Default()
		if net.Bandwidth!=54000000 {
			wirelesslan.Bandwidth=net.Bandwidth
		}
		if net.Range !=275{
			wirelesslan.Range=net.Range
		}
		if net.Jitter != 0{
			wirelesslan.Jitter=net.Jitter
		}
		if net.Delay != 5000{
			wirelesslan.Delay=net.Delay
		}
		if net.Loss != 0.0{
			wirelesslan.Loss=net.Loss
		}
		if net.Promiscuous {
			wirelesslan.Promiscuous=net.Promiscuous
		}
		return wirelesslan
	case "wireless":
		wireless := networktypes.Wireless{}.Default()
		if net.Bandwidth != 54000000 {
			wireless.Bandwidth = net.Bandwidth
		}
		if net.Range != 400 {
			wireless.Range = net.Range
		}
		if net.Delay != 5000 {
			wireless.Bandwidth = net.Delay
		}
		if net.LossInitial != 0.0 {
			wireless.LossInitial = net.Loss
		}
		if net.LossFactor != 1.0 {
			wireless.LossFactor = net.Loss
		}
		if net.LossStartRange != 300.0 {
			wireless.LossStartRange = net.Loss
		}
		if net.Jitter != 0 {
			wireless.Jitter = net.Jitter
		}
		if !net.Movement {
			wireless.Movement = net.Promiscuous
		}

		return wireless
	case "emane":
		return networktypes.Emane{}.Default()
	case "hub":
		return networktypes.Hub{}.Default()
	case "switch":
		return networktypes.Switch{}.Default()
	default:
		logger.Error("While generating Experiments, could not find networktype")
		return 
	}
}
