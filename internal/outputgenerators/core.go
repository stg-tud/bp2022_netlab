package outputgenerators

import (
	"fmt"
	"os"
	"text/template"

	"github.com/korylprince/ipnetgen"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// The Core output generator generates a XML configuration file for CORE.
type Core struct{}

type coreData struct {
	ScenarioName string
	Devices      []device
	Networks     []network
}

type device struct {
	Id   int
	Name string

	IPv4 string
	IPv6 string
	Mac  string
}

type network struct {
	Id     int
	Prefix string

	IPv4Mask int
	IPv6Mask int

	Type        string
	Bandwidth   int
	Range       int
	Jitter      int
	Delay       int
	Error       int
	Promiscuous int

	Devices []device
}

var ScenarioIdCounter int

// getId returns an incrementing unique id used for all nodes in CORE configuration
func (Core) getId() int {
	ScenarioIdCounter++
	return ScenarioIdCounter - 1
}

// getMac returns an unique MAC address for different NodeGroups and index
func (Core) getMac(groupIndex int, nodeGroup experiment.NodeGroup, index int) string {
	return fmt.Sprintf("%02x:%02x:00:00:%02x:00", groupIndex, int(nodeGroup.Prefix[0]), index)
}

// getIpSpace returns an ip space consisting of an IPv4 and an IPv6 network. They are collision free for different index values.
func (Core) getIpSpace(index int) (IPv4Net *ipnetgen.IPNetGenerator, IPv4Mask int, IPv6Net *ipnetgen.IPNetGenerator, IPv6Mask int) {
	v4Net, err := ipnetgen.New(fmt.Sprintf("10.%d.0.0/24", index))
	if err != nil {
		panic(err)
	}
	v4Mask, _ := v4Net.Mask.Size()
	v4Net.Next()

	v6Net, err := ipnetgen.New(fmt.Sprintf("2001:%d::/120", index))
	if err != nil {
		panic(err)
	}
	v6Mask, _ := v6Net.Mask.Size()
	v6Net.Next()

	return v4Net, v4Mask, v6Net, v6Mask
}

// networkType returns the correct CORE string for given NetworkType
func (Core) networkType(nt experiment.NetworkType) string {
	switch nt {
	case experiment.WIRELESS_LAN:
		return "WIRELESS_LAN"

	default:
		return ""
	}
}

// Generate generates the XML configuration for CORE for a given Experiment.
func (c Core) Generate(exp experiment.Experiment) {
	ScenarioIdCounter = 0

	fbuffer, err := os.Create("output/core.xml")
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
				Id:   c.getId(),
				Name: fmt.Sprintf("%s%d", nodeGroup.Prefix, y),

				IPv4: IPv4.String(),
				IPv6: IPv6.String(),
				Mac:  c.getMac(i+1, nodeGroup, y+1),
			}

			devices = append(devices, dev)
		}
		networks = append(networks, network{
			Id:     c.getId(),
			Prefix: nodeGroup.Prefix,

			IPv4Mask: IPv4Mask,
			IPv6Mask: IPv6Mask,

			Type:        c.networkType(nodeGroup.NetworkType),
			Bandwidth:   nodeGroup.Bandwidth,
			Range:       nodeGroup.Range,
			Jitter:      nodeGroup.Jitter,
			Delay:       nodeGroup.Delay,
			Error:       nodeGroup.Error,
			Promiscuous: nodeGroup.Promiscuous,

			Devices: devices,
		})
	}
	replacements := coreData{
		ScenarioName: exp.Name,
		Networks:     networks,
	}
	xmlTemplate, err := template.ParseFiles("internal/outputgenerators/templates/core.xml")
	if err != nil {
		panic(err)
	}
	xmlTemplate.Execute(fbuffer, replacements)
}
