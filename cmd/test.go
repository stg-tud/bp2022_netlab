package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:   "test [filename]",
	Short: "Tests the given TOML file for errors without generating files",
	Long: `test will run some tests on the given TOML file. It will try to
parse the file and echo syntactical errors as well as invalid values. No
further steps are taken, e.g. no output files will be generated.`,
	Args: cobra.MatchAll(cobra.ExactArgs(1), cobra.OnlyValidArgs),
	Run:  test,
}

func init() {
}

func test(cmd *cobra.Command, args []string) {
	fmt.Println("Not yet implemented")
}
