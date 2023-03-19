package outputgenerators_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
	"github.com/stretchr/testify/assert"
)

func TestTheOne(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})
	to := outputgenerators.TheOne{}
	testingExperiment := GetTestingExperiment()
	outputFolder := folderstructure.GetOutputFolder(testingExperiment)
	to.Generate(testingExperiment)

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, outputgenerators.TheOneOutput))
	assert.NoError(t, err)
	expectedClean := strings.ReplaceAll(string(expected), "\r\n", "\n")

	actual, err := os.ReadFile(filepath.Join(outputFolder, outputgenerators.TheOneOutput))
	assert.NoError(t, err)
	actualClean := strings.ReplaceAll(string(actual), "\r\n", "\n")

	assert.Equal(t, expectedClean, actualClean)
}
