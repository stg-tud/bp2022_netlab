package experiment

import (
	"errors"

	"github.com/stg-tud/bp2022_netlab/internal/eventgenerators"
)

// A Network represents a network that nodes from NodeGroups can connect to
type EventGenerator struct {
	Name string
	Type eventgenerators.EventGenerator
}

// NewNetwork returns a Network of the given NetworkType
func NewEventGenerator(name string, eventGenerator eventgenerators.EventGenerator) (EventGenerator, error) {
	if len(name) == 0 {
		return EventGenerator{}, errors.New("name of the EventGenerator must consist of at least on character")
	}
	return EventGenerator{
		Name: name,
		Type: eventGenerator,
	}, nil
}

// NewDefaultNetwork returns a Network of the default NetworkType
func NewDefaultEventGenerator(name string) (EventGenerator, error) {
	return NewEventGenerator(name, eventgenerators.MessageEventGenerator{}.Default())
}
