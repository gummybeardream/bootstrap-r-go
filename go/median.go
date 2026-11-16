package main

import (
	"errors"
	"sort"
)

func calculateMedian(values []float64) (float64, error) {
	sortedValues := append([]float64(nil), values...)

	return calculateMedianInPlace(sortedValues)
}

func calculateMedianInPlace(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("median cannot be calculated from empty input")
	}

	sort.Float64s(values)

	valueCount := len(values)
	middleIndex := valueCount / 2

	if valueCount%2 == 1 {
		return values[middleIndex], nil
	}

	leftMiddle := values[middleIndex-1]
	rightMiddle := values[middleIndex]
	median := (leftMiddle + rightMiddle) / 2.0

	return median, nil
}
