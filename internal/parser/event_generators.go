package parser

import (
	"fmt"
	"strings"

	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// Input format of an EventGenerator configuration
type inputEventGenerator struct {
	Name   any `required:"true"`
	Type   any `default:"Message Burst Generator"`
	Prefix any `default:""`

	Interval inputInterval
	Size     inputInterval
	Hosts    inputInterval
	ToHosts  inputInterval
}

// Input format of an Interval configuration
type inputInterval struct {
	From any `default:"-1"`
	To   any `default:"-1"`
}

// Intermediate representation of an EventGeneratorType
type intermediateEventGeneratorType struct {
	Prefix   string
	Interval customtypes.Interval
	Size     customtypes.Interval
	Hosts    customtypes.Interval
	ToHosts  customtypes.Interval
}

// Intermediate representation of a EventGenerator
type intermediateEventGenerator struct {
	Name string
	Type string
}

// Parses all given inputEventGenerator to a list of valid experiment.EventGenerator
func parseEventGenerators(input []inputEventGenerator) ([]experiment.EventGenerator, error) {
	output := []experiment.EventGenerator{}
	names := make(map[string]bool)

	for i, inEvenetGenerator := range input {
		intermediate, err := fillDefaults[inputEventGenerator, intermediateEventGenerator](inEvenetGenerator)
		if err != nil {
			return output, fmt.Errorf("error parsing event generator %d: %s", i, err)
		}

		_, exists := names[strings.ToLower(intermediate.Name)]
		if exists {
			return output, fmt.Errorf("an event generator with the name \"%s\" already exists", intermediate.Name)
		}
		names[strings.ToLower(intermediate.Name)] = true

		eventGeneratorType, err := parseEventGeneratorType(inEvenetGenerator, intermediate.Type, intermediate.Name)
		if err != nil {
			return output, fmt.Errorf("error parsing node group %d: %s", i, err)
		}

		outputEventGenerator, err := experiment.NewEventGenerator(intermediate.Name, eventGeneratorType)
		if err != nil {
			return output, fmt.Errorf("error parsing event generator %d: %s", i, err)
		}

		fmt.Printf("%#v\n", outputEventGenerator)

		output = append(output, outputEventGenerator)
	}

	return output, nil
}

// Parses a given inputEventGenerator with a given type and name as strings to a valid eventgeneratortypes.EventGeneratorType
func parseEventGeneratorType(input inputEventGenerator, eventGeneratorType string, name string) (eventgeneratortypes.EventGeneratorType, error) {
	var output eventgeneratortypes.EventGeneratorType
	prefixes := make(map[string]bool)

	intermediate, err := fillDefaults[inputEventGenerator, intermediateEventGeneratorType](input)
	if err != nil {
		return output, err
	}

	if intermediate.Prefix == "" {
		intermediate.Prefix = name
	}
	_, exists := prefixes[strings.ToLower(intermediate.Prefix)]
	if exists {
		return output, fmt.Errorf("an event generator with the prefix \"%s\" already exists", intermediate.Prefix)
	}
	prefixes[strings.ToLower(intermediate.Prefix)] = true

	intermediate.Interval, err = fillDefaults[inputInterval, customtypes.Interval](input.Interval)
	if err != nil {
		return output, err
	}

	intermediate.Size, err = fillDefaults[inputInterval, customtypes.Interval](input.Size)
	if err != nil {
		return output, err
	}

	intermediate.Hosts, err = fillDefaults[inputInterval, customtypes.Interval](input.Hosts)
	if err != nil {
		return output, err
	}

	intermediate.ToHosts, err = fillDefaults[inputInterval, customtypes.Interval](input.ToHosts)
	if err != nil {
		return output, err
	}

	switch strings.ToLower(eventGeneratorType) {
	case "messageburstgenerator", "message burst generator", "messageburst", "message burst", "burst":
		output, err := fillDefaults[intermediateEventGeneratorType, eventgeneratortypes.MessageBurstGenerator](intermediate)
		if err != nil {
			return output, err
		}
		defaults := eventgeneratortypes.MessageBurstGenerator{}.Default()
		output.Interval = compareIntervals(output.Interval, defaults.Interval)
		output.Size = compareIntervals(output.Size, defaults.Size)
		output.Hosts = compareIntervals(output.Hosts, defaults.Hosts)
		output.ToHosts = compareIntervals(output.ToHosts, defaults.ToHosts)
		return output, nil

	case "messageeventgenerator", "message event generator", "messageevent", "message event", "event":
		output, err := fillDefaults[intermediateEventGeneratorType, eventgeneratortypes.MessageEventGenerator](intermediate)
		if err != nil {
			return output, err
		}
		defaults := eventgeneratortypes.MessageEventGenerator{}.Default()
		output.Interval = compareIntervals(output.Interval, defaults.Interval)
		output.Size = compareIntervals(output.Size, defaults.Size)
		output.Hosts = compareIntervals(output.Hosts, defaults.Hosts)
		output.ToHosts = compareIntervals(output.ToHosts, defaults.ToHosts)
		return output, nil

	default:
		return output, fmt.Errorf("event generator type \"%s\" not found", eventGeneratorType)
	}
}

// Compares an input Interval with an defaults Interval. Applies default values wherever the input has -1 as value.
func compareIntervals(input customtypes.Interval, defaults customtypes.Interval) customtypes.Interval {
	if input.From != -1 {
		defaults.From = input.From
	}
	if input.To != -1 {
		defaults.To = input.To
	}
	return defaults
}
