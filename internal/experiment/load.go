package experiment

import (
	"errors"
	"os"
	"strings"

	logger "github.com/gookit/slog"
	"github.com/pelletier/go-toml/v2"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"

	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type expConf struct {
	Name            string
	Runs            uint
	Networks        []network
	RandomSeed      int64
	Duration        uint
	WorldSize       customtypes.Area
	NodeGroups      []nodegroup
	Targets         []string
	Warmup          uint
	EventGenerators []eventgenerator
}
type eventgenerator struct {
	Name      string
	Prefix    string
	Intervall customtypes.Interval
	Size      customtypes.Interval
	Hosts     customtypes.Interval
	ToHosts   customtypes.Interval
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
	Model                 string
	MinSpeed              int
	MaxSpeed              int
	MinPause              int
	MaxPause              int
	Range                 int
	Clusters              int
	Alpha                 float32
	MinFlight             int
	MaxFlight             int
	Beta                  float32
	NumberOfWaypoints     int
	LevyExponent          float32
	HurstParameter        float32
	DistanceWeight        float32
	ClusteringRange       float32
	ClusterRatio          int
	WaypointRatio         int
	Radius                float32
	CellDistanceWeight    float32
	NodeSpeedMultiplier   float32
	WaitingTimeExponent   float32
	WaitingTimeUpperBound float32
}

// parse toml file into experiment struct
func LoadFromFile(file string) (exp Experiment, returnError error) {
	logger.Info("Generate experiment")
	var conf expConf
	// read file
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
	var replaceTargets []Target
	for _, targetString := range conf.Targets {
		switch strings.ToLower(targetString) {
		case "core", "coreemu", "core-emu":
			replaceTargets = append(replaceTargets, TargetCore)
		case "coreemulab", "coreemu-lab", "core-emulab", "core-emu-lab", "clab":
			replaceTargets = append(replaceTargets, TargetCore)
		case "the one", "theone", "one":
			replaceTargets = append(replaceTargets, TargetTheOne)
		default:
			return exp, errors.New("error getting targets, could not find target")
		}
	}
	exp.Targets = replaceTargets

	// experiment other field
	exp.Duration = conf.Duration
	exp.Name = conf.Name
	exp.Runs = conf.Runs
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
		}
		exp.NodeGroups = append(exp.NodeGroups, node)
		//networks of nodegroups
		nets := nodes[i].Networks
		for k := range nets {

			indexOfNetwork, e := setNetworkInNodeGroup(conf, nets[k])
			if e != nil {
				logger.Error("Error setting up Networks for Nodegroup")
			}
			exp.NodeGroups[i].Networks = append(exp.NodeGroups[i].Networks, &exp.Networks[indexOfNetwork])

		}
		//movementmodel of nodegroups
		model := strings.ToLower(nodes[i].MovementModel.Model)
		modelConfig := conf.NodeGroups[i].MovementModel

		switch model {
		case "randomwaypoint", "random waypoint", "random":
			movementModel := movementpatterns.RandomWaypoint{}.Default().(movementpatterns.RandomWaypoint)
			movementModel.MinSpeed = modelConfig.MinSpeed
			movementModel.MaxSpeed = modelConfig.MaxSpeed
			movementModel.MaxPause = modelConfig.MaxPause
			exp.NodeGroups[i].MovementModel = movementModel
		case "smooth":
			movementModel := movementpatterns.SMOOTH{}.Default().(movementpatterns.SMOOTH)
			movementModel.Range = modelConfig.Range
			movementModel.Clusters = modelConfig.Clusters
			movementModel.Alpha = modelConfig.Alpha
			movementModel.MinFlight = modelConfig.MinFlight
			movementModel.MaxFlight = modelConfig.MaxFlight
			movementModel.Beta = modelConfig.Beta
			movementModel.MinPause = modelConfig.MinPause
			movementModel.MaxPause = modelConfig.MaxPause
			exp.NodeGroups[i].MovementModel = movementModel
		case "slaw":
			movementModel := movementpatterns.SLAW{}.Default().(movementpatterns.SLAW)
			movementModel.NumberOfWaypoints = modelConfig.NumberOfWaypoints
			movementModel.MinPause = modelConfig.MinPause
			movementModel.MaxPause = modelConfig.MaxPause
			movementModel.LevyExponent = modelConfig.LevyExponent
			movementModel.HurstParameter = modelConfig.HurstParameter
			movementModel.DistanceWeight = modelConfig.DistanceWeight
			movementModel.ClusteringRange = modelConfig.ClusteringRange
			movementModel.ClusterRatio = modelConfig.ClusterRatio
			movementModel.WaypointRatio = modelConfig.WaypointRatio
			exp.NodeGroups[i].MovementModel = movementModel
		case "swim":
			movementModel := movementpatterns.SWIM{}.Default().(movementpatterns.SWIM)
			movementModel.Radius = modelConfig.Radius
			movementModel.CellDistanceWeight = modelConfig.CellDistanceWeight
			movementModel.NodeSpeedMultiplier = modelConfig.NodeSpeedMultiplier
			movementModel.WaitingTimeExponent = modelConfig.WaitingTimeExponent
			movementModel.WaitingTimeUpperBound = modelConfig.WaitingTimeUpperBound
			exp.NodeGroups[i].MovementModel = movementModel
		case "static", "none":
		default:
			exp.NodeGroups[i].MovementModel = movementpatterns.Static{}.Default()

		}

	}

	//eventgeneratortypes
	events := []EventGenerator{}
	for _, eve := range conf.EventGenerators {
		buffer, e := setDefaultEventGenerator(eve.Name, eve)
		if e != nil {
			logger.Error("Error setting up event generators")
		}
		newEvent, err := NewEventGenerator(eve.Name, buffer)
		if err != nil {
			logger.Error("Error setting up new event generators")
		}
		events = append(events, newEvent)
	}
	exp.EventGenerators = events
	//finished
	logger.Trace("Finished generation")
	return exp, nil
}
func setDefaultEventGenerator(Name string, eve eventgenerator) (eventgeneratortypes.EventGeneratorType, error) {
	switch Name {

	case "MessageEventGenerator":
		msg := eventgeneratortypes.MessageEventGenerator{}.Default()
		if eve.Intervall.From != 25 && eve.Intervall.From != 0 {
			msg.Interval.From = eve.Intervall.From
		}
		if eve.Intervall.To != 35 && eve.Intervall.To != 0 {
			msg.Interval.From = eve.Intervall.From
		}
		if eve.Size.From != 80 && eve.Size.From != 120 {
			msg.Size.From = eve.Size.From
		}
		if eve.Size.To != 80 && eve.Size.To != 120 {
			msg.Size.To = eve.Size.To
		}
		if eve.Hosts.From != 5 && eve.Hosts.From != 0 {
			msg.Hosts.From = eve.Hosts.From
		}
		if eve.Hosts.To != 15 && eve.Hosts.To != 0 {
			msg.Hosts.To = eve.Hosts.To
		}
		if eve.ToHosts.From != 16 && eve.ToHosts.From != 0 {
			msg.ToHosts.From = eve.ToHosts.From
		}
		if eve.ToHosts.To != 17 && eve.ToHosts.To != 0 {
			msg.ToHosts.To = eve.ToHosts.To
		}
		if eve.Prefix != "M" && eve.Prefix != "" {
			msg.Prefix = eve.Prefix
		}
		return msg, nil

	case "MessageBurstGenerator":
		burst := eventgeneratortypes.MessageBurstGenerator{}.Default()
		if eve.Intervall.From != 25 && eve.Intervall.From != 0 {
			burst.Interval.From = eve.Intervall.From
		}
		if eve.Intervall.To != 35 && eve.Intervall.To != 0 {
			burst.Interval.From = eve.Intervall.From
		}
		if eve.Size.From != 80 && eve.Size.From != 120 {
			burst.Size.From = eve.Size.From
		}
		if eve.Size.To != 80 && eve.Size.To != 120 {
			burst.Size.To = eve.Size.To
		}
		if eve.Hosts.From != 5 && eve.Hosts.From != 0 {
			burst.Hosts.From = eve.Hosts.From
		}
		if eve.Hosts.To != 15 && eve.Hosts.To != 0 {
			burst.Hosts.To = eve.Hosts.To
		}
		if eve.ToHosts.From != 16 && eve.ToHosts.From != 0 {
			burst.ToHosts.From = eve.ToHosts.From
		}
		if eve.ToHosts.To != 17 && eve.ToHosts.To != 0 {
			burst.ToHosts.To = eve.ToHosts.To
		}
		if eve.Prefix != "M" && eve.Prefix != "" {
			burst.Prefix = eve.Prefix
		}
		return burst, nil
	default:
		logger.Error("Error while generating eventgenerator, name not found")
		return eventgeneratortypes.MessageEventGenerator{}.Default(), errors.New("error while generating eventgenerator, name not found")
	}
}

