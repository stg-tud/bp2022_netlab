package outputgenerators_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestDebugGeneration(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	og := outputgenerators.Debug{}
	testingExperiment := GetTestingExperiment()
	outputFolder := folderstructure.GetOutputFolder(testingExperiment)
	og.Generate(testingExperiment)

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, outputgenerators.DebugOutputFile))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}
	expectedClean := strings.ReplaceAll(string(expected), "\r\n", "\n")

	actual, err := os.ReadFile(filepath.Join(outputFolder, outputgenerators.DebugOutputFile))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}
	actualClean := strings.ReplaceAll(string(actual), "\r\n", "\n")

	if string(actualClean) != string(expectedClean) {
		t.Fatal("Output does not match expected output!")
	}
}
