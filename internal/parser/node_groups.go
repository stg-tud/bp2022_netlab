package parser

import (
	"fmt"
	"strings"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// Input format of a NodeGroup configuration
type inputNodeGroup struct {
	Prefix        any `required:"true"`
	NoNodes       any `default:"1"`
	MovementModel any `default:"Static"`
	NodesType     any `default:"PC"`
	Networks      []string

	MinSpeed any `default:"0"`
	MaxSpeed any `default:"1"`
	MaxPause any `default:"0"`
}

// Intermediate representation of a NodeGroup
type intermediateNodeGroup struct {
	Prefix        string
	NoNodes       uint
	MovementModel string
	NodesType     string
}

// Parses all given inputNodeGroups to a list of valid experiment.NodeGroup
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

// Parses an input string to a valid experiment.NodeType
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

// Generates a list of *experiment.Network for a given list of network names as strings
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
		if count == 0 {
			output = append(output, networkPointers[lowerNetworkName])
			availableNetworkNames[lowerNetworkName] = count + 1
		}
	}
	return output, nil
}

// Parses a inputNodeGroup with to a movementpatterns.MovementPattern which fits the model string
func parseMovementModel(input inputNodeGroup, model string) (movementpatterns.MovementPattern, error) {
	switch strings.ToLower(model) {
	case "randomwaypoint", "random waypoint", "random_waypoint", "random":
		return fillDefaults[inputNodeGroup, movementpatterns.RandomWaypoint](input)

	case "static", "none":
		return fillDefaults[inputNodeGroup, movementpatterns.Static](input)

	default:
		logger.Warnf("Unknown movement pattern \"%s\". Using static instead. Please check your config.", model)
		return movementpatterns.Static{}, nil
	}
}
