package experiment

// A NodeType represents the type of a node in the network simmulation, e.g. a Router or a PC.
type NodeType int

const (
	// Router
	NODE_TYPE_ROUTER NodeType = iota
	// PC
	NODE_TYPE_PC
)

func (n NodeType) String() string {
	switch n {
	case NODE_TYPE_PC:
		return "PC"

	case NODE_TYPE_ROUTER:
		return "Router"

	default:
		return ""
	}
}
