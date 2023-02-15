package outputgenerators_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestTheone(t *testing.T) {
	

	to := outputgenerators.Theone{}
	testingExperiment := GetTestingExperiment()
	outputFolder := folderstructure.GetOutputFolder(testingExperiment)
	to.Generate(testingExperiment)

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, outputgenerators.TheoneOutput))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile(filepath.Join(outputFolder, outputgenerators.TheoneOutput))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		t.Fatal("Output does not match expected output!")
	}
}
