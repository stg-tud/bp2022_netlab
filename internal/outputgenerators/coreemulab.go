package outputgenerators

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"text/template"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
)

type CoreEmulab struct{}

func (CoreEmulab) String() string {
	return "coreemu-lab"
}

type coreEmuData struct {
	Name      string
	Scenario  string
	Automator string

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

// The name of the template to generate the config
const CoreEmulabTemplate = "experiment.conf"

// The name of the file that should be written to
const CoreEmulabOutput = "experiment.conf"

// The default values for experiment.conf
var defaultValuesCoreEmulab = coreEmuData{

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

// generates a conf configuartion for CoreEmulab with a given experiment
func (c CoreEmulab) Generate(exp experiment.Experiment) {
	logger.Info("Generating CoreEmulab output")
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	outputFilePath := filepath.Join(outputFolder, CoreEmulabOutput)
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
	replace := defaultValuesCoreEmulab
	replace.Name = folderstructure.FileSystemEscape(exp.Name)
	replace.RandomSeed = exp.RandomSeed
	replace.Warmup = exp.Warmup
	replace.Runtime = exp.Duration

	if exp.ExternalMovement.Active {
		replace.Automator = exp.ExternalMovement.FileName
		path, err := os.Getwd()
		if err != nil {
			logger.Error("Error getting current directory")
		}
		source, err := os.Open(filepath.Join(path, exp.ExternalMovement.FileName))
		if err != nil {
			logger.Error("Error opening external file", err)
		}
		defer source.Close()
		destination, err := os.Create(filepath.Join(outputFolder, exp.ExternalMovement.FileName))
		if err != nil {
			logger.Error("Error creating external file", err)
		}
		defer destination.Close()

		_, err = io.Copy(destination, source)
		if err != nil {
			logger.Error("Error copying external file", err)
		}

	}
	confTemplate, err := template.ParseFS(TemplatesFS, fmt.Sprintf("%s/%s", TemplatesFolder, "experiment.conf"))
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
