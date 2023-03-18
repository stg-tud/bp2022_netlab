package tests

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/cmd"
	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stretchr/testify/assert"
)

func TestMinialFile(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	assert.FileExists(t, filepath.Join(ExamplesFolder, "minimal.toml"))

	netlabCmd := cmd.RootCmd
	netlabCmd.SetArgs([]string{"generate", "-d", filepath.Join(ExamplesFolder, "minimal.toml")})
	err := netlabCmd.Execute()
	assert.NoError(t, err)

	assert.FileExists(t, filepath.Join(folderstructure.OutputFolderName, "netlab.log"))
	assert.FileExists(t, filepath.Join(folderstructure.OutputFolderName, "A_very_simple_experiment", "cluster_settings.txt"))
	assert.FileExists(t, filepath.Join(folderstructure.OutputFolderName, "A_very_simple_experiment", "core.xml"))
	assert.FileExists(t, filepath.Join(folderstructure.OutputFolderName, "A_very_simple_experiment", "debug_out.toml"))
	assert.FileExists(t, filepath.Join(folderstructure.OutputFolderName, "A_very_simple_experiment", "experiment.conf"))

	compareFiles := []string{
		"cluster_settings.txt",
		"core.xml",
		"debug_out.toml",
		"experiment.conf",
	}

	for _, fileName := range compareFiles {
		expectedFile, err := os.ReadFile(filepath.Join(TestDataFolder, "minimal", fileName))
		assert.NoError(t, err)
		expectedClean := strings.ReplaceAll(string(expectedFile), "\r\n", "\n")

		actualFile, err := os.ReadFile(filepath.Join(folderstructure.OutputFolderName, "A_very_simple_experiment", fileName))
		assert.NoError(t, err)
		actualClean := strings.ReplaceAll(string(actualFile), "\r\n", "\n")

		assert.Equal(t, expectedClean, actualClean)
	}
}
