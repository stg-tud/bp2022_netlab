package customtypes

// X to Y seconds specifies that every x to y seconds something should happen
type Interval struct {
	From int
	To   int
}
