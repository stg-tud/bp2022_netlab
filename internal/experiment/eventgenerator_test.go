package experiment_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stretchr/testify/assert"
)

func TestEventGeneratorWithoutName(t *testing.T) {
	_, err := experiment.NewDefaultEventGenerator("")
	assert.Error(t, err)
	assert.Equal(t, errors.New("name of the EventGenerator must consist of at least on character"), err)
	_, err = experiment.NewEventGenerator("", eventgeneratortypes.MessageEventGenerator{})
	assert.Error(t, err)
	assert.Equal(t, errors.New("name of the EventGenerator must consist of at least on character"), err)
}

func TestNewEventGenerator(t *testing.T) {
	eventGenerator := []eventgeneratortypes.EventGeneratorType{
		eventgeneratortypes.MessageBurstGenerator{}.Default(),
		eventgeneratortypes.MessageEventGenerator{}.Default(),
	}
	for _, eventGenerator := range eventGenerator {
		eventGeneratorName := fmt.Sprintf("eventGenerator_under_test_%s", strings.ToLower(eventGenerator.String()))
		eventGeneratorUnderTest, err := experiment.NewEventGenerator(eventGeneratorName, eventGenerator)
		assert.NoError(t, err)
		assert.Equal(t, eventGeneratorName, eventGeneratorUnderTest.Name)
		assert.Equal(t, eventGenerator, eventGeneratorUnderTest.Type)
	}
}
