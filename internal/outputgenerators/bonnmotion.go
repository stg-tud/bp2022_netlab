package outputgenerators

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// The name of the executable to run BonnMotion.
const BonnMotionExecutable = "bonnmotion"

// The name of the file that the taken steps should be written into.
const BonnMotionStepFile = "bonnmotion.steps"

// The Bonnmotion output generator calles BonnMotion with the correct parameters.
type Bonnmotion struct {
	outputFolder string
	stepFilePath string
}

// Returns the correct BonnMotion platform name for the given Target.
func (Bonnmotion) platform(t experiment.Target) (bool, string) {
	switch t {
	case experiment.TargetTheOne:
		return true, "TheONEFile"

	case experiment.TargetCore:
		return true, "NSFile"

	default:
		return false, ""
	}
}

// Returns whether the given Target is (currently) supported by this output generator.
func (b Bonnmotion) IsSupported(t experiment.Target) bool {
	supported, _ := b.platform(t)
	return supported
}

// Returns the parameter set for Random Waypoint movement model for a given NodeGroup inside an Experiment.
func (Bonnmotion) randomWaypointParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	movementmodel := nodeGroup.MovementModel.(movementpatterns.RandomWaypoint)
	return []string{
		"RandomWaypoint",
		fmt.Sprintf("-h%d", movementmodel.MaxSpeed),
		fmt.Sprintf("-l%d", movementmodel.MinSpeed),
		fmt.Sprintf("-p%d", movementmodel.MaxPause),
	}
}

// Returns the general parameters for a given NodeGroup inside an Experiment.
func (Bonnmotion) generalParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	return []string{
		fmt.Sprintf("-d%d", exp.Duration),
		fmt.Sprintf("-n%d", nodeGroup.NoNodes),
		fmt.Sprintf("-x%d", exp.WorldSize.Width),
		fmt.Sprintf("-y%d", exp.WorldSize.Height),
		fmt.Sprintf("-R%d", exp.RandomSeed),
	}
}

// Writes the command to the step file and executes it
func (b Bonnmotion) execute(command []string) error {
	logger.Trace("Running command:", command)
	logger.Tracef("Writing file \"%s\"", b.stepFilePath)
	stepFile, err := os.OpenFile(b.stepFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Error opening step file:", err)
		return err
	}
	defer func() {
		if cerr := stepFile.Close(); cerr != nil {
			logger.Error("Error closing step file:", cerr)
			err = cerr
		}
	}()
	_, err = stepFile.WriteString(fmt.Sprintln(command))
	if err != nil {
		logger.Error("Error writing step file:", err)
		return err
	}
	execCommand := exec.Command(BonnMotionExecutable, command...)
	execCommand.Dir = b.outputFolder
	// Check if the function is currently unit tested and do not execute actual BonnMotion command if so.
	if flag.Lookup("test.v") != nil {
		logger.Debug("Detected test. Skipping actual command execution")
		return nil
	}
	_, err = execCommand.Output()
	return err
}

// Calls BonnMotion to generate the Random Waypoint data for a given NodeGroup inside an Experiment.
func (b Bonnmotion) generateRandomWaypointNodeGroup(exp experiment.Experiment, nodeGroup experiment.NodeGroup) {
	logger.Trace("Generating Random Waypoint movements")
	command := []string{
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	command = append(command, b.randomWaypointParameters(exp, nodeGroup)...)
	command = append(command, b.generalParameters(exp, nodeGroup)...)
	err := b.execute(command)
	if err != nil {
		logger.Error("Error running command:", err)
	}
}

// Calls BonnMotion to convert the BonnMotion output to the given Target's format for a given NodeGroup.
func (b Bonnmotion) convertToTargetFormat(target experiment.Target, nodeGroup experiment.NodeGroup) {
	logger.Tracef("Converting to target format \"%s\"", target.String())
	supported, model := b.platform(target)
	if !supported {
		logger.Debug("Target platform is currently not supported. Skipping\n")
		return
	}
	command := []string{
		model,
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	err := b.execute(command)
	if err != nil {
		logger.Error("Error running command:", err)
	}
}

// Generate generates output for the given Experiment with BonnMotion.
func (b Bonnmotion) Generate(exp experiment.Experiment) {
	logger.Info("Generating BonnMotion output")
	outputFolder, err := folderstructure.GetAndCreateOutputFolder(exp, "movements")
	if err != nil {
		logger.Error("Could not create output folder!", err)
		return
	}
	b.outputFolder = outputFolder

	b.stepFilePath = filepath.Join(b.outputFolder, BonnMotionStepFile)
	allowedToWriteStepFile := folderstructure.MayCreatePath(b.stepFilePath)
	if !allowedToWriteStepFile {
		logger.Error("Not allowed to write step file!")
		return
	}
	for _, nodeGroup := range exp.NodeGroups {
		logger.Tracef("Processing NodeGroup \"%s\"", nodeGroup.Prefix)
		if !folderstructure.MayCreatePath(filepath.Join(b.outputFolder, fmt.Sprintf("%s.movements.gz", nodeGroup.Prefix))) {
			logger.Error("Not allowed to write output file!")
			return
		}
		switch nodeGroup.MovementModel.(type) {
		case movementpatterns.RandomWaypoint:
			b.generateRandomWaypointNodeGroup(exp, nodeGroup)
		default:
			logger.Debugf("Movement model \"%s\" is currently not supported. Skipping", reflect.TypeOf(nodeGroup.MovementModel))
			continue
		}
		for _, target := range exp.Targets {
			b.convertToTargetFormat(target, nodeGroup)
		}
	}
	logger.Trace("Finished generation")
}
