package movementpatterns

type RandomWaypoint struct {
	MinSpeed int
	MaxSpeed int
	MaxPause int
}

func (RandomWaypoint) String() string {
	return "Random Waypoint"
}
