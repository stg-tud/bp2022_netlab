package experiment_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestLoad(t *testing.T) {

	exp := experiment.LoadFromFile("testdata/example.toml")
	exp.Duration = 1
	t.Error()
}
