package networktypes

// Emane represents a emane network configuration
type Emane struct{}

func (Emane) String() string {
	return "Emane"
}

// Returns a new configuration of Emane with default values applied.
func (Emane) Default() Emane {
	return Emane{}
}
