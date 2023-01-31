// Package folderstructure provides functions for the folder strucure of output data
package folderstructure

import (
	"os"
	"path/filepath"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

// The name of the (top level) output folder
const OutputFolderName = "output"

// GetOutputFolder returns the relative path of the output folder
func GetOutputFolder(exp experiment.Experiment, subfolders ...string) string {
	outputFolder := filepath.Join(OutputFolderName, FileSystemEscape(exp.Name))
	for _, subfolder := range subfolders {
		outputFolder = filepath.Join(outputFolder, FileSystemEscape(subfolder))
	}
	return outputFolder
}

// GetOutputFolder returns the relative path of the output folder and creates the folder if it does not exist
func GetAndCreateOutputFolder(exp experiment.Experiment, subfolders ...string) (string, error) {
	outputFolder := GetOutputFolder(exp, subfolders...)
	err := os.MkdirAll(outputFolder, 0755)
	if err != nil {
		return "", err
	}
	return outputFolder, nil
}
