package movementpatterns

type Swim struct {
	Radius                float32
	CellDistanceWeight    float32
	NodeSpeedMultiplier   float32
	WaitingTimeExponent   float32
	WaitingTimeUpperBound float32
}

func (Swim) String() string {
	return "SWIM: Small World In Motion"
}

func (Swim) Default() MovementPattern {
	return Swim{
		Radius:                0.200,
		CellDistanceWeight:    0.500,
		NodeSpeedMultiplier:   1.500,
		WaitingTimeExponent:   1.550,
		WaitingTimeUpperBound: 100.000,
	}
}
