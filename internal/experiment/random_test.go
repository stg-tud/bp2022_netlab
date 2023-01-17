package experiment_test

import (
	"math/rand"
	"testing"

	"github.com/stg-tud/bp2022_netlab/internal/experiment"
)

func TestGenerateRandomSeed(t *testing.T) {
	const RUNS = 15
	seeds := make(map[int64]bool)
	for i := 0; i < RUNS; i++ {
		seed := experiment.GenerateRandomSeed()
		_, present := seeds[seed]
		if present {
			t.Fatal("GenerateRandomSeed() returned the same seed multiple times!")
		}
		seeds[seed] = true
	}
}

func TestGetRandomSource(t *testing.T) {
	const RUNS = 15
	var seed int64 = 1673916420049
	randObjects := [RUNS]*rand.Rand{}
	for i := 0; i < RUNS; i++ {
		randObjects[i] = experiment.GetRandomSource(seed)
	}
	for y := 0; y < RUNS; y++ {
		var previousValue int = randObjects[0].Int()
		for i := 1; i < RUNS; i++ {
			currentValue := randObjects[i].Int()
			if currentValue != previousValue {
				t.Fatal("GetRandomSource() returned Rand objects that do not have the same outputs!")
			}
		}
	}
}
