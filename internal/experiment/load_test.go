package experiment_test

import (
	
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestLoad(t *testing.T) {

	exp := experiment.LoadFromFile("testdata/format.toml")
	exp.Duration =1
	

}
