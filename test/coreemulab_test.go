package main

import (
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/outputgenerators"
)

func TestOutput(t *testing.T) {

	core := outputgenerators.CoreEmulab{}
	core.Generate()
}
