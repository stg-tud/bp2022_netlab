package parser_test

import (
	"errors"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/networktypes"
	"github.com/stg-tud/bp2022_netlab/internal/parser"
	"github.com/stretchr/testify/assert"
)

func TestMissingRequiredValuesText(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Duration\": field is required"), err)
}

func TestInvalidValueIntToUint(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = -1
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Duration\": invalid input data type -- expected uint"), err)
}

func TestInvalidValueFloatToUint(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1.0
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Duration\": invalid input data type -- expected uint"), err)
}

func TestInvalidValueStringToUint(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = "1"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Duration\": invalid input data type -- expected uint"), err)
}

func TestInvalidValueBoolToUint(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = true
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Duration\": invalid input data type -- expected uint"), err)
}

func TestInvalidValueFloatToInt(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1

	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Range = 1.0
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Range\": invalid input data type -- expected int"), err)
}

func TestInvalidValueStringToInt(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1

	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Range = "1"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Range\": invalid input data type -- expected int"), err)
}

func TestInvalidValueBoolToInt(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Range = false
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Range\": invalid input data type -- expected int"), err)
}

func TestValidValueIntToFloat(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Loss = 3
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)
	assert.EqualValues(t, 3.0, exp.Networks[0].Type.(networktypes.WirelessLAN).Loss)
}

func TestInvalidValueStringToFloat(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Loss = "3.0"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Loss\": invalid input data type -- expected float32"), err)
}

func TestInvalidValueBoolToFloat(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Loss = false
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Loss\": invalid input data type -- expected float32"), err)
}

func TestInvalidValueIntToString(t *testing.T) {
	toml := `
	Name = 1
	Duration = 1
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Name\": invalid input data type -- expected string"), err)
}

func TestInvalidValueBoolToString(t *testing.T) {
	toml := `
	Name = false
	Duration = 1
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing field \"Name\": invalid input data type -- expected string"), err)
}

func TestValidValueIntToBool(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Promiscuous = 1
	`

	exp, err := parser.ParseText([]byte(toml))

	assert.NoError(t, err)
	assert.True(t, exp.Networks[0].Type.(networktypes.WirelessLAN).Promiscuous)
}

func TestInvalidValueIntToBool(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Promiscuous = 2
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Promiscuous\": invalid input data type -- expected bool"), err)
}

func TestInvalidValueFloatToBool(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Promiscuous = 1.0
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Promiscuous\": invalid input data type -- expected bool"), err)
}

func TestInvalidValueStringToBool(t *testing.T) {
	toml := `
	Name = "Testing Experiment"
	Duration = 1
	
	[[Network]]
	Name = "Wifi"
	Type = "wireless lan"
	Promiscuous = "true"
	`

	_, err := parser.ParseText([]byte(toml))

	assert.Error(t, err)
	assert.Equal(t, errors.New("error parsing network 0: error parsing field \"Promiscuous\": invalid input data type -- expected bool"), err)
}
