package main

import (
	"fmt"
	"os"

	"github.com/stg-tud/bp2022_netlab/cmd"
)

func main() {
	err := cmd.Execute()
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
}
