package experiment

// A NodeType represents the type of a node in the network simmulation, e.g. a Router or a PC.
type NodeType int

const (
	// Router
	NodeTypeRouter NodeType = iota
	// PC
	NodeTypePC
)

func (n NodeType) String() string {
	switch n {
	case NodeTypePC:
		return "PC"

	case NodeTypeRouter:
		return "Router"

	default:
		return ""
	}
}
