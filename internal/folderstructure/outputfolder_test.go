package folderstructure_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
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

		if outputFolder != expectedOutput {
			t.Fatalf("Output \"%s\" does not match expected output \"%s\"!", outputFolder, expectedOutput)
		}
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
	if outputFolder != expectedOutput {
		t.Fatalf("Output \"%s\" does not match expected output \"%s\"!", outputFolder, expectedOutput)
	}
}

func TestGetAndCreateOutputFolder(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})
	if _, err := os.Stat(folderstructure.OutputFolderName); !os.IsNotExist(err) {
		// Folder already exists. Removing it in order to check generation.
		err = os.RemoveAll(folderstructure.OutputFolderName)
		if err != nil {
			t.Fatal("Could not remove existing output folder!")
		}
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
		if err != nil {
			t.Fatal("Error running GetAndCreateOutputFolder:", err)
		}
		if outputFolder != expectedOutput {
			t.Fatalf("Output \"%s\" does not match expected output \"%s\"!", outputFolder, expectedOutput)
		}
		if _, err := os.Stat(outputFolder); os.IsNotExist(err) {
			t.Fatal("Method did not create folder!")
		}
	}
}
