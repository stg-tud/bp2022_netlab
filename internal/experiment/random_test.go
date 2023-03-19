package experiment_test

import (
	"math/rand"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
	"github.com/stretchr/testify/assert"
)

func TestGenerateRandomSeedRandomness(t *testing.T) {
	const Runs = 15
	seeds := make(map[int64]bool)
	for i := 0; i < Runs; i++ {
		seed := experiment.GenerateRandomSeed()
		_, present := seeds[seed]
		assert.False(t, present, "GenerateRandomSeed() returned the same seed multiple times!")
		seeds[seed] = true
	}
}

func TestGetRandomSourceDeterminancy(t *testing.T) {
	const Runs = 15
	var seed int64 = 1673916420049
	randObjects := [Runs]*rand.Rand{}
	for i := 0; i < Runs; i++ {
		randObjects[i] = experiment.GetRandomSource(seed)
	}
	for y := 0; y < Runs; y++ {
		var previousValue int = randObjects[0].Int()
		for i := 1; i < Runs; i++ {
			currentValue := randObjects[i].Int()
			assert.Equal(t, previousValue, currentValue, "GetRandomSource() returned Rand objects that do not have the same outputs!")
		}
	}
}
