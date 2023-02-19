package movementpatterns

type SMOOTH struct {
	// Transmission range of mobile nodes
	Range int
	// Total number of clusters in the network
	Clusters int
	// Alpha value for the flight distribution
	Alpha float32
	// Minimum value for the flight distribution
	MinFlight int
	// Maximum value for the flight distribution
	MaxFlight int
	// Beta value for the pause-time distribution
	Beta float32
	// Minimum value for the pause-time distribution
	MinPause int
	// Maximum value for the pause-time distribution
	MaxPause int
}

func (SMOOTH) String() string {
	return "SMOOTH"
}

func (SMOOTH) Default() MovementPattern {
	return SMOOTH{
		Range:     100,
		Clusters:  40,
		Alpha:     1.450,
		MinFlight: 1,
		MaxFlight: 14000,
		Beta:      1.500,
		MinPause:  10,
		MaxPause:  3600,
	}
}
