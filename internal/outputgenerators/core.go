package outputgenerators

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	"github.com/korylprince/ipnetgen"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

// The Core output generator generates a XML configuration file for CORE.
type Core struct{}

type coreData struct {
	ScenarioName string
	Devices      []device
	Networks     []network
	WorldSize    customtypes.Area
}

type device struct {
	Id            uint
	IdInNodeGroup uint
	Name          string
	Position      customtypes.Position
	Type          string

	IPv4 string
	IPv6 string
	Mac  string
}

type network struct {
	Id       uint
	Prefix   string
	Position customtypes.Position

	IPv4Mask int
	IPv6Mask int

	NetworkTypeName string
	NetworkType     networktypes.NetworkType

	Devices []device
}

var ScenarioIdCounter uint
var lastPosition customtypes.Position

const NODE_SIZE int = 100

// getPosition returns an incrementing position for nodes to be placed on the canvas without overlap
func (Core) getPosition(worldSize customtypes.Area) customtypes.Position {
	lastPosition.X = lastPosition.X + NODE_SIZE
	if lastPosition.X >= int(worldSize.Width) {
		lastPosition.X = NODE_SIZE
		lastPosition.Y = lastPosition.Y + NODE_SIZE
		if lastPosition.Y >= int(worldSize.Height) {
			lastPosition.Y = int(worldSize.Height) - NODE_SIZE
		}
	}
	return lastPosition
}

// getId returns an incrementing unique id used for all nodes in CORE configuration
func (Core) getId() uint {
	ScenarioIdCounter++
	return ScenarioIdCounter - 1
}

// getMac returns an unique MAC address for different NodeGroups and index
func (Core) getMac(groupIndex uint, nodeGroup experiment.NodeGroup, index uint) string {
	return fmt.Sprintf("02:%02x:%02x:00:00:%02x", groupIndex, int(nodeGroup.Prefix[0]), index)
}

// getIpSpace returns an ip space consisting of an IPv4 and an IPv6 network. They are collision free for different index values.
func (Core) getIpSpace(index uint) (IPv4Net *ipnetgen.IPNetGenerator, IPv4Mask int, IPv6Net *ipnetgen.IPNetGenerator, IPv6Mask int) {
	v4Net, err := ipnetgen.New(fmt.Sprintf("10.%d.0.0/24", index))
	if err != nil {
		panic(err)
	}
	v4Mask, _ := v4Net.Mask.Size()
	v4Net.Next()

	v6Net, err := ipnetgen.New(fmt.Sprintf("2001:%d::/64", index))
	if err != nil {
		panic(err)
	}
	v6Mask, _ := v6Net.Mask.Size()
	v6Net.Next()

	return v4Net, v4Mask, v6Net, v6Mask
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
	case experiment.NODE_TYPE_ROUTER:
		return "router"

	case experiment.NODE_TYPE_PC:
		return "PC"

	default:
		return ""
	}
}

// Generate generates the XML configuration for CORE for a given Experiment.
func (c Core) Generate(exp experiment.Experiment) {
	ScenarioIdCounter = 5
	lastPosition = customtypes.Position{X: 0, Y: NODE_SIZE}

	os.Mkdir(OutputFolder, 0755)
	fbuffer, err := os.Create(filepath.Join(OutputFolder, "core.xml"))
	if err != nil {
		panic(err)
	}

	networks := []network{}
	var i uint
	for i = 0; int(i) < len(exp.NodeGroups); i++ {
		nodeGroup := exp.NodeGroups[i]
		devices := []device{}

		IPv4Net, IPv4Mask, IPv6Net, IPv6Mask := c.getIpSpace(i + 1)

		var y uint
		for y = 0; y < nodeGroup.NoNodes; y++ {
			IPv4 := IPv4Net.Next()
			IPv6 := IPv6Net.Next()

			devType := nodeGroup.NodesType
			// If NodeGroup consists of PCs, make first device a Router so they have a gateway
			if nodeGroup.NodesType == experiment.NODE_TYPE_PC && y == 0 {
				devType = experiment.NODE_TYPE_ROUTER
			}

			dev := device{
				Id:            c.getId(),
				IdInNodeGroup: y,
				Name:          fmt.Sprintf("%s%d", nodeGroup.Prefix, y+1),
				Position:      c.getPosition(exp.WorldSize),

				Type: c.deviceType(devType),

				IPv4: IPv4.String(),
				IPv6: IPv6.String(),
				Mac:  c.getMac(i+1, nodeGroup, y+1),
			}

			devices = append(devices, dev)
		}
		networks = append(networks, network{
			Id:       c.getId(),
			Prefix:   nodeGroup.Prefix,
			Position: c.getPosition(exp.WorldSize),

			IPv4Mask: IPv4Mask,
			IPv6Mask: IPv6Mask,

			NetworkTypeName: c.networkType(nodeGroup.NetworkType),
			NetworkType:     nodeGroup.NetworkType,

			Devices: devices,
		})
		// Get a new position to create empty space between networks
		c.getPosition(exp.WorldSize)
	}
	replacements := coreData{
		ScenarioName: exp.Name,
		Networks:     networks,
		WorldSize:    exp.WorldSize,
	}
	xmlTemplate, err := template.ParseFiles(filepath.Join(GetTemplatesFolder(), "core.xml"))
	if err != nil {
		panic(err)
	}
	xmlTemplate.Execute(fbuffer, replacements)
}
