package experiment_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/eventgenerators"
	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestEventGeneratorWithoutName(t *testing.T) {
	_, error1 := experiment.NewDefaultEventGenerator("")
	_, error2 := experiment.NewEventGenerator("", eventgenerators.MessageEventGenerator{})
	if error1 == nil || error2 == nil {
		t.Fatal("EventGenerators without name should not be allowed!")
	}
}

func TestNewEventGenerator(t *testing.T) {
	eventGenerator := []eventgenerators.EventGenerator{
		eventgenerators.MessageBurstGenerator{}.Default(),
		eventgenerators.MessageEventGenerator{}.Default(),
	}
	for _, eventGenerator := range eventGenerator {
		eventGeneratorName := fmt.Sprintf("eventGenerator_under_test_%s", strings.ToLower(eventGenerator.String()))
		eventGenerator_under_test, err := experiment.NewNetwork(eventGeneratorName, eventGenerator)
		if err != nil {
			t.Fatalf("Error creating new '%s' EventGenerator: %s", eventGenerator.String(), err)
		}
		if eventGenerator_under_test.Name != eventGeneratorName {
			t.Fatalf("EventGenerator has wrong name '%s', expected '%s'!", eventGenerator_under_test.Name, eventGeneratorName)
		}
		if eventGenerator_under_test.Type != eventGenerator {
			t.Fatalf("EventGenerator has wrong type '%s', expected '%s'!", eventGenerator_under_test.Type, eventGenerator)
		}
	}
}
