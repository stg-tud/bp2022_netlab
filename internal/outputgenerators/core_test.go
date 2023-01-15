package outputgenerators_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestCore(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(outputgenerators.OUTPUT_FOLDER)
	})

	og := outputgenerators.Core{}
	og.Generate(GetTestingExperiment())

	expected, err := os.ReadFile(fmt.Sprintf("%s/core.xml", TESTDATA_FOLDER))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	actual, err := os.ReadFile(fmt.Sprintf("%s/core.xml", outputgenerators.OUTPUT_FOLDER))
	if err != nil {
		t.Fatal("Could not read output file", err)
	}

	if string(actual) != string(expected) {
		t.Fatal("Output does not match expected output!")
	}
}