// return a networktype with the given type and sets them to deafault/ custom
func setDefaultNet(netType string, net network) (networkType networktypes.NetworkType, err error) {

	switch netType {
	case "wireless_lan":
		wirelesslan := networktypes.WirelessLAN{}.Default()
		if net.Bandwidth != 54000000 && net.Bandwidth != 0 {
			wirelesslan.Bandwidth = net.Bandwidth
		}
		if net.Range != 275 && net.Range != 0 {
			wirelesslan.Range = net.Range
		}
		if net.Jitter != 0 {
			wirelesslan.Jitter = net.Jitter
		}
		if net.Delay != 5000 && net.Delay != 0 {
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
		if net.Bandwidth != 54000000 && net.Bandwidth != 0 {
			wireless.Bandwidth = net.Bandwidth
		}
		if net.Range != 400 && net.Range != 0 {
			wireless.Range = net.Range
		}
		if net.Delay != 5000 && net.Delay != 0 {
			wireless.Bandwidth = net.Delay
		}
		if net.LossInitial != 0.0 {
			wireless.LossInitial = net.Loss
		}
		if net.LossFactor != 1.0 && net.LossFactor != 0 {
			wireless.LossFactor = net.Loss
		}
		if net.LossStartRange != 300.0 && net.LossStartRange != 0 {
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

// set nodegroups network with the corresponding name
func setNetworkInNodeGroup(conf expConf, nameOfNetwork string) (i int, err error) {

	for i, network := range conf.Networks {
		if nameOfNetwork == network.Name {
			return i, nil
		}
	}
	return i, errors.New("could not find the network for nodegroup")
}
