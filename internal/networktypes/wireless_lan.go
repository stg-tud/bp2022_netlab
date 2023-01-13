package networktypes

// WirelessLAN represents a wireless LAN network configuration
type WirelessLAN struct {
	// Bandwidth (bps)
	Bandwidth int
	// Wireless range (pixels)
	Range int
	// Transmission jitter (usec)
	Jitter int
	// Transmission delay (usec)
	Delay int
	// Loss (%)
	Loss float32
	// Promiscuous mode
	Promiscuous bool
}

var defaultValues = WirelessLAN{
	Bandwidth:   54000000,
	Range:       275,
	Jitter:      0,
	Delay:       5000,
	Loss:        0.0,
	Promiscuous: false,
}

func (WirelessLAN) String() string {
	return "Wireless LAN"
}

// Returns a new configuration of WirelessLAN with default values applied.
func (WirelessLAN) Default() WirelessLAN {
	return defaultValues
}
