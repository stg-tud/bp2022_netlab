package main

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestLoad(t *testing.T) {
	name := "WIRELESS_LAN"
	exp := experiment.Loading("format.toml")

	if exp.NodeGroups[0].NetworkType != name {
		t.Error()
	}

}
