package eventgenerators

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

type MessageEventGenerator struct {
	// Class of the first event generator
	Class string
	// Creation interval in seconds (one new message every X to Y seconds)
	Interval customtypes.Position
	// Size of the message
	Size customtypes.Position
	// range of message source/destination addresses
	Hosts customtypes.Position
	// Message ID prefix
	Prefix string
}

func (MessageEventGenerator) String() string {
	return "MessageEventGenerator"
}

// Returns a new configuration of Messageeventgenerator with default values applied.
func (MessageEventGenerator) Default() MessageEventGenerator {
	return MessageEventGenerator{
		Class: "MessageEventGenerator",
		Interval: customtypes.Position{
			X: 25,
			Y: 35,
		},
		Size: customtypes.Position{
			X: 80,
			Y: 120,
		},
		Hosts: customtypes.Position{
			X: 5,
			Y: 15,
		},
		Prefix: "M",
	}
}
