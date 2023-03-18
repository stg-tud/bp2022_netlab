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

func TestMultipleNetworksAndGroupsFile(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	assert.FileExists(t, filepath.Join(ExamplesFolder, "multiple_networks_and_groups.toml"))

	netlabCmd := cmd.RootCmd
	netlabCmd.SetArgs([]string{"generate", "-d", filepath.Join(ExamplesFolder, "multiple_networks_and_groups.toml")})
	err := netlabCmd.Execute()
	assert.NoError(t, err)

	compareFiles := []string{
		"cluster_settings.txt",
		"core.xml",
		"debug_out.toml",
		"experiment.conf",
		filepath.Join("movements", "bonnmotion.steps"),
	}

	for _, fileName := range compareFiles {
		expectedFile, err := os.ReadFile(filepath.Join(TestDataFolder, "multi_networks_groups", fileName))
		assert.NoError(t, err)
		expectedClean := strings.ReplaceAll(string(expectedFile), "\r\n", "\n")

		actualFile, err := os.ReadFile(filepath.Join(folderstructure.OutputFolderName, "Multiple_Networks___Groups", fileName))
		assert.NoError(t, err)
		actualClean := strings.ReplaceAll(string(actualFile), "\r\n", "\n")

		assert.Equal(t, expectedClean, actualClean)
	}
}
