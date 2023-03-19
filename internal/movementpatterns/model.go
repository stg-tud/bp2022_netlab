package movementpatterns

type MovementPattern interface {
	String() string
	Default() MovementPattern
}
