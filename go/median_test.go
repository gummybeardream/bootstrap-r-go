package main

import (
	"slices"
	"testing"
)

func TestCalculateMedianOddCount(t *testing.T) {
	values := []float64{9, 1, 5}
	expected := 5.0

	actual, err := calculateMedian(values)
	if err != nil {
		t.Fatalf("calculateMedian returned an unexpected error: %v", err)
	}

	if actual != expected {
		t.Errorf("calculateMedian() = %v; expected %v", actual, expected)
	}
}

func TestCalculateMedianEvenCount(t *testing.T) {
	values := []float64{8, 2, 6, 4}
	expected := 5.0

	actual, err := calculateMedian(values)
	if err != nil {
		t.Fatalf("calculateMedian returned an unexpected error: %v", err)
	}

	if actual != expected {
		t.Errorf("calculateMedian() = %v; expected %v", actual, expected)
	}
}

func TestCalculateMedianSingleValue(t *testing.T) {
	values := []float64{7.5}
	expected := 7.5

	actual, err := calculateMedian(values)
	if err != nil {
		t.Fatalf("calculateMedian returned an unexpected error: %v", err)
	}

	if actual != expected {
		t.Errorf("calculateMedian() = %v; expected %v", actual, expected)
	}

}

func TestCalculateMedianEmptyInput(t *testing.T) {
	values := []float64{}

	_, err := calculateMedian(values)
	if err == nil {
		t.Fatal("calculateMedian() expected an error for empty input")
	}
}

func TestCalculateMedianPreservesInput(t *testing.T) {
	values := []float64{3, 1, 2}
	original := append([]float64(nil), values...)

	_, err := calculateMedian(values)
	if err != nil {
		t.Fatalf("calculateMedian returned an unexpected error: %v", err)
	}
	if !slices.Equal(values, original) {
		t.Errorf("calculateMedian modified the input: got %v; expected %v", values, original)
	}
}

func TestCalculateMedianInPlace(t *testing.T) {
	values := []float64{9, 1, 5}
	expectedMedian := 5.0
	expectedOrder := []float64{1, 5, 9}

	actualMedian, err := calculateMedianInPlace(values)
	if err != nil {
		t.Fatalf(
			"calculateMedianInPlace returned an unexpected error: %v",
			err,
		)
	}

	if actualMedian != expectedMedian {
		t.Errorf(
			"calculateMedianInPlace() = %v; expected %v",
			actualMedian,
			expectedMedian,
		)
	}

	if !slices.Equal(values, expectedOrder) {
		t.Errorf(
			"calculateMedianInPlace order = %v; expected %v",
			values,
			expectedOrder,
		)
	}
}
