package experiment

type NodeGroup struct {
	Prefix  string
	NoNodes int

	MovementModel string

	IPv4Net  string
	IPv4Mask int
	IPv6Net  string
	IPv6Mask int

	NetworkType string
	Bandwidth   int
	Range       int
	Jitter      int
	Delay       int
	Error       int
	Promiscuous int
}

var defaultValues = NodeGroup{
	MovementModel: "static",

	IPv4Net:  "10.0.0.0",
	IPv4Mask: 24,
	IPv6Net:  "2001::",
	IPv6Mask: 120,

	NetworkType: "WIRELESS_LAN",
	Range:       180,
	Bandwidth:   54000000,
	Jitter:      0,
	Delay:       20000,
	Error:       0,
	Promiscuous: 0,
}

func NewNodeGroup(prefix string, noNodes int) NodeGroup {
	out := defaultValues
	out.Prefix = prefix
	out.NoNodes = noNodes
	return out
}
