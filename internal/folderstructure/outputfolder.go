// Package folderstructure provides functions for the folder strucure of output data
package folderstructure

import (
	"os"
	"path/filepath"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// The name of the (top level) output folder
const OutputFolderName = "output"

func GetOutputFolder(exp experiment.Experiment, subfolders ...string) string {
	outputFolder := filepath.Join(OutputFolderName, FileSystemEscape(exp.Name))
	for _, subfolder := range subfolders {
		outputFolder = filepath.Join(outputFolder, FileSystemEscape(subfolder))
	}
	return outputFolder
}

func GetAndCreateOutputFolder(exp experiment.Experiment, subfolders ...string) (string, error) {
	outputFolder := GetOutputFolder(exp, subfolders...)
	os.MkdirAll(outputFolder, 0755)
	return outputFolder, nil
}
