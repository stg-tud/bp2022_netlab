package folderstructure_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/folderstructure"
	"github.com/stretchr/testify/assert"
)

func TestFilesystemEncode(t *testing.T) {
	tests := map[string]string{
		"nothingToEncode":                "nothingToEncode",
		"Only-minus-and-digits-123":      "Only-minus-and-digits-123",
		"A löt of §bad$ chars":           "A_l_t_of__bad__chars",
		"Slashes/Should/Also/Be/Escaped": "Slashes_Should_Also_Be_Escaped",
	}
	for input, expectedOutput := range tests {
		output := folderstructure.FileSystemEscape(input)
		assert.Equal(t, expectedOutput, output)
	}
}
