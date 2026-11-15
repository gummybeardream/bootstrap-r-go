package main

import (
	"math/rand"
	"path/filepath"
	"testing"

	"github.com/seehuhn/mt19937"
)

var (
	benchmarkFloatResult float64
	benchmarkValueCount  int
)

var benchmarkInputPath = filepath.Join(
	"..",
	"data",
	"symmetric.csv",
)

func BenchmarkReadValuesCSV(b *testing.B) {
	b.ReportAllocs()

	for iteration := 0; iteration < b.N; iteration++ {
		values, err := readValuesCSV(benchmarkInputPath)
		if err != nil {
			b.Fatalf("readValuesCSV returned an unexpected error: %v", err)
		}

		benchmarkValueCount = len(values)
	}
}

func BenchmarkCalculateMedian(b *testing.B) {
	values, err := readValuesCSV(benchmarkInputPath)
	if err != nil {
		b.Fatalf("read benchmark data: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for iteration := 0; iteration < b.N; iteration++ {
		result, err := calculateMedian(values)
		if err != nil {
			b.Fatalf("calculateMedian returned an unexpected error: %v", err)
		}

		benchmarkFloatResult = result
	}
}

func BenchmarkBootstrapMedianStandardError(b *testing.B) {
	values, err := readValuesCSV(benchmarkInputPath)
	if err != nil {
		b.Fatalf("read benchmark data: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for iteration := 0; iteration < b.N; iteration++ {
		source := mt19937.New()
		source.Seed(bootstrapSeed)
		rng := rand.New(source)

		result, err := bootstrapMedianStandardError(
			values,
			bootstrapRepetitions,
			rng,
		)
		if err != nil {
			b.Fatalf(
				"bootstrapMedianStandardError returned an unexpected error: %v",
				err,
			)
		}

		benchmarkFloatResult = result
	}
}
