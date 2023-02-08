package outputgenerators_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestCoreemulabGeneration(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})
	coreemulab := outputgenerators.CoreEmulab{}
	testingExperiment := GetTestingExperiment()
	outputFolder := folderstructure.GetOutputFolder(testingExperiment)
	coreemulab.Generate(testingExperiment)

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, "experiment.conf"))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile(filepath.Join(outputFolder, "experiment.conf"))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		println(expected)
		t.Fatal("Output does not match expected output!")
	}

}
