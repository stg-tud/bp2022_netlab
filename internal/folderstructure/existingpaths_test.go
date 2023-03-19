package folderstructure_test

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stretchr/testify/assert"
)

func TestExistingPathsAllowance(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	folderstructure.OverwriteExisting = true

	err := os.MkdirAll(folderstructure.OutputFolderName, 0755)
	assert.NoError(t, err)

	testingFile := filepath.Join(folderstructure.OutputFolderName, "test.test")
	assert.True(t, folderstructure.MayCreatePath(testingFile), "Creation denied while it should be allowed!")

	fbuffer, err := os.Create(testingFile)
	assert.NoError(t, err)
	defer fbuffer.Close()

	assert.True(t, folderstructure.MayCreatePath(testingFile), "Creation denied while it should be allowed!")
}

func TestExistingPathsDenial(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})
	if _, err := os.Stat(folderstructure.OutputFolderName); !os.IsNotExist(err) {
		// Folder already exists. Removing it in order to check generation.
		err = os.RemoveAll(folderstructure.OutputFolderName)
		assert.NoError(t, err)
	}

	folderstructure.OverwriteExisting = false

	err := os.MkdirAll(folderstructure.OutputFolderName, 0755)
	assert.NoError(t, err)
	testingFile := filepath.Join(folderstructure.OutputFolderName, "test.test")
	assert.True(t, folderstructure.MayCreatePath(testingFile), "Creation denied while it should be allowed!")

	fbuffer, err := os.Create(testingFile)
	assert.NoError(t, err)
	defer fbuffer.Close()

	assert.False(t, folderstructure.MayCreatePath(testingFile), "Creation allowed while it should be denied!")
}
