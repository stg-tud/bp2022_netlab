package networktypes

// Wireless represents a wireless network configuration
type Wireless struct {
	// Movement enabled
	Movement bool
	// Bandwidth (bps)
	Bandwidth int
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

func (Wireless) String() string {
	return "Wireless"
}

// Returns a new configuration of Wireless with default values applied.
func (Wireless) Default() Wireless {
	return Wireless{
		Movement:       true,
		Bandwidth:      54000000,
		Range:          400,
		Jitter:         0,
		Delay:          5000,
		LossInitial:    0.0,
		LossFactor:     1.0,
		LossStartRange: 300.0,
	}
}
