package movementpatterns

type Slaw struct {
	NumberOfWaypoints int
	MinPause          int
	MaxPause          int
	// Levy exponent for pause time
	LevyExponentPause float32
	// Hurst parameter of self-similarity of waypoints
	HurstParameter float32
	DistanceWeight float32
	// Clustering range (meter)
	ClusteringRange int
	ClusterRatio    float32
	WaypointRatio   float32
}

func (Slaw) String() string {
	return "SLAW: Self-similar Least Action Walk"
}

func (Slaw) Default() MovementPattern {
	return Slaw{
		NumberOfWaypoints: 500,
		MinPause:          10,
		MaxPause:          50,
		LevyExponentPause: 1.000,
		HurstParameter:    0.750,
		DistanceWeight:    3.000,
		ClusteringRange:   50,
		ClusterRatio:      3.000,
		WaypointRatio:     5.000,
	}
}
