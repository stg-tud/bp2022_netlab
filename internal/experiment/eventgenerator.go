package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/eventgeneratortypes"
)

// A EventGenerator generates events
type EventGenerator struct {
	Name string
	Type eventgeneratortypes.EventGeneratorType
}

// EventGenerator returns a EventGenerator of the given eventgeneratortypes
func NewEventGenerator(name string, EventGeneratorType eventgeneratortypes.EventGeneratorType) (EventGenerator, error) {
	if len(name) == 0 {
		return EventGenerator{}, errors.New("name of the EventGenerator must consist of at least on character")
	}
	return EventGenerator{
		Name: name,
		Type: EventGeneratorType,
	}, nil
}

// NewDefaultEventGenerator returns a EventGenerator of the default eventgeneratortypes
func NewDefaultEventGenerator(name string) (EventGenerator, error) {
	return NewEventGenerator(name, eventgeneratortypes.MessageEventGenerator{}.Default())
}
