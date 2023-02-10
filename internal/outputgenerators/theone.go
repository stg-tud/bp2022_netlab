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

type data struct {
	ScenarioName                string
	ScenarioSimulateConnections bool
	ScenarioUpdateInterval      float64
	ScenarioEndTime             string

	FirstinterfaceType string

	GUI           uint
	RandomSeed    int64
	PidStat       uint
	PidParam      string
	Net           uint
	NetParam      string
	XY            uint
	XYParam       uint
	Contacts      uint
	ContactsParam uint
	Shutdown      string
	Warmup        uint
	Runtime       uint
}

// The name of the file that should be written to
const TheoneOutput = "cluster_settings.txt"

// The default values for sluster_settings.txt
var defaultValuesTheone = data{

	ScenarioSimulateConnections: true,
	ScenarioUpdateInterval:      0.1,
	ScenarioEndTime:             "43k",

	GUI:           1,
	PidStat:       0,
	PidParam:      "vnoded",
	Net:           0,
	NetParam:      "eth0",
	XY:            1,
	XYParam:       5,
	Contacts:      1,
	ContactsParam: 5,
	Shutdown:      "",
}

// generates a txt for Theone with a given experiment
func (Theone) Generate(exp experiment.Experiment) {
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
	replace.Warmup = exp.Warmup
	replace.Runtime = exp.Duration

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
