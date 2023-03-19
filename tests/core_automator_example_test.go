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

func TestCoreAutomatorExampleFile(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	assert.FileExists(t, filepath.Join(ExamplesFolder, "core_automator_example.toml"))

	netlabCmd := cmd.RootCmd
	netlabCmd.SetArgs([]string{"generate", "-d", filepath.Join(ExamplesFolder, "core_automator_example.toml")})
	err := netlabCmd.Execute()
	assert.NoError(t, err)

	compareFiles := []string{
		"core.xml",
		"experiment.conf",
		"debug_out.toml",
	}

	for _, fileName := range compareFiles {
		expectedFile, err := os.ReadFile(filepath.Join(TestDataFolder, "core_automator_example", fileName))
		assert.NoError(t, err)
		expectedClean := strings.ReplaceAll(string(expectedFile), "\r\n", "\n")

		actualFile, err := os.ReadFile(filepath.Join(folderstructure.OutputFolderName, "core-automator", fileName))
		assert.NoError(t, err)
		actualClean := strings.ReplaceAll(string(actualFile), "\r\n", "\n")

		assert.Equal(t, expectedClean, actualClean)
	}
}
