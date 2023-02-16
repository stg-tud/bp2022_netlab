package outputgenerators

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/korylprince/ipnetgen"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// The Core output generator generates a XML configuration file for CORE.
type Core struct{}

func (Core) String() string {
	return "CORE"
}

// The name of the file where the XML config should be written to
const CoreOutputFile string = "core.xml"

type coreData struct {
	ScenarioName string
	Devices      []device
	Networks     []network
	WorldSize    customtypes.Area
}

type network struct {
	Id               uint
	Name             string
	Position         customtypes.Position
	Type             networktypes.NetworkType
	TypeName         string
	DevicesConnected uint

	IPv4Net  *ipnetgen.IPNetGenerator
	IPv4Mask int
	IPv6Net  *ipnetgen.IPNetGenerator
	IPv6Mask int
}

type networkInterface struct {
	Network     *network
	IdInNetwork uint
	IPv4        string
	IPv6        string
	Mac         string
}

type device struct {
	Id            uint
	IdInNodeGroup uint
	Name          string
	Position      customtypes.Position
	Type          string
	Interfaces    []networkInterface
}

// The first id to be assigned inside a scenario
const ScenarioIdOffset uint = 5

var scenarioIdCounter uint
var lastPosition customtypes.Position

const nodeSize int = 100

// getPosition returns an incrementing position for nodes to be placed on the canvas without overlap
func (Core) getPosition(worldSize customtypes.Area) customtypes.Position {
	logger.Trace("Getting new position")
	lastPosition.X = lastPosition.X + nodeSize
	if lastPosition.X >= int(worldSize.Width) {
		lastPosition.X = nodeSize
		lastPosition.Y = lastPosition.Y + nodeSize
		if lastPosition.Y >= int(worldSize.Height) {
			lastPosition.Y = int(worldSize.Height) - nodeSize
		}
	}
	return lastPosition
}

// getId returns an incrementing unique id used for all nodes in CORE configuration
func (Core) getId() uint {
	scenarioIdCounter++
	return scenarioIdCounter - 1
}

// getMac returns an unique MAC address for different NodeGroups and index
func (Core) getMac(network network, index uint) string {
	return fmt.Sprintf("02:%02x:%02x:00:00:%02x", network.Id, int(network.Name[0]), index)
}

// getIpSpace returns an ip space consisting of an IPv4 and an IPv6 network. They are collision free for different index values.
func (Core) getIpSpace(index uint) (IPv4Net *ipnetgen.IPNetGenerator, IPv4Mask int, IPv6Net *ipnetgen.IPNetGenerator, IPv6Mask int, error error) {
	logger.Trace("Getting new IP space for index", index)
	v4Net, err := ipnetgen.New(fmt.Sprintf("10.%d.0.0/24", index))
	if err != nil {
		logger.Error("Error generating IPv4 space", err)
		return &ipnetgen.IPNetGenerator{}, 0, &ipnetgen.IPNetGenerator{}, 0, err
	}
	v4Mask, _ := v4Net.Mask.Size()
	v4Net.Next()

	v6Net, err := ipnetgen.New(fmt.Sprintf("2001:%d::/64", index))
	if err != nil {
		logger.Error("Error generating IPv6 space", err)
		return &ipnetgen.IPNetGenerator{}, 0, &ipnetgen.IPNetGenerator{}, 0, err
	}
	v6Mask, _ := v6Net.Mask.Size()
	v6Net.Next()

	return v4Net, v4Mask, v6Net, v6Mask, nil
}

// networkType returns the correct CORE string for given NetworkType
func (Core) networkType(networkType networktypes.NetworkType) string {
	switch networkType.(type) {
	case networktypes.WirelessLAN:
		return "WIRELESS_LAN"

	case networktypes.Switch:
		return "SWITCH"

	case networktypes.Wireless:
		return "WIRELESS"

	case networktypes.Hub:
		return "HUB"

	case networktypes.Emane:
		return "EMANE"

	default:
		return ""
	}
}

// deviceType returns the correct CORE string for a given NodeType
func (Core) deviceType(deviceType experiment.NodeType) string {
	switch deviceType {
	case experiment.NodeTypeRouter:
		return "router"

	case experiment.NodeTypePC:
		return "PC"

	default:
		return ""
	}
}

