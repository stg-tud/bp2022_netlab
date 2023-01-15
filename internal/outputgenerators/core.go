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
	Id       int
	Name     string
	Position customtypes.Position
	Type     string

	IPv4 string
	IPv6 string
	Mac  string
}

type network struct {
	Id       int
	Prefix   string
	Position customtypes.Position

	IPv4Mask int
	IPv6Mask int

	NetworkTypeName string
	NetworkType     networktypes.NetworkType

	Devices []device
}

var ScenarioIdCounter int
var lastPosition customtypes.Position

const NODE_SIZE int = 100

// getPosition returns an incrementing position for nodes to be placed on the canvas without overlap
func (Core) getPosition(worldSize customtypes.Area) customtypes.Position {
	lastPosition.X = lastPosition.X + NODE_SIZE
	if lastPosition.X >= worldSize.Width {
		lastPosition.X = NODE_SIZE
		lastPosition.Y = lastPosition.Y + NODE_SIZE
		if lastPosition.Y >= worldSize.Height {
			lastPosition.Y = worldSize.Height - NODE_SIZE
		}
	}
	return lastPosition
}

// getId returns an incrementing unique id used for all nodes in CORE configuration
func (Core) getId() int {
	ScenarioIdCounter++
	return ScenarioIdCounter - 1
}

// getMac returns an unique MAC address for different NodeGroups and index
func (Core) getMac(groupIndex int, nodeGroup experiment.NodeGroup, index int) string {
	return fmt.Sprintf("02:%02x:%02x:00:00:%02x", groupIndex, int(nodeGroup.Prefix[0]), index)
}

// getIpSpace returns an ip space consisting of an IPv4 and an IPv6 network. They are collision free for different index values.
func (Core) getIpSpace(index int) (IPv4Net *ipnetgen.IPNetGenerator, IPv4Mask int, IPv6Net *ipnetgen.IPNetGenerator, IPv6Mask int) {
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
func (Core) networkType(nt networktypes.NetworkType) string {
	switch nt.(type) {
	case networktypes.WirelessLAN:
		return "WIRELESS_LAN"

	case networktypes.Switch:
		return "SWITCH"

	default:
		return ""
	}
}

// Generate generates the XML configuration for CORE for a given Experiment.
func (c Core) Generate(exp experiment.Experiment) {
	ScenarioIdCounter = 5
	lastPosition = customtypes.Position{X: 0, Y: NODE_SIZE}

	os.Mkdir(OUTPUT_FOLDER, 0755)
	fbuffer, err := os.Create(filepath.Join(OUTPUT_FOLDER, "core.xml"))
	if err != nil {
		panic(err)
	}

	networks := []network{}
	for i := 0; i < len(exp.NodeGroups); i++ {
		nodeGroup := exp.NodeGroups[i]
		devices := []device{}

		IPv4Net, IPv4Mask, IPv6Net, IPv6Mask := c.getIpSpace(i + 1)

		for y := 0; y < nodeGroup.NoNodes; y++ {
			IPv4 := IPv4Net.Next()
			IPv6 := IPv6Net.Next()

			dev := device{
				Id:       c.getId(),
				Name:     fmt.Sprintf("%s%d", nodeGroup.Prefix, y+1),
				Position: c.getPosition(exp.WorldSize),

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
