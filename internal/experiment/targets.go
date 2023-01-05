package experiment

// A Target is a platform for which output should be generated.
type Target int

const (
	// Target: Core
	TARGET_CORE Target = iota
	// Target: The ONE
	TARGET_THEONE
)

func (t Target) String() string {
	switch t {
	case TARGET_CORE:
		return "Core"

	case TARGET_THEONE:
		return "The ONE"
	}
	return "Unknown target"
}
