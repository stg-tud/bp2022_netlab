package outputgenerators

import (
	"os"
	"path/filepath"
	"text/template"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

type CoreEmulab struct{}

type data struct {
	Name      string
	Scenario  string
	Automator string
	GUI       int
	PidStat   string
	PidParam  string
	Net       int
	NetParam  string
	XY        int
	Contacts  int
	Shutdown  string
	Runtime   int
}

// generates a XML and a conf configuartion for CoreEmulab with a given experiment
func (c CoreEmulab) Generate(exp experiment.Experiment) {

	os.Mkdir(OutputFolder, 0755)

	fbuffer, err := os.Create(filepath.Join(OutputFolder, "experiment.conf"))
	if err != nil {
		panic(err)
	}
	replace := data{
		Name: exp.Name,

		Runtime: exp.Duration,
	}

	confTemplate, err := template.ParseFiles(filepath.Join(GetTemplatesFolder(), "experiment.conf"))
	if err != nil {
		panic(err)
	}
	confTemplate.Execute(fbuffer, replace)

}
