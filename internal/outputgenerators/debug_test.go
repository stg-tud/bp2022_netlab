package outputgenerators_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestDebug(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(outputgenerators.OUTPUT_FOLDER)
	})

	og := outputgenerators.Debug{}
	og.Generate(GetTestingExperiment())

	expected, err := os.ReadFile(fmt.Sprintf("%s/debug_out.toml", TESTDATA_FOLDER))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile(fmt.Sprintf("%s/debug_out.toml", outputgenerators.OUTPUT_FOLDER))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		t.Fatal("Output does not match expected output!")
	}
}
