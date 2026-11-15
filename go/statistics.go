package main

import (
	"errors"
	"math"
)

func calculateSampleStandardDeviation(values []float64) (float64, error) {
	if len(values) < 2 {
		return 0, errors.New("sample standard deviation requires at least two values")
	}

	var sum float64
	for _, value := range values {
		sum += value
	}

	mean := sum / float64(len(values))

	var squaredDifferenceSum float64
	for _, value := range values {
		difference := value - mean
		squaredDifferenceSum += difference * difference
	}

	variance := squaredDifferenceSum / float64(len(values)-1)

	return math.Sqrt(variance), nil
}
