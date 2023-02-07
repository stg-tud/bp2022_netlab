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

	GUI           int
	PidStat       string
	PidParam      string
	Net           int
	NetParam      string
	XY            int
	XYParam       int
	Contacts      int
	ContactsParam int
	Shutdown      string
	Warmup        int
	Runtime       int
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
	replace := data{
		Name: exp.Name,

		Runtime: int(exp.Duration),
	}

	confTemplate, err := template.ParseFiles(filepath.Join(GetTemplatesFolder(), "experiment.conf"))
	if err != nil {
		panic(err)
	}
	confTemplate.Execute(fbuffer, replace)

}
