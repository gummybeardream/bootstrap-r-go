package main

import (
	"errors"
	"math/rand"
)

func bootstrapSample(values []float64, rng *rand.Rand) ([]float64, error) {
	if len(values) == 0 {
		return nil, errors.New("bootstrap sample cannot be created from empty input")
	}

	if rng == nil {
		return nil, errors.New("random-number generator cannot be nil")
	}

	sample := make([]float64, len(values))

	for index := range sample {
		randomIndex := rng.Intn(len(values))
		sample[index] = values[randomIndex]
	}

	return sample, nil
}
