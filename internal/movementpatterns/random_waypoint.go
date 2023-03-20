package movementpatterns

type RandomWaypoint struct {
	MinSpeed int
	MaxSpeed int
	MaxPause int
}

func (RandomWaypoint) String() string {
	return "Random Waypoint"
}

func (RandomWaypoint) Default() MovementPattern {
	return RandomWaypoint{
		MinSpeed: 1,
		MaxSpeed: 5,
		MaxPause: 3600,
	}
}
