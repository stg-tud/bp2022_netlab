package experiment

import (
	"math/rand"
	"time"

	logger "github.com/gookit/slog"
)

// GenerateRandomSeed returns a random seed used for random generations
func GenerateRandomSeed() int64 {
	logger.Trace("Generating new random seed")
	time.Sleep(1000) // Sleep a few nanoseconds in order for Windows to update its time and allow unique timestamps
	source := rand.NewSource(time.Now().UnixNano())
	randObj := rand.New(source)
	seed := randObj.Int63()
	return seed
}

// GetRandomSource returns a randomness Rand object for the given randomSeed
func GetRandomSource(randomSeed int64) *rand.Rand {
	logger.Trace("Generating new random source")
	source := rand.NewSource(randomSeed)
	return rand.New(source)
}
