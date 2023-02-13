package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/eventgenerators"
)

// A EventGenerator generates events 
type EventGenerator struct {
	Name string
	Type eventgenerators.EventGenerator
	NoEventGenerator uint
}

// EventGenerator returns a EventGenerator of the given EventGeneratorType
func NewEventGenerator(name string, eventGenerator eventgenerators.EventGenerator, NoEventGenerator uint) (EventGenerator, error) {
	if len(name) == 0 {
		return EventGenerator{}, errors.New("name of the EventGenerator must consist of at least on character")
	}
	return EventGenerator{
		Name: name,
		Type: eventGenerator,
		NoEventGenerator: NoEventGenerator,
	}, nil
}

// NewDefaultEventGEnerator returns a EventGenerator of the default EventGeneratorType
func NewDefaultEventGenerator(name string) (EventGenerator, error) {
	return NewEventGenerator(name, eventgenerators.MessageEventGenerator{}.Default(), 1)
}
