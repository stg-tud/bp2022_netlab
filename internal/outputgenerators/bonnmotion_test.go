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

func TestBonnmotionGeneration(t *testing.T) {
	t.Cleanup(func() {
		os.RemoveAll(folderstructure.OutputFolderName)
	})

	outputgenerators.BonnMotionExecutable = "bonnmotion"

	og := outputgenerators.Bonnmotion{}
	testingExperiment := GetTestingExperiment()
	outputFolder := folderstructure.GetOutputFolder(testingExperiment, "movements")
	og.Generate(testingExperiment)

	expected, err := os.ReadFile(filepath.Join(TestDataFolder, outputgenerators.BonnMotionStepFile))
	assert.NoError(t, err)
	expectedClean := strings.ReplaceAll(string(expected), "\r\n", "\n")

	actual, err := os.ReadFile(filepath.Join(outputFolder, outputgenerators.BonnMotionStepFile))
	assert.NoError(t, err)
	actualClean := strings.ReplaceAll(string(actual), "\r\n", "\n")

	assert.Equal(t, expectedClean, actualClean)
}
