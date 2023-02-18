package outputgenerators

import (
	"fmt"
	"os"
	"path/filepath"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type TheOne struct{}

type group struct {
	Id             string
	NrofHosts      uint
	NrofInterfaces int
	Interfaces     []*experiment.Network
	MovementModel  string
}

type theOneData struct {
	ScenarioName                string
	ScenarioSimulateConnections bool
	ScenarioUpdateInterval      float64
	ScenarioEndTime             string
	WorldSizeHeight             uint
	WorldSizeWidth              uint
	NrofHostGroups              int
	Groups                      []group
	Interfaces                  []networkInterfaceTheOne
	EventGenerators             []eventGenerator
	RandomSeed                  int64
	Warmup                      uint
	Runtime                     uint
	NoEventGenerator            int
}
type eventGenerator struct {
	Name string
	Type eventgeneratortypes.EventGeneratorType
}
type networkInterfaceTheOne struct {
	Name      string
	Bandwidth int
	Range     int
}

// The name of the file that should be written to
const TheOneOutput = "cluster_settings.txt"

func (TheOne) String() string {
	return "The ONE"
}

// add function for template
func add(x, y int) int {
	return x + y
}

// returns the names of the movement pattern types in the way needed
func (t TheOne) movementPattern(movementPatternType movementpatterns.MovementPattern) string {
	switch movementPatternType.(type) {
	case movementpatterns.RandomWaypoint:
		return "RandomWaypoint"
	case movementpatterns.Static:
		return "Static"
	default:
		return ""
	}
}

// generates the group for cluster_settings.txt
func (t TheOne) buildGroups(exp experiment.Experiment) []group {
	logger.Trace("Building Groups")
	groups := []group{}

	for _, expNodeGroups := range exp.NodeGroups {
		group := group{
			Id:             expNodeGroups.Prefix,
			NrofHosts:      expNodeGroups.NoNodes,
			Interfaces:     expNodeGroups.Networks,
			MovementModel:  t.movementPattern(expNodeGroups.MovementModel),
			NrofInterfaces: len(expNodeGroups.Networks),
		}
		groups = append(groups, group)

	}
	return groups
}

// generates the interfaces for cluster_settings.txt
// 6750 bandwidth and 100 range are default values for wireless_lan, the prefered type if nothing else is set
func (t TheOne) buildInterfaces(exp experiment.Experiment) (networks []networkInterfaceTheOne) {
	logger.Trace("Building Interfaces")
	for _, expNetwork := range exp.Networks {
		bandwidth := 6570
		rangeOfNetwork := 100
		switch networkType := expNetwork.Type.(type) {
		case networktypes.WirelessLAN:
			bandwidth = networkType.Bandwidth
			rangeOfNetwork = networkType.Range

		case networktypes.Wireless:
			bandwidth = networkType.Bandwidth
			rangeOfNetwork = networkType.Range
		}

		nt := networkInterfaceTheOne{
			Name:      expNetwork.Name,
			Bandwidth: bandwidth,
			Range:     rangeOfNetwork,
		}
		networks = append(networks, nt)
	}

	return networks
}

// generates the eventgeneratortypes for cluster_settings.txt
func (t TheOne) buildEventGenerators(exp experiment.Experiment) (eventGenerators []eventGenerator) {
	logger.Trace("Building Event Generators")
	for _, expEventGenerator := range exp.EventGenerators {
		evg := eventGenerator{
			Name: expEventGenerator.Name,
			Type: expEventGenerator.Type,
		}

		eventGenerators = append(eventGenerators, evg)
	}
	return eventGenerators
}

// generates a txt for Theone with a given experiment
func (t TheOne) Generate(exp experiment.Experiment) {
	logger.Info("Generating TheOne output")
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	outputFilePath := filepath.Join(outputFolder, TheOneOutput)
	if !folderstructure.MayCreatePath(outputFilePath) {
		logger.Error("Not allowed to write output file!")
		return
	}
	logger.Tracef("Opening file \"%s\"", outputFilePath)
	fbuffer, err := os.Create(outputFilePath)
	if err != nil {
		logger.Error("Error creating output file:", err)
		return
	}
	defer func() {
		if cerr := fbuffer.Close(); cerr != nil {
			logger.Error("Error closing output file:", cerr)
		}
	}()

	// WorldSizes are multiplied by 20, because the Size of The ONE is about 20 times bigger compared to the values of clab, Core, etc.
	// WarmUp is multiplied by 200, because the warm up period for THE ONE is about 200 times longer compared to the values of clab, Core, etc.
	replace := theOneData{
		ScenarioName:     folderstructure.FileSystemEscape(exp.Name),
		RandomSeed:       exp.RandomSeed,
		Warmup:           exp.Warmup * 200,
		Runtime:          exp.Duration,
		NrofHostGroups:   len(exp.NodeGroups),
		WorldSizeHeight:  exp.WorldSize.Height * 20,
		WorldSizeWidth:   exp.WorldSize.Width * 20,
		Interfaces:       t.buildInterfaces(exp),
		Groups:           t.buildGroups(exp),
		NoEventGenerator: len(exp.EventGenerators),
		EventGenerators:  t.buildEventGenerators(exp),
	}

	funcs := template.FuncMap{"add": add}
	txtTemplate := template.New(TheOneOutput).Funcs(funcs)
	txtTemplate, err = txtTemplate.ParseFS(TemplatesFS, fmt.Sprintf("%s/%s", TemplatesFolder, TheOneOutput))
	if err != nil {
		logger.Error("Error opening template file:", err)
		return
	}
	err = txtTemplate.Execute(fbuffer, replace)
	if err != nil {
		logger.Error("Could not execute txt template:", err)
		return
	}
	logger.Trace("Finished generation")

}
