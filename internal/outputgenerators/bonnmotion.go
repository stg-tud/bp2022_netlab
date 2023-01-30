package outputgenerators

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"text/template"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/movementpatterns"
)

// The name of the executable to run BonnMotion.
const BonnMotionExecutable = "bonnmotion"

// The name of the file that the taken steps should be written into.
const BonnMotionStepFile = "bonnmotion.steps"

// The Bonnmotion output generator calles BonnMotion with the correct parameters.
type Bonnmotion struct{}

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
func (Bonnmotion) execute(command []string) error {
	stepFile, err := os.OpenFile(filepath.Join(OutputFolder, BonnMotionStepFile), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer stepFile.Close()
	stepFile.WriteString(fmt.Sprintln(command))
	execCommand := exec.Command(BonnMotionExecutable, command...)
	execCommand.Dir = OutputFolder
	// Check if the function is currently unit tested and do not execute actual BonnMotion command if so.
	if flag.Lookup("test.v") != nil {
		return nil
	}
	_, err = execCommand.Output()
	return err
}

// Calls BonnMotion to generate the Random Waypoint data for a given NodeGroup inside an Experiment.
func (b Bonnmotion) generateRandomWaypointNodeGroup(exp experiment.Experiment, nodeGroup experiment.NodeGroup) {
	command := []string{
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	command = append(command, b.randomWaypointParameters(exp, nodeGroup)...)
	command = append(command, b.generalParameters(exp, nodeGroup)...)
	err := b.execute(command)
	if err != nil {
		panic(err)
	}
}

// Writes the .ns_params file containing the relevant Experiment information
func (Bonnmotion) writeNSParamsFile(exp experiment.Experiment) error {
	paramsOutBuffer, err := os.Create(filepath.Join(OutputFolder, "core.ns_params"))
	if err != nil {
		return err
	}
	defer paramsOutBuffer.Close()
	paramsFile, err := template.ParseFiles(filepath.Join(GetTemplatesFolder(), "core.ns_params"))
	if err != nil {
		return err
	}

	var noNodes uint = 0
	for _, nodeGroup := range exp.NodeGroups {
		noNodes += nodeGroup.NoNodes
	}

	type paramsData struct {
		Experiment experiment.Experiment
		NoNodes    uint
	}

	fillData := paramsData{
		Experiment: exp,
		NoNodes:    noNodes,
	}

	paramsFile.Execute(paramsOutBuffer, fillData)
	return nil
}

func (Bonnmotion) parseNSMovementLine3(line string) (idInFile uint, posX float64, posY float64, error error) {
	parts := strings.Split(line, " ")
	if len(parts[0]) < 8 || parts[0][:7] != "$node_(" {
		error = errors.New("line is invalid")
		return
	}
	nodeId := strings.Split(parts[0], "(")[1]
	nodeId = nodeId[:len(nodeId)-1]
	sidInFile, err := strconv.ParseUint(nodeId, 10, 32)
	if err != nil {
		error = err
		return
	}
	idInFile = uint(sidInFile)
	pos, err := strconv.ParseFloat(parts[3], 64)
	if err != nil {
		error = err
		return
	}
	if parts[2] == "X_" {
		posX = pos
	} else if parts[2] == "Y_" {
		posY = pos
	}
	return
}

func (Bonnmotion) parseNSMovementLine2(line string) (timestamp float64, idInFile uint, destX float64, destY float64, destZ float64, error error) {
	parts := strings.Split(line, " ")
	if parts[0] != "$ns_" {
		error = errors.New("line is invalid")
		return
	}
	timestamp, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		error = err
		return
	}
	nodeId := strings.Split(parts[3], "(")[1]
	nodeId = nodeId[:len(nodeId)-1]
	sidInFile, err := strconv.ParseUint(nodeId, 10, 32)
	if err != nil {
		error = err
		return
	}
	idInFile = uint(sidInFile)
	destX, err = strconv.ParseFloat(parts[5], 64)
	if err != nil {
		error = err
		return
	}
	destY, err = strconv.ParseFloat(parts[6], 64)
	if err != nil {
		error = err
		return
	}
	destZ, err = strconv.ParseFloat(parts[7][:len(parts[7])-1], 64)
	if err != nil {
		error = err
		return
	}
	return
}

func (b Bonnmotion) parseNSMovementLine(line string) string {
	type nodePos struct {
		Timestamp float64
		IdInFile  uint
		DestX     float64
		DestY     float64
		DestZ     float64
	}
	timestamp, idInFile, destX, destY, destZ, err := b.parseNSMovementLine2(line)
	orderedList := make(map[float64]nodePos)
	if err == nil {
		_, exists := orderedList[timestamp]
		if exists {
			panic("Timestamp already exists!")
		}
		currentLine := nodePos{
			Timestamp: timestamp,
			IdInFile:  idInFile,
			DestX:     destX,
			DestY:     destY,
			DestZ:     destZ,
		}
		orderedList[timestamp] = currentLine
	}
	// fmt.Println(timestamp, idInFile, destX, destY, destZ, err)
	// idInFile, posX, posY, err := b.parseNSMovementLine3(line)
	return line + "\n"
}

func (b Bonnmotion) parseNSMovementFile(nodeGroup experiment.NodeGroup) error {
	fmt.Println("Processing NodeGroup " + nodeGroup.Prefix)
	nsMovementFile, err := os.Open(filepath.Join(OutputFolder, nodeGroup.Prefix+".ns_movements"))
	if err != nil {
		if os.IsNotExist(err) {
			return errors.New("No movement existing!")
		}
		return err
	}
	defer nsMovementFile.Close()
	nsMovementFileScanner := bufio.NewScanner(nsMovementFile)
	nsMovementFileScanner.Split(bufio.ScanLines)

	type nodePos struct {
		Timestamp float64
		IdInFile  uint
		DestX     float64
		DestY     float64
		DestZ     float64
	}

	orderedList := make(map[float64][]nodePos)
	for nsMovementFileScanner.Scan() {
		timestamp, idInFile, destX, destY, destZ, err := b.parseNSMovementLine2(nsMovementFileScanner.Text())
		if err != nil {
			fmt.Println(err)
			continue
		}
		currentLine := nodePos{
			Timestamp: timestamp,
			IdInFile:  idInFile,
			DestX:     destX,
			DestY:     destY,
			DestZ:     destZ,
		}
		_, exists := orderedList[timestamp]
		if !exists {
			orderedList[timestamp] = []nodePos{}
		}
		orderedList[timestamp] = append(orderedList[timestamp], currentLine)
	}
	fmt.Println(orderedList)
	return nil
}

// Writes the .ns_movements file containing all nodes movements
func (b Bonnmotion) writeNSMovementFile(exp experiment.Experiment) error {
	movementOutBuffer, err := os.Create(filepath.Join(OutputFolder, "core.ns_movements"))
	if err != nil {
		return err
	}
	defer movementOutBuffer.Close()
	var indexOffset uint = ScenarioIdOffset
	for _, nodeGroup := range exp.NodeGroups {
		b.parseNSMovementFile(nodeGroup)
		indexOffset += nodeGroup.NoNodes
	}
	return nil
}

// Merges all generated NSFiles for CORE to one that can be used by CORE.
func (b Bonnmotion) mergeFilesForCore(exp experiment.Experiment) error {
	err := b.writeNSParamsFile(exp)
	if err != nil {
		return err
	}
	err = b.writeNSMovementFile(exp)
	if err != nil {
		return err
	}
	return nil
}

// Calls BonnMotion to convert the BonnMotion output to the given Target's format for a given NodeGroup.
func (b Bonnmotion) convertToTargetFormat(target experiment.Target, nodeGroup experiment.NodeGroup) {
	supported, model := b.platform(target)
	if !supported {
		fmt.Printf("Target platform \"%s\" is currently not supported.\n", target.String())
		return
	}
	command := []string{
		model,
		fmt.Sprintf("-f%s", nodeGroup.Prefix),
	}
	err := b.execute(command)
	if err != nil {
		panic(err)
	}
}

// Generate generates output for the given Experiment with BonnMotion.
func (b Bonnmotion) Generate(exp experiment.Experiment) {
	os.Mkdir(OutputFolder, 0755)
	os.Create(filepath.Join(OutputFolder, BonnMotionStepFile))
	// targetCoreExists := false
	for _, nodeGroup := range exp.NodeGroups {
		switch nodeGroup.MovementModel.(type) {
		case movementpatterns.RandomWaypoint:
			b.generateRandomWaypointNodeGroup(exp, nodeGroup)
		default:
			fmt.Printf("Movement model \"%s\" is currently not supported.\n", reflect.TypeOf(nodeGroup.MovementModel))
			continue
		}
		for _, target := range exp.Targets {
			b.convertToTargetFormat(target, nodeGroup)
			// if target == experiment.TargetCore {
			// 	targetCoreExists = true
			// }
		}
	}
	// if targetCoreExists {
	// 	b.mergeFilesForCore(exp)
	// }
}
