package eventgenerators

type MessageBurstEventGenerator struct {
	// Movement enabled
	Interval uint
	// Bandwidth (bps)
	size uint
	// Max range (pixels)
	Range int
	// Transmission jitter (usec)
	Jitter int
	// Transmission delay (usec)
	Delay int
	// Loss Initial
	LossInitial float32
	// Loss Factor
	LossFactor float32
	// Loss Start Range (pixels)
	LossStartRange float32
}

func (MessageBurstEventGenerator) String() string {
	return "MessageBurstEventGenerator"
}

// Returns a new configuration of Wireless with default values applied.
func (MessageBurstEventGenerator) Default() MessageBurstEventGenerator {
	return MessageBurstEventGenerator{
		Interval:       20,
		size:           54000000,
		Range:          400,
		Jitter:         0,
		Delay:          5000,
		LossInitial:    0.0,
		LossFactor:     1.0,
		LossStartRange: 300.0,
	}
}
