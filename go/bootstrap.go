package main

import (
	"errors"
	"fmt"
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

func bootstrapMedianStandardError(
	values []float64,
	repetitions int,
	rng *rand.Rand,
) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("bootstrap cannot run with empty input")
	}

	if repetitions < 2 {
		return 0, errors.New("bootstrap requires at least two repetitions")
	}

	if rng == nil {
		return 0, errors.New("random-number generator cannot be nil")
	}

	medians := make([]float64, repetitions)

	for repetition := 0; repetition < repetitions; repetition++ {
		sample, err := bootstrapSample(values, rng)
		if err != nil {
			return 0, fmt.Errorf(
				"bootstrap repetition %d: %w",
				repetition+1,
				err,
			)
		}

		median, err := calculateMedian(sample)
		if err != nil {
			return 0, fmt.Errorf(
				"calculate median for bootstrap repetition %d: %w",
				repetition+1,
				err,
			)
		}

		medians[repetition] = median
	}

	standardError, err := calculateSampleStandardDeviation(medians)
	if err != nil {
		return 0, fmt.Errorf("calculate bootstrap standard error: %w", err)
	}

	return standardError, nil
}
