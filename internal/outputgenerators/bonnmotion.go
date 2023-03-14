package outputgenerators

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	logger "github.com/gookit/slog"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// The name of the executable to run BonnMotion.
var BonnMotionExecutable = "bm"

// The name of the file that the taken steps should be written into.
const BonnMotionStepFile = "bonnmotion.steps"

// The Bonnmotion output generator calles BonnMotion with the correct parameters.
type Bonnmotion struct {
	outputFolder string
	stepFile     *os.File
}

func (Bonnmotion) String() string {
	return "BonnMotion"
}

// Returns the correct BonnMotion platform name for the given Target.
func (Bonnmotion) platform(t experiment.Target) (bool, string) {
	switch t {
	case experiment.TargetTheOne:
		return true, "TheONEFile"

	case experiment.TargetCore, experiment.TargetCoreEmulab:
		return true, "NSFile"

	default:
		return false, ""
	}
}

// Returns whether the given Target is (currently) supported by this output generator.
func (b Bonnmotion) TargetIsSupported(t experiment.Target) bool {
	supported, _ := b.platform(t)
	return supported
}

// Returns whether the given MovementPattern is (currently) supported by this output generator.
func (b Bonnmotion) MovementPatternIsSupported(movementPattern movementpatterns.MovementPattern) bool {
	switch movementPattern.(type) {
	case movementpatterns.RandomWaypoint, movementpatterns.SMOOTH, movementpatterns.SLAW, movementpatterns.SWIM:
		return true

	default:
		return false
	}
}

// Returns the parameter set for SMOOTH movement model for a given NodeGroup inside an Experiment.
func (Bonnmotion) smoothParameters(movementModel movementpatterns.SMOOTH) []string {
	return []string{
		"SMOOTH",
		"-g",
		fmt.Sprintf("%d", movementModel.Range),
		"-h",
		fmt.Sprintf("%d", movementModel.Clusters),
		"-k",
		fmt.Sprintf("%v", movementModel.Alpha),
		"-l",
		fmt.Sprintf("%d", movementModel.MinFlight),
		"-m",
		fmt.Sprintf("%d", movementModel.MaxFlight),
		"-o",
		fmt.Sprintf("%v", movementModel.Beta),
		"-p",
		fmt.Sprintf("%d", movementModel.MinPause),
		"-q",
		fmt.Sprintf("%d", movementModel.MaxPause),
	}
}

// Returns the parameter set for SLAW movement model for a given NodeGroup inside an Experiment.
func (Bonnmotion) slawParameters(movementModel movementpatterns.SLAW) []string {
	return []string{
		"SLAW",
		"-w",
		fmt.Sprintf("%d", movementModel.NumberOfWaypoints),
		"-p",
		fmt.Sprintf("%d", movementModel.MinPause),
		"-P",
		fmt.Sprintf("%d", movementModel.MaxPause),
		"-b",
		fmt.Sprintf("%v", movementModel.LevyExponent),
		"-h",
		fmt.Sprintf("%v", movementModel.HurstParameter),
		"-l",
		fmt.Sprintf("%v", movementModel.DistanceWeight),
		"-r",
		fmt.Sprintf("%v", movementModel.ClusteringRange),
		"-Q",
		fmt.Sprintf("%v", movementModel.ClusterRatio),
		"-W",
		fmt.Sprintf("%v", movementModel.WaypointRatio),
	}
}

// Returns the parameter set for SWIM movement model for a given NodeGroup inside an Experiment.
func (Bonnmotion) swimParameters(movementModel movementpatterns.SWIM) []string {
	return []string{
		"SWIM",
		"-r",
		fmt.Sprintf("%v", movementModel.Radius),
		"-c",
		fmt.Sprintf("%v", movementModel.CellDistanceWeight),
		"-m",
		fmt.Sprintf("%v", movementModel.NodeSpeedMultiplier),
		"-e",
		fmt.Sprintf("%v", movementModel.WaitingTimeExponent),
		"-u",
		fmt.Sprintf("%v", movementModel.WaitingTimeUpperBound),
	}
}

// Returns the parameter set for Random Waypoint movement model for a given NodeGroup inside an Experiment.
func (Bonnmotion) randomWaypointParameters(movementModel movementpatterns.RandomWaypoint) []string {
	return []string{
		"RandomWaypoint",
		"-h",
		fmt.Sprintf("%d", movementModel.MaxSpeed),
		"-l",
		fmt.Sprintf("%d", movementModel.MinSpeed),
		"-p",
		fmt.Sprintf("%d", movementModel.MaxPause),
	}
}

