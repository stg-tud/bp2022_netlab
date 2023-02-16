package cmd

import (
	"fmt"

	logger "github.com/gookit/slog"
	"github.com/spf13/cobra"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/logging"
)

var testCmd = &cobra.Command{
	Use:   "test [filename]",
	Short: "Tests the given TOML file for errors without generating files",
	Long: `test will run some tests on the given TOML file. It will try to
parse the file and echo syntactical errors as well as invalid values. No
further steps are taken, e.g. no output files will be generated.`,
	Args:         cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	SilenceUsage: true,
	Run:          test,
}

// test user's input file for errors
func test(cmd *cobra.Command, args []string) {
	logging.Init(debug)
	_, err := experiment.LoadFromFile(args[0])
	if err != nil {
		logger.Error("Error parsing file:", err)
		fmt.Println()
		fmt.Println("File has errors. See above for more details.")
		if !debug {
			fmt.Println("Enable debug output with --debug to get further insights.")
		}
		return
	}
	fmt.Println("File is OK")
}
