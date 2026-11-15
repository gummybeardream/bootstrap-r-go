package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/seehuhn/mt19937"
)

const (
	inputPath                  = "data/symmetric.csv"
	bootstrapSeed        int64 = 2026
	bootstrapRepetitions       = 10000
)

func main() {
	log.SetPrefix("bootstrap-r-go: ")
	log.Println("application started")

	log.Printf("reading input file %s", inputPath)

	values, err := readValuesCSV(inputPath)
	if err != nil {
		log.Fatalf("read input data: %v", err)
	}

	log.Printf("loaded %d observations", len(values))

	median, err := calculateMedian(values)
	if err != nil {
		log.Fatalf("calculate original median: %v", err)
	}

	source := mt19937.New()
	source.Seed(bootstrapSeed)
	rng := rand.New(source)

	log.Printf(
		"running %d bootstrap repetitions with seed %d",
		bootstrapRepetitions,
		bootstrapSeed,
	)

	standardError, err := bootstrapMedianStandardError(
		values,
		bootstrapRepetitions,
		rng,
	)
	if err != nil {
		log.Fatalf("calculate bootstrap standard error: %v", err)
	}

	log.Println("bootstrap calculation completed")

	fmt.Println("Input file:", inputPath)
	fmt.Println("Observation count:", len(values))
	fmt.Println("Bootstrap seed:", bootstrapSeed)
	fmt.Println("Bootstrap repetitions:", bootstrapRepetitions)
	fmt.Printf("Original median: %.9f\n", median)
	fmt.Printf("Bootstrap standard error: %.9f\n", standardError)

	log.Println("application completed")
}
