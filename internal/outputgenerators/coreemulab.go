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
	logger.Info("Generating CORE output")
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp)
	os.Mkdir(outputFolder, 0755)

	fbuffer, err := os.Create(filepath.Join(outputFolder, "experiment.conf"))
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
