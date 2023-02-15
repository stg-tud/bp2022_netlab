package cmd

import (
	"errors"
	"strings"

	logger "github.com/gookit/slog"

	"github.com/spf13/cobra"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/logging"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

var overwriteExisting bool
var targetsOverwrite []string
var outputFolder string

var generateCmd = &cobra.Command{
	Use:     "generate [filename]",
	Aliases: []string{"gen"},
	Short:   "Loads the given TOML file and generates output files",
	Long: `generate will load the given TOML file and generate outputs
for all targets specified in it. If targets are specified. as flags,
the targets from the file will be ignored.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:  generate,
}

func init() {
	generateCmd.PersistentFlags().BoolVarP(&overwriteExisting, "overwrite", "o", false, "overwrite existing files for the same configuration")
	generateCmd.PersistentFlags().StringArrayVarP(&targetsOverwrite, "targets", "t", []string{}, "generate configs for the following targets (no matter which targets are configured in the TOML file)")
	generateCmd.PersistentFlags().StringVarP(&outputFolder, "folder", "f", "output", "name of the folder the output should be written to (default: output)")
}

func stringTargetMapping(input string) (experiment.Target, error) {
	cleanedInput := strings.ToLower(input)
	cleanedInput = strings.TrimSpace(cleanedInput)
	switch cleanedInput {
	case "the-one", "theone", "one":
		return experiment.TargetTheOne, nil

	case "core", "coreemu":
		return experiment.TargetCore, nil

	case "coreemulab", "core-emulab", "coreemu-lab", "clab":
		return experiment.TargetCoreEmulab, nil

	default:
		return 0, errors.New("no matching target found")
	}
}

func buildTargets(exp experiment.Experiment) []experiment.Target {
	targets := exp.Targets
	if len(targetsOverwrite) > 0 {
		targets = []experiment.Target{}
		for _, targetString := range targetsOverwrite {
			target, err := stringTargetMapping(targetString)
			if err != nil {
				logger.Warnf("Error parsing target \"%s\": %s", targetString, err)
				continue
			}
			targets = append(targets, target)
		}
	}
	logger.Debug("Generating output for following targets:", targets)
	return targets
}

func targetOutputGeneratorMapping(input experiment.Target) (outputgenerators.OutputGenerator, error) {
	switch input {
	case experiment.TargetCore:
		return outputgenerators.Core{}, nil

	case experiment.TargetCoreEmulab:
		return outputgenerators.CoreEmulab{}, nil

	default:
		return outputgenerators.Debug{}, errors.New("no matching output generator found")
	}
}

func buildOutputGenerators(exp experiment.Experiment) []outputgenerators.OutputGenerator {
	outputGenerators := []outputgenerators.OutputGenerator{}
	for _, target := range exp.Targets {
		outputGenerator, err := targetOutputGeneratorMapping(target)
		if err != nil {
			continue
		}
		outputGenerators = append(outputGenerators, outputGenerator)
		if target == experiment.TargetCoreEmulab {
			logger.Debug("Output generator \"coreemu-lab\" implies \"CORE\", therefore adding it")
			outputGenerators = append(outputGenerators, outputgenerators.Core{})
		}
	}
	for _, nodeGroup := range exp.NodeGroups {
		supported := outputgenerators.Bonnmotion{}.MovementPatternIsSupported(nodeGroup.MovementModel)
		if supported {
			logger.Debugf("Selecting BonnMotion output generator because node group \"%s\" needs a movement model", nodeGroup.Prefix)
			outputGenerators = append(outputGenerators, outputgenerators.Bonnmotion{})
			break
		}
	}
	if debug {
		logger.Debugf("Selecting Debug output generator because debug logging is active")
		outputGenerators = append(outputGenerators, outputgenerators.Debug{})
	}

	// Making sure no OutputGenerator appears more than once
	unique := make(map[outputgenerators.OutputGenerator]bool)
	cleanedOutputGenerators := []outputgenerators.OutputGenerator{}
	for _, outputGenerator := range outputGenerators {
		if _, present := unique[outputGenerator]; !present {
			unique[outputGenerator] = true
			cleanedOutputGenerators = append(cleanedOutputGenerators, outputGenerator)
		}
	}

	logger.Debug("Generating output for following output generators:", cleanedOutputGenerators)
	return cleanedOutputGenerators
}

func generate(cmd *cobra.Command, args []string) {
	folderstructure.OutputFolderName = outputFolder
	folderstructure.OverwriteExisting = overwriteExisting

	logging.Init(debug)
	logger.Info("Starting")

	exp, err := experiment.LoadFromFile(args[0])
	if err != nil {
		logger.Error("Error loading file:", err)
		return
	}

	logger.Info("Using random seed", exp.RandomSeed)

	exp.Targets = buildTargets(exp)
	outputGenerators := buildOutputGenerators(exp)

	for _, outputGenerator := range outputGenerators {
		outputGenerator.Generate(exp)
	}

	logger.Info("Finished")
}
