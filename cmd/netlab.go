package cmd

import (
	"github.com/spf13/cobra"
)

var debug bool

var rootCmd = &cobra.Command{
	Use:     "netlab",
	Short:   "netlab is a tool to generate network experiment configuration files",
	Version: "0.0.2",
	Long: `            _   _       _     
 _ __   ___| |_| | __ _| |__  
| '_ \ / _ \ __| |/ _' | '_ \ 
| | | |  __/ |_| | (_| | |_) |
|_| |_|\___|\__|_|\__,_|_.__/
                              

netlab helps you to quickly generate configuration files for network experiment
softwares such as Core, Coreemu-Lab and The ONE. Its aim is to make your work
easier, thus it handles annoying tasks such as generating IP addresses, movement
patterns and multiple parameterized runs.

netlab was developed as part of a bachelor internship at TU Darmstadt. For more
information and documentation visit https://github.com/stg-stud/bp2022_netlab.`,
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "enable debug logging")
	rootCmd.AddCommand(generateCmd)
	rootCmd.AddCommand(testCmd)
}

func Execute() error {
	return rootCmd.Execute()
}
