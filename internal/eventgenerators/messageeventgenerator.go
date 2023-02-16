package eventgenerators

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

type MessageEventGenerator struct {
	// Creation interval in seconds (one new message every X to Y seconds)
	Interval customtypes.XToYSeconds
	// Size of the message
	Size customtypes.XToYSeconds
	// range of message source/destination addresses
	Hosts customtypes.XToYSeconds
	// distance to the host
	ToHosts customtypes.XToYSeconds
	// Message ID prefix
	Prefix string
}

func (MessageEventGenerator) String() string {
	return "MessageEventGenerator"
}

// Returns a new configuration of Messageeventgenerator with default values applied.
func (MessageEventGenerator) Default() MessageEventGenerator {
	return MessageEventGenerator{
		Interval: customtypes.XToYSeconds{
			XSeconds: 25,
			YSeconds: 35,
		},
		Size: customtypes.XToYSeconds{
			XSeconds: 80,
			YSeconds: 120,
		},
		Hosts: customtypes.XToYSeconds{
			XSeconds: 5,
			YSeconds: 15,
		},
		ToHosts: customtypes.XToYSeconds{
			XSeconds: 16,
			YSeconds: 17,
		},
		Prefix: "M",
	}
}
