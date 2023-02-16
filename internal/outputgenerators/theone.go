package outputgenerators

import (
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
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

	for i := range exp.NodeGroups {

		expNodeGroups := exp.NodeGroups[i]

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
		if reflect.TypeOf(exp.Networks[i].Type).String() == "networktypes.WirelessLAN" || reflect.TypeOf(exp.Networks[i].Type).String() == "networktypes.Wireless" {
			bandwidth = int(reflect.ValueOf(exp.Networks[i].Type).FieldByName("Bandwidth").Int()) / 100000
			rangeOfNetwork = int(reflect.ValueOf(exp.Networks[i].Type).FieldByName("Range").Int()) / 20
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
	for i := range exp.EventGenerators {
		evg := eventGenerator{
			Index: uint(i + 1),
			Name:  exp.EventGenerators[i].Name,

			Prefix: reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Prefix").String(),

			SizeX: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Size").FieldByName("XSeconds").Int()),
			SizeY: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Size").FieldByName("YSeconds").Int()),

			HostsX: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Hosts").FieldByName("XSeconds").Int()),
			HostsY: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Hosts").FieldByName("YSeconds").Int()),

			ToHostsX: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("ToHosts").FieldByName("XSeconds").Int()),
			ToHostsY: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("ToHosts").FieldByName("YSeconds").Int()),

			IntervalX: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Interval").FieldByName("XSeconds").Int()),
			IntervalY: int(reflect.ValueOf(exp.EventGenerators[i].Type).FieldByName("Interval").FieldByName("YSeconds").Int()),
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
	txtTemplate := template.Must(template.New(TheOneOutput).Funcs(funcs).ParseFiles(filepath.Join(GetTemplatesFolder(), TheOneOutput)))
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
