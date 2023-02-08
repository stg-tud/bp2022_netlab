package outputgenerators_test

import (
	"os"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestOutput(t *testing.T) {

	core := outputgenerators.CoreEmulab{}

	exp := experiment.Experiment{
		Name:     "Automator",
		Duration: 245,
	}
	core.Generate(exp)

	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

}
