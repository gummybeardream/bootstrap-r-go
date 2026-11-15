package main

import (
	"fmt"
	"log"
)

func main() {
	inputPath := "data/symmetric.csv"
	values, err := readValuesCSV(inputPath)
	if err != nil {
		log.Fatalf("read input data: %v", err)
	}

	median, err := calculateMedian(values)
	if err != nil {
		log.Fatalf("calculate median: %v", err)
	}

	fmt.Println("Input file:", inputPath)
	fmt.Println("Observation count:", len(values))
	fmt.Printf("Original median: %.9f\n", median)
}
