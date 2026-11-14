package main

import (
	"os"
	"path/filepath"
	"slices"
	"testing"
)

func writeTestCSV(t *testing.T, content string) string {
	t.Helper()

	path := filepath.Join(t.TempDir(), "test.csv")

	err := os.WriteFile(path, []byte(content), 0o600)
	if err != nil {
		t.Fatalf("write temporary CSV: %v", err)
	}

	return path
}

func TestReadValuesCSVValidFile(t *testing.T) {
	path := writeTestCSV(t, "value\n1.5\n2.5\n3.5\n")
	expected := []float64{1.5, 2.5, 3.5}

	actual, err := readValuesCSV(path)
	if err != nil {
		t.Fatalf("readValuesCSV returned an unexpected error: %v", err)
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("readValuesCSV() = %v; expected %v", actual, expected)
	}
}

func TestReadValuesCSVValueColumnNotFirst(t *testing.T) {
	path := writeTestCSV(
		t,
		"label,value\nfirst,1.5\nsecond,2.5\n",
	)
	expected := []float64{1.5, 2.5}

	actual, err := readValuesCSV(path)
	if err != nil {
		t.Fatalf("readValuesCSV returned an unexpected error: %v", err)
	}

	if !slices.Equal(actual, expected) {
		t.Errorf("readValuesCSV() = %v; expected %v", actual, expected)
	}
}

func TestReadValuesCSVMissingValueColumn(t *testing.T) {
	path := writeTestCSV(t, "amount\n1.5\n")

	_, err := readValuesCSV(path)
	if err == nil {
		t.Fatal("readValuesCSV() expected an error for a missing value column")
	}
}

func TestReadValuesCSVInvalidNumber(t *testing.T) {
	path := writeTestCSV(t, "value\nnot-a-number\n")

	_, err := readValuesCSV(path)
	if err == nil {
		t.Fatal("readValuesCSV() expected an error for invalid numerical text")
	}
}

func TestReadValuesCSVEmptyValue(t *testing.T) {
	path := writeTestCSV(t, "value\n\"\"\n")

	_, err := readValuesCSV(path)
	if err == nil {
		t.Fatal("readValuesCSV() expected an error for an empty value")
	}
}

func TestReadValuesCSVHeaderOnly(t *testing.T) {
	path := writeTestCSV(t, "value\n")

	_, err := readValuesCSV(path)
	if err == nil {
		t.Fatal("readValuesCSV() expected an error for a header-only file")
	}
}

func TestReadValuesCSVMissingFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "missing.csv")

	_, err := readValuesCSV(path)
	if err == nil {
		t.Fatal("readValuesCSV() expected an error for a missing file")
	}
}

func TestReadValuesCSVEmptyFile(t *testing.T) {
	path := writeTestCSV(t, "")

	_, err := readValuesCSV(path)
	if err == nil {
		t.Fatal("readValuesCSV() expected an error for an empty file")
	}
}
