package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readValuesCSV(path string) ([]float64, error) {
	// open the CSV file
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open input file %q: %w", path, err)
	}

	defer file.Close()

	// create a new CSV reader and read all values
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("failed to read CSV %q: %w", path, err)
	}

	if len(records) == 0 {
		return nil, fmt.Errorf("CSV file %q is empty", path)
	}

	valueIndex := -1
	header := records[0]

	for index, columnName := range header {
		if columnName == "value" {
			valueIndex = index
			break
		}
	}

	if valueIndex == -1 {
		return nil, fmt.Errorf("CSV file %q does not contain a value column", path)
	}
	values := make([]float64, 0, len(records)-1)

	for rowIndex, record := range records[1:] {
		csvRow := rowIndex + 2
		if valueIndex >= len(record) {
			return nil, fmt.Errorf("CSV row %d does not contain a value field", csvRow)
		}

		textValue := strings.TrimSpace(record[valueIndex])
		if textValue == "" {
			return nil, fmt.Errorf("CSV row %d contains an empty value", csvRow)
		}

		numericValue, err := strconv.ParseFloat(textValue, 64)
		if err != nil {
			return nil, fmt.Errorf(
				"parse value %q on CSV row %d: %w",
				textValue,
				csvRow,
				err,
			)
		}

		values = append(values, numericValue)
	}

	if len(values) == 0 {
		return nil, fmt.Errorf("CSV file %q contains no observations", path)
	}

	return values, nil
}
