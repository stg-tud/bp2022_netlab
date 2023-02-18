package movementpatterns

type Static struct{}

func (Static) String() string {
	return "Static"
}

func (Static) Default() MovementPattern {
	return Static{}
}
