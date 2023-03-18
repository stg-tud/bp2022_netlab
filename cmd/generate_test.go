package cmd

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

var stringTest string

//var exTarget experiment.Target

var exptest experiment.Experiment = experiment.Experiment{}

// var outTest []outputgenerators.OutputGenerator
var boolCore bool
var boolClab bool
var boolTheOne bool

// test
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

func TestBuildTargets(t *testing.T) {

	exptest.Targets = append(exptest.Targets, experiment.TargetCore)
	exptest.Targets = append(exptest.Targets, experiment.TargetTheOne)

	var expectedTar [2]experiment.Target
	expectedTar[0] = experiment.TargetCore
	expectedTar[1] = experiment.TargetTheOne

	actual := buildTargets(exptest)

	if len(expectedTar) != len(actual) {
		t.Fatal("wrong number of Targets")
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expectedTar[i] {
			t.Fatal("Wrong Target")
		}
	}

}

func TestBuildOutputGenerators(t *testing.T) {

	var outgen1 outputgenerators.OutputGenerator = outputgenerators.Core{}
	var outgen2 outputgenerators.OutputGenerator = outputgenerators.CoreEmulab{}
	var outgen3 outputgenerators.OutputGenerator = outputgenerators.TheOne{}
	var outgen4 outputgenerators.OutputGenerator = outputgenerators.Bonnmotion{}

	exptest.Targets = append(exptest.Targets, experiment.TargetCore)
	exptest.Targets = append(exptest.Targets, experiment.TargetTheOne)
	exptest.Targets = append(exptest.Targets, experiment.TargetCoreEmulab)

	var nodetest experiment.NodeGroup
	nodetest.MovementModel = movementpatterns.RandomWaypoint{}

	exptest.NodeGroups = append(exptest.NodeGroups, nodetest)
	// exptest.Targets = append(exptest.Targets, experiment.TargetB
	// man kann in den bonnmotion fall nie reinlaufen?

	outTest := buildOutputGenerators(exptest)

	boolCore = false
	boolClab = false
	boolTheOne = false
	var boolBonnmotion = false

	for i := 0; i < len(outTest); i++ {
		if outgen2 == outTest[i] {
			boolClab = true
		}
		if outgen3 == outTest[i] {
			boolTheOne = true
		}
		if outgen1 == outTest[i] {
			boolCore = true
		}
		if outgen4 == outTest[i] {
			boolBonnmotion = true
		}
	}

	if !boolClab {
		t.Fatal(" created  non outputgenerator CoreEmuLab ")
	}

	if !boolCore {
		t.Fatal(" created  non outputgenerator Core ")
	}

	if !boolTheOne {
		t.Fatal(" created  non outputgenerator TheOne")
	}

	if !boolBonnmotion {
		t.Fatal(" created  non outputgenerator TheOne")
	}
}
