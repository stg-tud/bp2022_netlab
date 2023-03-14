package parser

import (
	"fmt"
	"strings"

	logger "github.com/gookit/slog"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type inputNetwork struct {
	Name any `required:"true"`
	Type any `default:"WirelessLAN"`

	Movement       any `default:"true"`
	Bandwidth      any `default:"54000000"`
	Range          any `default:"400"`
	Jitter         any `default:"0"`
	Delay          any `default:"5000"`
	Loss           any `default:"0.0"`
	LossInitial    any `default:"0.0"`
	LossFactor     any `default:"1.0"`
	LossStartRange any `default:"300.0"`
	Promiscuous    any `default:"false"`
}

type intermediateNetwork struct {
	Name string
	Type string
}

func parseNetworks(input []inputNetwork) ([]experiment.Network, error) {
	output := []experiment.Network{}
	names := make(map[string]bool)

	for i, inNetwork := range input {
		intermediate, err := fillDefaults[inputNetwork, intermediateNetwork](inNetwork)
		logger.Tracef("Processing network %d: \"%s\"", i, intermediate.Name)
		if err != nil {
			return output, fmt.Errorf("error parsing network %d: %s", i, err)
		}

		_, exists := names[strings.ToLower(intermediate.Name)]
		if exists {
			return output, fmt.Errorf("a network with the name \"%s\" already exists", intermediate.Name)
		}
		names[strings.ToLower(intermediate.Name)] = true

		networkType, err := parseNetworkType(inNetwork, intermediate.Type)
		if err != nil {
			return output, fmt.Errorf("error parsing network %d: %s", i, err)
		}

		outputNetwork, err := experiment.NewNetwork(intermediate.Name, networkType)
		if err != nil {
			return output, fmt.Errorf("error parsing network %d: %s", i, err)
		}

		output = append(output, outputNetwork)
	}

	return output, nil
}

func parseNetworkType(input inputNetwork, networkType string) (networktypes.NetworkType, error) {
	var output networktypes.NetworkType
	var err error
	switch strings.ToLower(networkType) {
	case "emane":
		output, err = fillDefaults[inputNetwork, networktypes.Emane](input)

	case "hub":
		output, err = fillDefaults[inputNetwork, networktypes.Hub](input)

	case "switch":
		output, err = fillDefaults[inputNetwork, networktypes.Switch](input)

	case "wireless":
		output, err = fillDefaults[inputNetwork, networktypes.Wireless](input)

	case "wirelesslan", "wireless-lan", "wireless lan", "wireless_lan":
		output, err = fillDefaults[inputNetwork, networktypes.WirelessLAN](input)

	default:
		return output, fmt.Errorf("network type \"%s\" not found", networkType)
	}

	if err != nil {
		return output, err
	}
	return output, nil
}
