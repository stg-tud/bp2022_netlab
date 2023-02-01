package experiment_test

import (
	
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestLoad(t *testing.T) {

	exp := experiment.LoadFromFile("testdata/format.toml")
	
	
	if exp.NodeGroups[0].IPv4Net != "10.0.0.0" {
		t.Error()
	}

	if exp.NodeGroups[3].IPv6Net != "2001::" {
		t.Error()
	}

	if exp.NodeGroups[3].Bandwidth != 54000000 {
		t.Error()
	}
}
