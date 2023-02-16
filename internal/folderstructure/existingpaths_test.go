package folderstructure_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
)

func TestExistingPathsAllowance(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	folderstructure.OverwriteExisting = true

	err := os.MkdirAll(folderstructure.OutputFolderName, 0755)
	if err != nil {
		t.Fatal("Error creating output folder", err)
	}
	testingFile := filepath.Join(folderstructure.OutputFolderName, "test.test")
	if !folderstructure.MayCreatePath(testingFile) {
		t.Fatal("Creation denied while it should be allowed!")
	}
	fbuffer, err := os.Create(testingFile)
	if err != nil {
		t.Fatal("Error creating file", err)
	}
	defer fbuffer.Close()
	if !folderstructure.MayCreatePath(testingFile) {
		t.Fatal("Creation denied while it should be allowed!")
	}
}

func TestExistingPathsDenial(t *testing.T) {
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

	folderstructure.OverwriteExisting = false

	err := os.MkdirAll(folderstructure.OutputFolderName, 0755)
	if err != nil {
		t.Fatal("Error creating output folder", err)
	}
	testingFile := filepath.Join(folderstructure.OutputFolderName, "test.test")
	if !folderstructure.MayCreatePath(testingFile) {
		t.Fatal("Creation denied while it should be allowed!")
	}
	fbuffer, err := os.Create(testingFile)
	if err != nil {
		t.Fatal("Error creating file", err)
	}
	defer fbuffer.Close()
	if folderstructure.MayCreatePath(testingFile) {
		t.Fatal("Creation allowed while it should be denied!")
	}
}
