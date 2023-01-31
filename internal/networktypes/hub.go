package networktypes

// Hub represents a hub based network configuration
type Hub struct{}

func (Hub) String() string {
	return "Hub"
}

// Returns a new configuration of Hub with default values applied.
func (Hub) Default() Hub {
	return Hub{}
}
