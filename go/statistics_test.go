package main

import "testing"

func TestCalculateSampleStandardDeviationKnownValues(t *testing.T) {
	values := []float64{1, 2, 3}
	expected := 1.0

	actual, err := calculateSampleStandardDeviation(values)
	if err != nil {
		t.Fatalf(
			"calculateSampleStandardDeviation returned an unexpected error: %v",
			err,
		)
	}

	if actual != expected {
		t.Errorf(
			"calculateSampleStandardDeviation() = %v; expected %v",
			actual,
			expected,
		)
	}
}

func TestCalculateSampleStandardDeviationTooFewValues(t *testing.T) {
	values := []float64{5}

	_, err := calculateSampleStandardDeviation(values)
	if err == nil {
		t.Fatal(
			"calculateSampleStandardDeviation expected an error for fewer than two values",
		)
	}
}
