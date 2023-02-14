package experiment

import (
	"errors"
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
	NodeGroups []nodegroup
	Targets    []string
	Warmup     uint
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

type nodegroup struct {
	Prefix        string
	NoNodes       uint
	MovementModel movement
	NodesType     string
	Networks      []string
}
type movement struct {
	Model    string
	MinSpeed int
	MaxSpeed int
	MaxPause int
}

// parse toml file into experiment struct
func LoadFromFile(file string) (exp Experiment, returnError error) {
	logger.Info("Generate experiment")
	var conf expConf

	buf, e := os.ReadFile(file)
	if e != nil {
		logger.Error("could not find toml file")

		return exp, e
	}

	err := toml.Unmarshal(buf, &conf)
	if err != nil {
		logger.Error("Error parsing toml into struct")
		return exp, err
	}
	// actual experiment
	exp = Experiment{}
	var replaceTargets [2]Target
	for i := range conf.Targets {
		switch conf.Targets[i] {
		case "Core":
			replaceTargets[i] = 0
		case "The One":
			replaceTargets[i] = 1
		default:
			return exp, errors.New("wrong targets")
		}
	}
	exp.Duration = conf.Duration
	exp.Name = conf.Name
	exp.RandomSeed = conf.RandomSeed
	exp.WorldSize = conf.WorldSize
	exp.Warmup = conf.Warmup
	// network slices
	nets := conf.Networks
	for i := range nets {

		netType, e := setDefaultNet(nets[i].Type, nets[i])
		if e != nil {
			logger.Error("While generating Experiments, could not find networktype")
			return exp, e
		}

		net, err := NewNetwork(nets[i].Name, netType)
		if err != nil {
			logger.Error("Error creating new Networks")
		}
		exp.Networks = append(exp.Networks, net)

	}
	// nodegroups slices
	nodes := conf.NodeGroups
	for i := range nodes {
		node, err := NewNodeGroup(nodes[i].Prefix, nodes[i].NoNodes)
		if err != nil {
			logger.Error("Error creating new Nodegroups")
		}

		switch nodes[i].NodesType {

		case "PC":
			node.NodesType = NodeTypePC
		case "Router":
			node.NodesType = NodeTypeRouter
		case "":
			logger.Error("not found nodetype")
			return exp, errors.New("wrong nodetype")
		default:
			node.NodesType = NodeTypePC
		}

		exp.NodeGroups = append(exp.NodeGroups, node)
		nets := nodes[i].Networks

		//networks of nodegroups
		for k := range nets {

			indexOfNetwork, e := setNetworkInNodeGroup(conf, nets[k])
			if e != nil {
				logger.Error("Error setting up Networks for Nodegroup")
			}
			exp.NodeGroups[i].Networks = append(exp.NodeGroups[i].Networks, &exp.Networks[indexOfNetwork])

		}
		//movementmodel of nodegroups
		model := nodes[i].MovementModel.Model

		switch model {

		case "Static":
			exp.NodeGroups[i].MovementModel = movementpatterns.Static{}
		case "":
			exp.NodeGroups[i].MovementModel = movementpatterns.RandomWaypoint{}
		case "RandomWaypoint":
			exp.NodeGroups[i].MovementModel = movementpatterns.RandomWaypoint{
				MinSpeed: conf.NodeGroups[i].MovementModel.MinSpeed,
				MaxSpeed: conf.NodeGroups[i].MovementModel.MaxSpeed,
				MaxPause: conf.NodeGroups[i].MovementModel.MaxPause,
			}
		default:
			exp.NodeGroups[i].MovementModel = movementpatterns.RandomWaypoint{
				MinSpeed: conf.NodeGroups[i].MovementModel.MinSpeed,
				MaxSpeed: conf.NodeGroups[i].MovementModel.MaxSpeed,
				MaxPause: conf.NodeGroups[i].MovementModel.MaxPause,
			}
		}

	}
	logger.Trace("Finished generation")
	return exp, nil
}

// return a networktype with the given name and sets them to deafault/ custom
func setDefaultNet(name string, net network) (networkType networktypes.NetworkType, err error) {

	switch name {
	case "wireless_lan":
		wirelesslan := networktypes.WirelessLAN{}.Default()
		if net.Bandwidth != 54000000 {
			wirelesslan.Bandwidth = net.Bandwidth
		}
		if net.Range != 275 {
			wirelesslan.Range = net.Range
		}
		if net.Jitter != 0 {
			wirelesslan.Jitter = net.Jitter
		}
		if net.Delay != 5000 {
			wirelesslan.Delay = net.Delay
		}
		if net.Loss != 0.0 {
			wirelesslan.Loss = net.Loss
		}
		if net.Promiscuous {
			wirelesslan.Promiscuous = net.Promiscuous
		}
		return wirelesslan, nil
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
		return wireless, nil
	case "emane":
		return networktypes.Emane{}.Default(), nil
	case "hub":
		return networktypes.Hub{}.Default(), nil
	case "switch":
		return networktypes.Switch{}.Default(), nil
	default:
		logger.Error("While generating Experiments, could not find networktype")
		return networktypes.WirelessLAN{}.Default(), errors.New("while generating Experiments, could not find networktype")
	}
}

func setNetworkInNodeGroup(conf expConf, nameOfNetwork string) (i int, err error) {

	for i, network := range conf.Networks {
		if nameOfNetwork == network.Name {
			return i, nil
		}
	}
	return i, errors.New("could not find the network for nodegroup")
}
