package experiment

type Target int

const (
	TARGET_CORE Target = iota
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