// buildNetworks generates a list of networks used for CORE configuration as well as a mapping of experiment.Network to (CORE) networks
func (c Core) buildNetworks(exp experiment.Experiment) ([]network, map[*experiment.Network]*network, error) {
	logger.Trace("Building networks")
	networks := []network{}
	networkMapping := make(map[*experiment.Network]*network)
	for i := range exp.Networks {
		expNetwork := &exp.Networks[i]
		logger.Tracef("Building network \"%s\"", expNetwork.Name)
		ui := uint(i)
		IPv4Net, IPv4Mask, IPv6Net, IPv6Mask, _ := c.getIpSpace(ui + 1)
		net := network{
			Id:               c.getId(),
			Position:         c.getPosition(exp.WorldSize),
			Name:             expNetwork.Name,
			Type:             expNetwork.Type,
			TypeName:         c.networkType(expNetwork.Type),
			DevicesConnected: 0,
			IPv4Net:          IPv4Net,
			IPv4Mask:         IPv4Mask,
			IPv6Net:          IPv6Net,
			IPv6Mask:         IPv6Mask,
		}
		networkMapping[expNetwork] = &net
		networks = append(networks, net)
	}
	return networks, networkMapping, nil
}

// buildDevice generates the device configuration for CORE
func (c Core) buildDevice(deviceId uint, exp experiment.Experiment, nodeGroup experiment.NodeGroup, nodeGroupNetworks []*network) (device, error) {
	logger.Tracef("Building device \"%s%d\"", nodeGroup.Prefix, deviceId)
	deviceNetworkInterfaces := []networkInterface{}
	for _, network := range nodeGroupNetworks {
		deviceNetworkInterface := networkInterface{
			Network:     network,
			IdInNetwork: network.DevicesConnected,
			IPv4:        network.IPv4Net.Next().String(),
			IPv6:        network.IPv6Net.Next().String(),
			Mac:         c.getMac(*network, deviceId+1),
		}
		network.DevicesConnected++
		deviceNetworkInterfaces = append(deviceNetworkInterfaces, deviceNetworkInterface)
	}

	dev := device{
		Id:            c.getId(),
		IdInNodeGroup: deviceId,
		Name:          fmt.Sprintf("%s%d", nodeGroup.Prefix, deviceId+1),
		Position:      c.getPosition(exp.WorldSize),
		Type:          c.deviceType(nodeGroup.NodesType),
		Interfaces:    deviceNetworkInterfaces,
	}
	return dev, nil
}

// buildDevices generates the configuration for devices for CORE
func (c Core) buildDevices(exp experiment.Experiment, networkMapping map[*experiment.Network]*network) ([]device, error) {
	logger.Trace("Building devices")
	devices := []device{}
	for _, nodeGroup := range exp.NodeGroups {
		logger.Tracef("Processing NodeGroup \"%s\"", nodeGroup.Prefix)
		nodeGroupNetworks := []*network{}
		for _, nodeGroupNetwork := range nodeGroup.Networks {
			network := networkMapping[nodeGroupNetwork]
			nodeGroupNetworks = append(nodeGroupNetworks, network)
		}
		var deviceId uint
		for deviceId = 0; deviceId < nodeGroup.NoNodes; deviceId++ {
			dev, err := c.buildDevice(deviceId, exp, nodeGroup, nodeGroupNetworks)
			if err != nil {
				return []device{}, err
			}
			devices = append(devices, dev)
		}
	}
	return devices, nil
}

// Generate generates the XML configuration for CORE for a given Experiment.
func (c Core) Generate(exp experiment.Experiment) {
	logger.Info("Generating CORE output")
	scenarioIdCounter = ScenarioIdOffset
	lastPosition = customtypes.Position{X: 0, Y: nodeSize}

	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	outputFilePath := filepath.Join(outputFolder, "core.xml")
	if !folderstructure.MayCreatePath(outputFilePath) {
		logger.Error("Not allowed to write output file!")
		return
	}
	logger.Tracef("Opening file \"%s\"", outputFilePath)
	fbuffer, err := os.Create(outputFilePath)
	if err != nil {
		logger.Error("Error creating output file:", err)
	}
	defer func() {
		if cerr := fbuffer.Close(); cerr != nil {
			logger.Error("Error closing step file:", cerr)
			err = cerr
		}
	}()

	networks, networkMapping, err := c.buildNetworks(exp)
	if err != nil {
		logger.Error("Error building networks:", err)
	}
	devices, err := c.buildDevices(exp, networkMapping)
	if err != nil {
		logger.Error("Error building devices:", err)
	}

	replacements := coreData{
		ScenarioName: exp.Name,
		Networks:     networks,
		Devices:      devices,
		WorldSize:    exp.WorldSize,
	}
	xmlTemplate, err := template.ParseFS(TemplatesFS, filepath.Join(TemplatesFolder, CoreOutputFile))
	if err != nil {
		logger.Error("Error opening template file:", err)
	}
	err = xmlTemplate.Execute(fbuffer, replacements)
	if err != nil {
		logger.Error("Could not execute XML template:", err)
		return
	}
	logger.Trace("Finished generation")
}
