package outputgenerators

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

const EXECUTABLE = "bonnmotion"

type Bonnmotion struct{}

type commandParams struct{}

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

func isSupported(t experiment.Target) bool {
	supported, _ := platform(t)
	return supported
}

func randomWaypointParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	movementmodel := nodeGroup.MovementModel.(movementpatterns.RandomWaypoint)
	return []string{
		"RandomWaypoint",
		fmt.Sprintf("-h%d", movementmodel.MaxSpeed),
		fmt.Sprintf("-l%d", movementmodel.MinSpeed),
		fmt.Sprintf("-p%d", movementmodel.MaxPause),
	}
}

func generalParameters(exp experiment.Experiment, nodeGroup experiment.NodeGroup) []string {
	return []string{
		fmt.Sprintf("-d%d", exp.Duration),
		fmt.Sprintf("-n%d", nodeGroup.NoNodes),
		fmt.Sprintf("-x%d", exp.WorldSize.Width),
		fmt.Sprintf("-y%d", exp.WorldSize.Height),
	}
}

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
