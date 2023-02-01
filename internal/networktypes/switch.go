package networktypes

// Switch represents a switch based network configuration
type Switch struct{}

func (Switch) String() string {
	return "Switch"
}

// Returns a new configuration of Switch with default values applied.
func (Switch) Default() Switch {
	return Switch{}
}
