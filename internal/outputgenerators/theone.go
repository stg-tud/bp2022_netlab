package outputgenerators

import (
	"os"
	"path/filepath"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/eventgenerators"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
)

type TheOne struct{}

type group struct {
	Index          uint
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
	Index     uint
	Name      string
	Type      eventgenerators.EventGeneratorType
	IntervalX int
	IntervalY int
	SizeX     int
	SizeY     int
	HostsX    int
	HostsY    int
	ToHostsX  int
	ToHostsY  int
	Prefix    string
}
type networkInterfaceTheOne struct {
	Name      string
	Bandwidth int
	Range     int
}

// The name of the file that should be written to
const TheOneOutput = "cluster_settings.txt"

func (TheOne) String() string {
	return "the-one"
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
	groupInterface := []group{}

	for i, expNodeGroups := range exp.NodeGroups {
		group := group{
			Index:          uint(i + 1),
			Id:             expNodeGroups.Prefix,
			NrofHosts:      expNodeGroups.NoNodes,
			Interfaces:     expNodeGroups.Networks,
			MovementModel:  t.movementPattern(expNodeGroups.MovementModel),
			NrofInterfaces: len(expNodeGroups.Networks),
		}
		groupInterface = append(groupInterface, group)

	}
	return groupInterface
}

// generates the interfaces for cluster_settings.txt
// bandwidth and range are both divided by a certain factor, because the ratio of the one is different
func (t TheOne) buildInterfaces(exp experiment.Experiment) (networks []networkInterfaceTheOne) {

	logger.Trace("Building Interfaces")
	for i := range exp.Networks {
		bandwidth := 6570
		rangeOfNetwork := 100
		switch networkType := exp.Networks[i].Type.(type) {
		case networktypes.WirelessLAN:
			bandwidth = networkType.Bandwidth
			rangeOfNetwork = networkType.Range

		case networktypes.Wireless:
			bandwidth = networkType.Bandwidth
			rangeOfNetwork = networkType.Range
		}

		nt := networkInterfaceTheOne{
			Name:      exp.Networks[i].Name,
			Bandwidth: bandwidth,
			Range:     rangeOfNetwork,
		}
		networks = append(networks, nt)
	}

	return networks
}

// generates the eventgenerators for cluster_settings.txt
func (t TheOne) buildEventGenerator(exp experiment.Experiment) (eventGenerators []eventGenerator) {
	logger.Trace("Building Event Generators")
	for i, expEventGenerator := range exp.EventGenerators {
		evg := eventGenerator{
			Index: uint(i + 1),
			Name:  expEventGenerator.Name,
			Type:  expEventGenerator.Type,
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

	// WorldSizes are multiplied by twenty because the Size of The One is about 20 times bigger
	//WarmUp is multiplied by 200, because the warm up period for the one is about 200 times longer
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
		EventGenerators:  t.buildEventGenerator(exp),
	}

	funcs := template.FuncMap{"add": add}
	txtTemplate := template.Must(template.New(TheOneOutput).Funcs(funcs).ParseFiles(filepath.Join(TemplatesFolder, TheOneOutput)))
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
