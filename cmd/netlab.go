package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "netlab",
	Short: "netlab is a tool to generate network experiment configuration files",
	Long: `            _   _       _     
 _ __   ___| |_| | __ _| |__  
| '_ \ / _ \ __| |/ _' | '_ \ 
| | | |  __/ |_| | (_| | |_) |
|_| |_|\___|\__|_|\__,_|_.__/  v0.0.1
                              

netlab helps you to quickly generate configuration files for network experiment
softwares such as Core, Coreemu-Lab and The ONE. Its aim is to make your work
easier, thus it handles annoying tasks such as generating IP addresses, movement
patterns and multiple parameterized runs.

netlab was developed as part of a bachelor internship at TU Darmstadt. For more
information and documentation visit https://github.com/stg-stud/bp2022_netlab.`,
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number",
	Long:  `Prints the current version of netlab and exits`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("netlab v0.0.1")
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests the given TOML file for errors without generating files",
	Long: `test will run some tests on the given TOML file. It will try to
parse the file and echo syntactical errors as well as invalid values. No
further steps are taken, e.g. no output files will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("netlab v0.0.1")
	},
}

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Loads the given TOML file and generates output files",
	Long: `test will run some tests on the given TOML file. It will try to
parse the file and echo syntactical errors as well as invalid values. No
further steps are taken, e.g. no output files will be generated.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("netlab v0.0.1")
	},
}

func init() {
	rootCmd.PersistentFlags().Bool("debug", true, "enable debug logging")
	rootCmd.AddCommand(versionCmd)
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(testCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