// Returns the general parameters for a given NodeGroup inside an Experiment.
func (Bonnmotion) generalParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	return []string{
		"-d",
		fmt.Sprintf("%d", exp.Duration),
		"-n",
		fmt.Sprintf("%d", nodeGroup.NoNodes),
		"-x",
		fmt.Sprintf("%d", exp.WorldSize.Width),
		"-y",
		fmt.Sprintf("%d", exp.WorldSize.Height),
		"-R",
		fmt.Sprintf("%d", exp.RandomSeed),
	}
}

// Writes the command to the step file and executes it
func (b Bonnmotion) execute(command []string) error {
	commandString := fmt.Sprintf("%s %s", BonnMotionExecutable, strings.Join(command, " "))
	logger.Trace("Running command:", commandString)
	_, err := b.stepFile.WriteString(fmt.Sprintln(commandString))
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

// Calls BonnMotion to generate the for a given NodeGroup inside an Experiment.
func (b Bonnmotion) generateNodeGroup(exp experiment.Experiment, nodeGroup experiment.NodeGroup) error {
	logger.Tracef("Generating \"%s\" movements", nodeGroup.MovementModel.String())
	var movementModelParameters []string
	switch movementModel := nodeGroup.MovementModel.(type) {
	case movementpatterns.RandomWaypoint:
		movementModelParameters = b.randomWaypointParameters(movementModel)
	case movementpatterns.SMOOTH:
		movementModelParameters = b.smoothParameters(movementModel)
	case movementpatterns.SLAW:
		movementModelParameters = b.slawParameters(movementModel)
	case movementpatterns.SWIM:
		movementModelParameters = b.swimParameters(movementModel)
	default:
		return fmt.Errorf("movement model \"%s\" is not supported", nodeGroup.MovementModel.String())
	}
	command := []string{
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	command = append(command, movementModelParameters...)
	command = append(command, b.generalParameters(exp, nodeGroup)...)
	err := b.execute(command)
	return err
}

// Calls BonnMotion to convert the BonnMotion output to the given Target's format for a given NodeGroup.
func (b Bonnmotion) convertToTargetFormat(target experiment.Target, nodeGroup experiment.NodeGroup) error {
	logger.Tracef("Converting to target format \"%s\"", target.String())
	supported, model := b.platform(target)
	if !supported {
		logger.Debug("Target platform is currently not supported. Skipping\n")
		return nil
	}
	if flag.Lookup("test.v") != nil {
		logger.Debug("Detected test. Skipping file existence test")
	} else {
		_, err := os.Stat(filepath.Join(b.outputFolder, fmt.Sprintf("%s.movements.gz", nodeGroup.Prefix)))
		if os.IsNotExist(err) {
			return err
		}
	}
	command := []string{
		model,
		"-f",
		nodeGroup.Prefix,
	}
	err := b.execute(command)
	return err
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

	stepFilePath := filepath.Join(b.outputFolder, BonnMotionStepFile)
	allowedToWriteStepFile := folderstructure.MayCreatePath(stepFilePath)
	if !allowedToWriteStepFile {
		logger.Error("Not allowed to write step file!")
		return
	}
	b.stepFile, err = os.OpenFile(stepFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logger.Error("Error opening step file:", err)
		return
	}
	defer func() {
		if cerr := b.stepFile.Close(); cerr != nil {
			logger.Error("Error closing step file:", cerr)
		}
	}()

	for _, nodeGroup := range exp.NodeGroups {
		logger.Tracef("Processing NodeGroup \"%s\"", nodeGroup.Prefix)
		if !b.MovementPatternIsSupported(nodeGroup.MovementModel) {
			logger.Debugf("Movement model \"%s\" is currently not supported. Skipping", nodeGroup.MovementModel.String())
			continue
		}
		if !folderstructure.MayCreatePath(filepath.Join(b.outputFolder, fmt.Sprintf("%s.movements.gz", nodeGroup.Prefix))) {
			logger.Error("Not allowed to write output file!")
			return
		}
		err = b.generateNodeGroup(exp, nodeGroup)
		if err != nil {
			logger.Error("Could not generate NodeGroup movement:", err)
			continue
		}
		for _, target := range exp.Targets {
			err = b.convertToTargetFormat(target, nodeGroup)
			if err != nil {
				logger.Error("Error converting to target format:", err)
			}
		}
	}
	logger.Trace("Finished generation")
}
