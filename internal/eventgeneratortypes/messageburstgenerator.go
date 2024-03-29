package eventgeneratortypes

import "github.com/stg-tud/bp2022_netlab/internal/customtypes"

type MessageBurstGenerator struct {
	// Creation interval in seconds (one new message every X to Y seconds)
	Interval customtypes.Interval
	// Size of the message
	Size customtypes.Interval
	// range of message source/destination addresses
	Hosts customtypes.Interval
	// distance to the host
	ToHosts customtypes.Interval
	// Message ID prefix
	Prefix string
}

func (MessageBurstGenerator) String() string {
	return "MessageBurstGenerator"
}

// Returns a new configuration of MessageBurstGenerator with default values applied.
func (MessageBurstGenerator) Default() MessageBurstGenerator {
	return MessageBurstGenerator{
		Interval: customtypes.Interval{
			From: 25,
			To:   35,
		},
		Size: customtypes.Interval{
			From: 80,
			To:   120,
		},
		Hosts: customtypes.Interval{
			From: 5,
			To:   15,
		},
		ToHosts: customtypes.Interval{
			From: 16,
			To:   17,
		},
		Prefix: "M",
	}
}
