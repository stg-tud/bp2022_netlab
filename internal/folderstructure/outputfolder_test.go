package folderstructure_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stretchr/testify/assert"
)

func TestGetOutputFolder(t *testing.T) {
	experiments := []experiment.Experiment{
		experiment.GetExampleExperiment(),
		{
			Name:       "This is a Test?",
			RandomSeed: experiment.GenerateRandomSeed(),
		},
	}
	for _, experiment := range experiments {
		outputFolder := folderstructure.GetOutputFolder(experiment)
		expectedOutput := filepath.Join(folderstructure.OutputFolderName, folderstructure.FileSystemEscape(experiment.Name))

		assert.Equal(t, expectedOutput, outputFolder)
	}
}

func TestGetOutputFolderWithSubfolders(t *testing.T) {
	exampleExperiment := experiment.GetExampleExperiment()
	subfolders := []string{
		"nothingToEncode",
		"Only-minus-and-digits-123",
		"A löt of §bad$ chars",
	}
	expectedSubfolders := []string{
		"nothingToEncode",
		"Only-minus-and-digits-123",
		"A_l_t_of__bad__chars",
	}
	outputFolder := folderstructure.GetOutputFolder(exampleExperiment, subfolders...)
	expectedOutput := filepath.Join(folderstructure.OutputFolderName, "Example_Experiment", expectedSubfolders[0], expectedSubfolders[1], expectedSubfolders[2])
	assert.Equal(t, expectedOutput, outputFolder)
}

func TestGetAndCreateOutputFolder(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})
	if _, err := os.Stat(folderstructure.OutputFolderName); !os.IsNotExist(err) {
		// Folder already exists. Removing it in order to check generation.
		err = os.RemoveAll(folderstructure.OutputFolderName)
		assert.NoError(t, err)
	}

	experiments := []experiment.Experiment{
		experiment.GetExampleExperiment(),
		{
			Name:       "This is a Test?",
			RandomSeed: experiment.GenerateRandomSeed(),
		},
	}
	// Add the second experiment a second time to check for errors with existing folders
	experiments = append(experiments, experiments[1])
	for _, experiment := range experiments {
		subfolders := []string{"test-subdir"}
		outputFolder, err := folderstructure.GetAndCreateOutputFolder(experiment, subfolders...)
		expectedOutput := filepath.Join(folderstructure.OutputFolderName, folderstructure.FileSystemEscape(experiment.Name), "test-subdir")
		assert.NoError(t, err)
		assert.Equal(t, expectedOutput, outputFolder)
		assert.DirExists(t, outputFolder)
	}
}
