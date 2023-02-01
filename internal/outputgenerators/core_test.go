package outputgenerators_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestCore(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	og := outputgenerators.Core{}
	testingExperiment := GetTestingExperiment()
	outputFolder := folderstructure.GetOutputFolder(testingExperiment)
	og.Generate(testingExperiment)

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, "core.xml"))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile(filepath.Join(outputFolder, "core.xml"))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		t.Fatal("Output does not match expected output!")
	}
}
