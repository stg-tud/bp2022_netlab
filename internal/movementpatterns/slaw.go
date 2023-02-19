package movementpatterns

type SLAW struct {
	// Number of waypoints to generate during initialization
	NumberOfWaypoints int
	// Minimum pause time in seconds
	MinPause int
	// Maximum pause time in seconds
	MaxPause int
	// Levy exponent for pause time
	LevyExponent float32
	// Hurst parameter of self-similarity of waypoints
	HurstParameter float32
	// Distance weight used by LATP algorithm
	DistanceWeight float32
	// Clustering range in meters
	ClusteringRange float32
	// Cluster ratio as divisor (ratio is then 1/divisor)
	ClusterRatio int
	// Waypoint ratio as divisor (ratio is then 1/divisor)
	WaypointRatio int
}

func (SLAW) String() string {
	return "SLAW: Self-similar Least Action Walk"
}

func (SLAW) Default() MovementPattern {
	return SLAW{
		NumberOfWaypoints: 500,
		MinPause:          10,
		MaxPause:          50,
		LevyExponent:      1.000,
		HurstParameter:    0.750,
		DistanceWeight:    3.000,
		ClusteringRange:   50.000,
		ClusterRatio:      3,
		WaypointRatio:     5,
	}
}
