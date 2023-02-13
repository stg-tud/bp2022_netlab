package eventgenerators

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

type MessageBurstGenerator struct {
	// Class of the first event generator
	Class string
	// Creation interval in seconds (one new message every X to Y seconds)
	Interval uint
	// Size of the message
	Size customtypes.Area
	// range of message source/destination addresses
	Hosts customtypes.Area
	// Message ID prefix
	Prefix string
}

func (MessageBurstGenerator) String() string {
	return "MessageBurstGenerator"
}

// Returns a new configuration of Wireless with default values applied.
func (MessageBurstGenerator) Default() MessageBurstGenerator {
	return MessageBurstGenerator{
		Class:    "MessageBurstGenerator",
		Interval: 20,
		Size: customtypes.Area{
			Height: 80,
			Width:  120,
		},
		Hosts: customtypes.Area{
			Height: 5,
			Width:  15,
		},
		Prefix: "M",
	}
}
