package cmd

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

var stringTest string

func TestStringTargetMappingTHEONE(t *testing.T) {

	stringTest = "the-one"
	expectedTheOne, errTheOne := stringTargetMapping(stringTest)

	if expectedTheOne != experiment.TargetTheOne {
		t.Fatal("Wrong Experiment Target", errTheOne)
	}

	stringTest = "the-lne"
	_, err := stringTargetMapping(stringTest)
	if err == nil {
		t.Fatal("created a target even though no correct one was given", err)
	}

	stringTest = "core"
	expectedCore, errCore := stringTargetMapping(stringTest)

	if expectedCore != experiment.TargetCore {
		t.Fatal("Wrong Experiment Target", errCore)
	}

	stringTest = "core-emulab"
	expectedClab, errClab := stringTargetMapping(stringTest)

	if expectedClab != experiment.TargetCoreEmulab {
		t.Fatal("Wrong Experiment Target", errClab)
	}

}
