package experiment

import (
	"math/rand"
	"time"
)

// GenerateRandomSeed returns a random seed used for random generations
func GenerateRandomSeed() int64 {
	source := rand.NewSource(time.Now().UnixNano())
	randObj := rand.New(source)
	seed := randObj.Int63()
	return seed
}

// GetRandomSource returns a randomness Rand object for the given randomSeed
func GetRandomSource(randomSeed int64) *rand.Rand {
	source := rand.NewSource(randomSeed)
	return rand.New(source)
}
