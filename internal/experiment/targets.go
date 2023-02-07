package experiment

// A Target is a platform for which output should be generated.
type Target uint

const (
	// Target: Core
	TargetCore Target = iota
	//
	TargetCoreEmulab
	// Target: The ONE
	TargetTheOne
)

func (t Target) String() string {
	switch t {
	case TargetCore:
		return "Core"

	case TargetCoreEmulab:
		return "CoreEmulab"

	case TargetTheOne:
		return "The ONE"

	default:
		return "Unknown target"
	}
}
