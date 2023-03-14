package parser

import (
	"fmt"
	"strings"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

type inputNodeGroup struct {
	Prefix        any `required:"true"`
	NoNodes       any `default:"1"`
	MovementModel any `default:"Static"`
	NodesType     any `default:"PC"`
	Networks      []string

	MinSpeed              any `default:"1"`
	MaxSpeed              any `default:"5"`
	MinPause              any `default:"10"`
	MaxPause              any `default:"3600"`
	Range                 any `default:"100"`
	Clusters              any `default:"40"`
	Alpha                 any `default:"1.45"`
	MinFlight             any `default:"1"`
	MaxFlight             any `default:"14000"`
	Beta                  any `default:"1.5"`
	NumberOfWaypoints     any `default:"500"`
	LevyExponent          any `default:"1.0"`
	HurstParameter        any `default:"0.75"`
	DistanceWeight        any `default:"3.0"`
	ClusteringRange       any `default:"50.0"`
	ClusterRatio          any `default:"3"`
	WaypointRatio         any `default:"5"`
	Radius                any `default:"0.2"`
	CellDistanceWeight    any `default:"0.5"`
	NodeSpeedMultiplier   any `default:"1.5"`
	WaitingTimeExponent   any `default:"1.55"`
	WaitingTimeUpperBound any `default:"100.0"`
}

type intermediateNodeGroup struct {
	Prefix        string
	NoNodes       uint
	MovementModel string
	NodesType     string
}

func parseNodeGroups(input []inputNodeGroup, exp *experiment.Experiment) ([]experiment.NodeGroup, error) {
	output := []experiment.NodeGroup{}
	prefixes := make(map[string]bool)

	for i, inNodeGroup := range input {
		intermediate, err := fillDefaults[inputNodeGroup, intermediateNodeGroup](inNodeGroup)
		if err != nil {
			return output, fmt.Errorf("error parsing node group %d: %s", i, err)
		}

		_, exists := prefixes[intermediate.Prefix]
		if exists {
			return output, fmt.Errorf("a node group with the prefix \"%s\" already exists", intermediate.Prefix)
		}
		prefixes[intermediate.Prefix] = true

		outputNodeGroup, err := experiment.NewNodeGroup(intermediate.Prefix, intermediate.NoNodes)
		if err != nil {
			return output, fmt.Errorf("error parsing node group %d: %s", i, err)
		}

		outputNodeGroup.NodesType, err = parseNodesType(intermediate.NodesType)
		if err != nil {
			return output, fmt.Errorf("error parsing node group %d: %s", i, err)
		}

		outputNodeGroup.Networks, err = parseNodeGroupNetworks(inNodeGroup.Networks, exp)
		if err != nil {
			return output, fmt.Errorf("error parsing node group %d: %s", i, err)
		}

		outputNodeGroup.MovementModel, err = parseMovementModel(inNodeGroup, intermediate.MovementModel)
		if err != nil {
			return output, fmt.Errorf("error parsing node group %d: %s", i, err)
		}

		output = append(output, outputNodeGroup)
	}

	return output, nil
}

func parseNodesType(input string) (experiment.NodeType, error) {
	switch strings.ToLower(input) {
	case "pc", "computer", "device", "node":
		return experiment.NodeTypePC, nil

	case "router":
		return experiment.NodeTypeRouter, nil

	default:
		return -1, fmt.Errorf("node type \"%s\" not found", input)
	}
}

func parseNodeGroupNetworks(input []string, exp *experiment.Experiment) ([]*experiment.Network, error) {
	output := []*experiment.Network{}
	availableNetworkNames := make(map[string]int)
	networkPointers := make(map[string]*experiment.Network)
	for i, network := range exp.Networks {
		availableNetworkNames[strings.ToLower(network.Name)] = 0
		networkPointers[strings.ToLower(network.Name)] = &exp.Networks[i]
	}

	for _, networkName := range input {
		lowerNetworkName := strings.ToLower(networkName)
		count, exists := availableNetworkNames[lowerNetworkName]
		if !exists {
			return output, fmt.Errorf("network \"%s\" not found", networkName)
		}
		if count > 0 {
			continue
		}
		output = append(output, networkPointers[lowerNetworkName])
		availableNetworkNames[lowerNetworkName] = count + 1
	}
	return output, nil
}

func parseMovementModel(input inputNodeGroup, model string) (movementpatterns.MovementPattern, error) {
	switch strings.ToLower(model) {
	case "randomwaypoint", "random waypoint", "random_waypoint", "random":
		return fillDefaults[inputNodeGroup, movementpatterns.RandomWaypoint](input)

	case "static", "none":
		return fillDefaults[inputNodeGroup, movementpatterns.Static](input)

	case "slaw":
		return fillDefaults[inputNodeGroup, movementpatterns.SLAW](input)

	case "smooth":
		return fillDefaults[inputNodeGroup, movementpatterns.SMOOTH](input)

	case "swim":
		return fillDefaults[inputNodeGroup, movementpatterns.SWIM](input)

	default:
		return movementpatterns.Static{}, fmt.Errorf("movement pattern \"%s\" not found", model)
	}
}
