package main

import (
	"math/rand"
	"slices"
	"testing"

	"github.com/seehuhn/mt19937"
)

func newTestRNG(seed int64) *rand.Rand {
	source := mt19937.New()
	source.Seed(seed)

	return rand.New(source)
}

func TestBootstrapSampleCorrectLength(t *testing.T) {
	values := []float64{10, 20, 30, 40}
	rng := newTestRNG(2026)

	sample, err := bootstrapSample(values, rng)
	if err != nil {
		t.Fatalf("bootstrapSample returned an unexpected error: %v", err)
	}

	if len(sample) != len(values) {
		t.Errorf(
			"bootstrapSample returned %d values; expected %d",
			len(sample),
			len(values),
		)
	}
}

func TestBootstrapSampleUsesInputValues(t *testing.T) {
	values := []float64{10, 20, 30, 40}
	rng := newTestRNG(2026)

	sample, err := bootstrapSample(values, rng)
	if err != nil {
		t.Fatalf("bootstrapSample returned an unexpected error: %v", err)
	}

	validValues := map[float64]bool{
		10: true,
		20: true,
		30: true,
		40: true,
	}

	for _, sampledValue := range sample {
		if !validValues[sampledValue] {
			t.Errorf(
				"bootstrapSample returned value %v, which is not in the input",
				sampledValue,
			)
		}
	}
}

func TestBootstrapSampleEmptyInput(t *testing.T) {
	rng := newTestRNG(2026)

	_, err := bootstrapSample([]float64{}, rng)
	if err == nil {
		t.Fatal("bootstrapSample expected an error for empty input")
	}
}

func TestBootstrapSampleNilGenerator(t *testing.T) {
	values := []float64{10, 20, 30, 40}

	_, err := bootstrapSample(values, nil)
	if err == nil {
		t.Fatal("bootstrapSample expected an error for a nil generator")
	}
}

func TestBootstrapSampleReproducible(t *testing.T) {
	values := []float64{10, 20, 30, 40}

	firstRNG := newTestRNG(2026)
	secondRNG := newTestRNG(2026)

	firstSample, err := bootstrapSample(values, firstRNG)
	if err != nil {
		t.Fatalf("first bootstrapSample call returned an unexpected error: %v", err)
	}

	secondSample, err := bootstrapSample(values, secondRNG)
	if err != nil {
		t.Fatalf("second bootstrapSample call returned an unexpected error: %v", err)
	}

	if !slices.Equal(firstSample, secondSample) {
		t.Errorf(
			"bootstrap samples differ: first %v; second %v",
			firstSample,
			secondSample,
		)
	}
}
