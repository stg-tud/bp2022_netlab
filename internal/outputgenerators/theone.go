package outputgenerators

import (
	"os"
	"path/filepath"
	"reflect"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/customtypes"

	//"github.com/stg-tud/bp2022_netlab/internal/eventgenerators"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

type Theone struct{}

type replacenetwork struct {
	Bandwidth int
	Range     int
}
type groups struct {
	Id             string
	NrofHosts      uint
	NrofInterfaces int
	InterfaceID    uint
	Interface      []*experiment.Network
	MovementModel  string
}

type data struct {
	ScenarioName                string
	ScenarioSimulateConnections bool
	ScenarioUpdateInterval      float64
	ScenarioEndTime             string
	WorldSizeHeight             uint
	WorldSizeWidth              uint
	NrofHostGroups              int
	Groups                      []groups
	Interfaces                  []networkInterFace
	EventGenerator              []eventGeneraTor
	RandomSeed                  int64
	Warmup                      uint
	Runtime                     uint
	NoEventGenerator            int
}
type eventGeneraTor struct {
	Name     string
	Interval uint
	Size     customtypes.Area
	Hosts    customtypes.Area
	Prefix   string
}
type networkInterFace struct {
	Name      string
	Bandwidth int
	Range     int
}

var InterfaceIdCounter uint = 1

func (t Theone) getInterfaceId() uint {
	InterfaceIdCounter++
	return InterfaceIdCounter - 1
}

// The name of the file that should be written to
const TheoneOutput = "cluster_settings.txt"

// The default values for sluster_settings.txt
var defaultValuesTheone = data{

	ScenarioSimulateConnections: true,
	ScenarioUpdateInterval:      0.1,
}

func (t Theone) movementpattern(movementpatterntype movementpatterns.MovementPattern) string {
	switch movementpatterntype.(type) {
	case movementpatterns.RandomWaypoint:
		return "RandomWaypoint"
	case movementpatterns.Static:
		return "Static"
	default:
		return ""
	}
}

func (t Theone) BuildGroups(exp experiment.Experiment) []groups {
	logger.Trace("Building Groups")
	groupInterface := []groups{}

	for i := 0; i < len(exp.NodeGroups); i++ {

		expNodeGroups := &exp.NodeGroups[i]
		group := groups{
			Id:        expNodeGroups.Prefix,
			NrofHosts: expNodeGroups.NoNodes,

			InterfaceID:   t.getInterfaceId(),
			Interface:     expNodeGroups.Networks,
			MovementModel: t.movementpattern(expNodeGroups.MovementModel),
		}
		groupInterface = append(groupInterface, group)

	}
	return groupInterface
}

// generates the networks for theone.txt
func (t Theone) BuildNetworks(exp experiment.Experiment) (networks []networkInterFace) {

	logger.Trace("Building Interfaces")
	for i := 0; i < len(exp.Networks); i++ {
		bandwidth := 250
		rangeOfNetwork := 10
		if reflect.TypeOf(exp.Networks[i].Type).String() == "networktypes.WirelessLAN" || reflect.TypeOf(exp.Networks[i].Type).String() == "networktypes.Wireless" {
			bandwidth = int(reflect.ValueOf(exp.Networks[i].Type).FieldByName("Bandwidth").Int()) / 100000
			rangeOfNetwork = int(reflect.ValueOf(exp.Networks[i].Type).FieldByName("Range").Int()) / 20
		}

		nt := networkInterFace{
			Name:      exp.Networks[i].Name,
			Bandwidth: bandwidth,
			Range:     rangeOfNetwork,
		}
		networks = append(networks, nt)
	}

	return networks
}

func (t Theone) BuildEventGenerator(exp experiment.Experiment) (eventGenerator []eventGeneraTor) {
	logger.Trace("Building Event Generators")
	for i := 0; i < len(exp.EventGenerators); i++ {
		evg := eventGeneraTor{
			Name: exp.EventGenerators[i].Name,
		}
		eventGenerator = append(eventGenerator, evg)
	}
	return eventGenerator
}

// generates a txt for Theone with a given experiment
func (t Theone) Generate(exp experiment.Experiment) {
	logger.Info("Generating Theone output")
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	outputFilePath := filepath.Join(outputFolder, TheoneOutput)
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
			logger.Error("Error closing output file:", cerr)
			err = cerr
		}
	}()

	if err != nil {
		panic(err)
	}
	replace := defaultValuesTheone
	replace.ScenarioName = folderstructure.FileSystemEscape(exp.Name)
	replace.RandomSeed = exp.RandomSeed
	replace.Warmup = exp.Warmup * 200
	replace.Runtime = exp.Duration
	replace.NrofHostGroups = len(exp.NodeGroups)

	replace.WorldSizeHeight = uint(exp.WorldSize.Height * 20)
	replace.WorldSizeWidth = uint(exp.WorldSize.Width * 20)
	replace.Interfaces = t.BuildNetworks(exp)

	replace.Groups = t.BuildGroups(exp)

	replace.EventGenerator = t.BuildEventGenerator(exp)
	for i, node := range exp.NodeGroups {

		replace.Groups[i].NrofInterfaces = len(node.Networks)
	}
	replace.NoEventGenerator = len(exp.EventGenerators)

	txtTemplate, err := template.ParseFiles(filepath.Join(GetTemplatesFolder(), "cluster_settings.txt"))
	if err != nil {
		logger.Error("Error opening template file:", err)
	}
	err = txtTemplate.Execute(fbuffer, replace)
	if err != nil {
		logger.Error("Could not execute txt template:", err)
		return
	}
	logger.Trace("Finished generation")

}
