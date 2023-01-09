package outputgenerators

import (
	"fmt"
	"net"
	"os"
	"text/template"

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

func getId() int {
	ScenarioIdCounter++
	return ScenarioIdCounter - 1
}

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

		IPv4 := net.ParseIP(nodeGroup.IPv4Net).To4()

		for y := 0; y < nodeGroup.NoNodes; y++ {
			dev := device{
				Id:   getId(),
				Name: fmt.Sprintf("%s%d", nodeGroup.Prefix, y),

				IPv4: net.IPv4(IPv4[0], IPv4[1], IPv4[2], IPv4[3]+byte(y+1)).String(),
				IPv6: nodeGroup.IPv6Net,
				Mac:  fmt.Sprintf("%02x:%02x:00:00:%02x:00", i, int(nodeGroup.Prefix[0]), y+1),
			}

			devices = append(devices, dev)
		}
		networks = append(networks, network{
			Id:     getId(),
			Prefix: nodeGroup.Prefix,

			IPv4Mask: nodeGroup.IPv4Mask,
			IPv6Mask: nodeGroup.IPv6Mask,

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
