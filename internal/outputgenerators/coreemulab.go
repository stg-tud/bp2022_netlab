package outputgenerators

import (
	"os"
	"path/filepath"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
)

type CoreEmulab struct{}

type data struct {
	Name      string
	Scenario  string
	Automator string

	GUI           uint
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

// The default values fpr experiment.conf
var defaultValues = data{

	Scenario:  "core.xml",
	Automator: "",

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

// generates a XML and a conf configuartion for CoreEmulab with a given experiment
func (c CoreEmulab) Generate(exp experiment.Experiment) {
	logger.Info("Generating CoreEmulab output")
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	outputFilePath := filepath.Join(outputFolder, "experiment.conf")
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
			logger.Error("Error closing step file:", cerr)
			err = cerr
		}
	}()

	if err != nil {
		panic(err)
	}
	replace := defaultValues
	replace.Name = exp.Name
	replace.Warmup = exp.Warmup
	replace.Runtime = exp.Duration

	confTemplate, err := template.ParseFiles(filepath.Join(GetTemplatesFolder(), "experiment.conf"))
	if err != nil {
		logger.Error("Error opening template file:", err)
	}
	err = confTemplate.Execute(fbuffer, replace)
	if err != nil {
		logger.Error("Could not execute conf template:", err)
		return
	}
	logger.Trace("Finished generation")

}
