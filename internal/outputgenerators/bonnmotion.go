package outputgenerators

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// The executable of BonnMotion to call.
const EXECUTABLE = "bonnmotion"

// The Bonnmotion output generator calles BonnMotion with the correct parameters.
type Bonnmotion struct{}

// Returns the correct BonnMotion platform name for the given Target.
func platform(t experiment.Target) (bool, string) {
	switch t {
	case experiment.TARGET_THEONE:
		return true, "TheONEFile"

	case experiment.TARGET_CORE:
		return true, "NSFile"

	default:
		return false, ""
	}
}

// Returns whether the given Target is (currently) supported by this output generator.
func isSupported(t experiment.Target) bool {
	supported, _ := platform(t)
	return supported
}

// Returns the parameter set for Random Waypoint movement model for a given NodeGroup inside an Experiment.
func randomWaypointParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	movementmodel := nodeGroup.MovementModel.(movementpatterns.RandomWaypoint)
	return []string{
		"RandomWaypoint",
		fmt.Sprintf("-h%d", movementmodel.MaxSpeed),
		fmt.Sprintf("-l%d", movementmodel.MinSpeed),
		fmt.Sprintf("-p%d", movementmodel.MaxPause),
	}
}

// Returns the general parameters for a given NodeGroup inside an Experiment.
func generalParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	return []string{
		fmt.Sprintf("-d%d", exp.Duration),
		fmt.Sprintf("-n%d", nodeGroup.NoNodes),
		fmt.Sprintf("-x%d", exp.WorldSize.Width),
		fmt.Sprintf("-y%d", exp.WorldSize.Height),
	}
}

// Calls BonnMotion to generate the Random Waypoint data for a given NodeGroup inside an Experiment.
func generateRandomWaypointNodeGroup(exp experiment.Experiment, nodeGroup experiment.NodeGroup) {
	command := []string{
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	command = append(command, randomWaypointParameters(exp, nodeGroup)...)
	command = append(command, generalParameters(exp, nodeGroup)...)
	fmt.Printf("Random Waypoint. Running: %v\n", command)
	execCommand := exec.Command(EXECUTABLE, command...)
	execCommand.Dir = OUTPUT_FOLDER
	_, err := execCommand.Output()
	if err != nil {
		panic(err)
	}
}

// Calls BonnMotion to convert the BonnMotion output to the given Target's format for a given NodeGroup.
func convertToTargetFormat(target experiment.Target, nodeGroup experiment.NodeGroup) {
	supported, model := platform(target)
	if !supported {
		fmt.Printf("Target platform \"%s\" is currently not supported.\n", target.String())
		return
	}
	command := []string{
		model,
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	execCommand := exec.Command(EXECUTABLE, command...)
	execCommand.Dir = OUTPUT_FOLDER
	_, err := execCommand.Output()
	if err != nil {
		panic(err)
	}
}

// Generate generates output for the given Experiment with BonnMotion.
func (t Bonnmotion) Generate(exp experiment.Experiment) {
	os.Mkdir(OUTPUT_FOLDER, 0755)
	for i := 0; i < len(exp.NodeGroups); i++ {
		nodeGroup := exp.NodeGroups[i]
		switch nodeGroup.MovementModel.(type) {
		case movementpatterns.RandomWaypoint:
			generateRandomWaypointNodeGroup(exp, nodeGroup)
		default:
			fmt.Printf("Movement model \"%s\" is currently not supported.\n", reflect.TypeOf(nodeGroup.MovementModel))
			continue
		}
		for y := 0; y < len(exp.Targets); y++ {
			convertToTargetFormat(exp.Targets[y], nodeGroup)
		}
	}
}
