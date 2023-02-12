package outputgenerators

import (
	"os"
	"path/filepath"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
)

type Theone struct{}

type groups struct {
	Id        string
	NrofHosts uint
	Interface string
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
	RandomSeed                  int64
	Warmup                      uint
	Runtime                     uint
}

type networkInterFace struct {
	Name string
	Bandwidth int
	Range     int
}

// The name of the file that should be written to
const TheoneOutput = "cluster_settings.txt"

// The default values for sluster_settings.txt
var defaultValuesTheone = data{

	ScenarioSimulateConnections: true,
	ScenarioUpdateInterval:      0.1,
	ScenarioEndTime:             "43k",
}

func (t Theone) BuildGroups(exp experiment.Experiment) []groups {
	logger.Trace("Building Groups")
	groupInterface := []groups{}
	for i := 0; i < len(exp.NodeGroups); i++ {
		group := groups{
			Id:        exp.NodeGroups[i].Prefix,
			NrofHosts: exp.NodeGroups[i].NoNodes,
		}
		groupInterface = append(groupInterface, group)

	}
	return groupInterface
}

func (t Theone)BuildNetworks(exp experiment.Experiment) (networks []networkInterFace){

	logger.Trace("Building Interfaces")
	for i := 0; i < len(exp.Networks); i++ {
		nt:=networkInterFace{
			Name: exp.Networks[i].Name,
		}
		networks = append(networks, nt)
	}
	
return networks
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

	replace.Groups = t.BuildGroups(exp)

	replace.Interfaces = t.BuildNetworks(exp)


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
