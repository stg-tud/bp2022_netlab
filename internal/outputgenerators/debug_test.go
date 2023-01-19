package outputgenerators_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestDebugGeneration(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(outputgenerators.OutputFolder)
	})

	og := outputgenerators.Debug{}
	og.Generate(GetTestingExperiment())

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, outputgenerators.DebugOutputFile))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile(filepath.Join(outputgenerators.OutputFolder, outputgenerators.DebugOutputFile))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		t.Fatal("Output does not match expected output!")
	}
}
