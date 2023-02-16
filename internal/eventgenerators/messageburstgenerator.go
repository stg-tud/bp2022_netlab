package eventgenerators

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

type MessageBurstGenerator struct {
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

func (MessageBurstGenerator) String() string {
	return "MessageBurstGenerator"
}

// Returns a new configuration of MessageBurstGenerator with default values applied.
func (MessageBurstGenerator) Default() MessageBurstGenerator {
	return MessageBurstGenerator{
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
