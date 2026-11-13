package main

import (
	"errors"
	"sort"
)

func calculateMedian(values []float64) (float64, error) {
	if len(values) == 0 {
		return 0, errors.New("median cannot be calculated from empty input")
	}

	sortedValues := append([]float64(nil), values...)
	sort.Float64s(sortedValues)

	valueCount := len(sortedValues)
	middleIndex := valueCount / 2

	if valueCount%2 == 1 {
		return sortedValues[middleIndex], nil
	}

	leftMiddle := sortedValues[middleIndex-1]
	rightMiddle := sortedValues[middleIndex]
	median := (leftMiddle + rightMiddle) / 2.0

	return median, nil
}
