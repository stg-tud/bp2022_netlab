package main

import (
	"os"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestLoad(t *testing.T) {

	doc, err := os.ReadFile("format.toml")
	if err != nil {
		panic(err)
	}
	name := "core-automator"
	exp := experiment.Loading(doc)

	if exp.Name != name {
		t.Error()
	}
	
}
