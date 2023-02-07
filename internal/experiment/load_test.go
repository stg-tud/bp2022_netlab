package experiment_test

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestLoad(t *testing.T) {

	exp := experiment.LoadFromFile("testdata/example.toml")

	if (exp.Duration!=132){
		t.Error()
	}
	if (exp.Networks[0].Name!="wireless_lan"){
		t.Error()
	}
	
	if (exp.Networks[2].Name!="changed_wifi"){
		t.Error()
	}
}
