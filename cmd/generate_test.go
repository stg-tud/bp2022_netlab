package cmd

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

var stringTest string
var exTarget experiment.Target

func TestStringTargetMappingTHEONE(t *testing.T) {

	stringTest = "the-one"
	expected, err := stringTargetMapping(stringTest)

	if expected != experiment.TargetTheOne {
		t.Fatal("Wrong output Experiment", err)
	}

}

func TestStringTargetMappingEmpty(t *testing.T) {

	stringTest = "the-lne"
	_, err := stringTargetMapping(stringTest)
	if err == nil {
		t.Fatal("Didn't create Experiment", err)
	}

}

func TestStringTargetMappingCore(t *testing.T) {

	stringTest = "core"
	expected, err := stringTargetMapping(stringTest)

	if expected != experiment.TargetCore {
		t.Fatal("Wrong output Experiment", err)
	}

}

func TestStringTargetMappingCoreemu(t *testing.T) {

	stringTest = "core-emulab"
	expected, err := stringTargetMapping(stringTest)

	if expected != experiment.TargetCoreEmulab {
		t.Fatal("Wrong output Experiment", err)
	}

}

func TestTargetOutputGeneratorMappingCore(t *testing.T) {

	exTarget = experiment.TargetCore

	expected, err := targetOutputGeneratorMapping(exTarget)

	if (expected != outputgenerators.Core{}) {
		t.Fatal(" created wrong or non outputgenerators", err)
	}

}

func TestTargetOutputGeneratorMappingCoreEmulab(t *testing.T) {

	exTarget = experiment.TargetCoreEmulab

	expected, err := targetOutputGeneratorMapping(exTarget)

	if (expected != outputgenerators.CoreEmulab{}) {
		t.Fatal(" created wrong or non outputgenerators", err)
	}

}

func TestTargetOutputGeneratorMappingOne(t *testing.T) {

	exTarget = experiment.TargetTheOne

	expected, err := targetOutputGeneratorMapping(exTarget)

	if (expected != outputgenerators.TheOne{}) {
		t.Fatal(" created wrong or non outputgenerators", err)
	}

}
