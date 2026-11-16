# Bootstrap Performance Comparison: R and Go
R and Go bootstrap sampling comparison

## Project Overview

## Management Problem

## Selected Statistical Method

The selected statistical method is the bootstrap estimation of the median's standard error. The median is the middle value after sorting all observations. Standard error measures how much an estimate would vary across repeated samples. Bootstrap resampling approximates how an estimate can vary across repeated samples. Bootstrap resampling replicates this variation by repeatedly sampling from the observed data with replacement. This method is computationally intensive because it may repeat the calculation thousands of times.

## Bootstrap Methodology

## R Package Research

The `boot` package was chosen because it is an established R package that handles repeated resampling while allowing the user to define the statistic being measured.

`boot`: https://cran.r-project.org/package=boot

`boot manual`: https://cran.r-project.org/web/packages/boot/boot.pdf

`R median`: https://stat.ethz.ch/R-manual/R-devel/library/stats/html/median.html

## Go Package Research

| Package | Type | Planned use |
|---|---|---|
| `encoding/csv` | Go standard library | Read the shared CSV dataset |
| `strconv` | Go standard library | Convert CSV text into numerical values |
| `sort` | Go standard library | Sort values when calculating the median |
| `math` | Go standard library | Calculate the square root used in standard deviation |
| `os` | Go standard library | Open the input file |
| `log` | Go standard library | Record application events and errors |
| `mt19937` | Third-party | Generate reproducible Mersenne Twister random numbers |

Standard-library packages are included with Go and will be used for the project. The `mt19937` package was selected because R uses a Mersenne Twister generator by default. The Go package implements the 64-bit Mersenne Twister and provides reproducible random-number generation. Using generators from the same algorithm family improves the experimental comparison, but R and Go are not expected to produce identical resamples from the same seed.

- Go standard library: https://pkg.go.dev/std
- Mersenne Twister package: https://github.com/seehuhn/mt19937

## Repository Structure

- `data/symmetric.csv`: Shared input dataset used by R and Go.
- `r/generate_data.R`: Generates the symmetric dataset.
- `r/bootstrap_median.R`: Runs the R bootstrap baseline.
- `go/data.go`: Reads and validates the shared CSV data.
- `go/data_test.go`: Tests valid and invalid CSV inputs.
- `go/median.go`: Calculates the median without changing the input.
- `go/median_test.go`: Tests the median calculation.
- `go/bootstrap.go`: Creates one bootstrap sample with replacement.
- `go/bootstrap_test.go`: Tests bootstrap sample size, values, validation, and reproducibility.
- `go/statistics.go`: Calculates sample standard deviation.
- `go/statistics_test.go`: Tests known results and invalid input.
- `go/main.go`: Runs the complete sequential Go bootstrap application.
- `go/statistics.go`: Calculates sample standard deviation.
- `go/statistics_test.go`: Tests the standard-deviation calculation.
- `results/logs/`: Contains saved program and test output.
- `results/logs/go_application.log`: Example Go application log and statistical output.

## Data and Experimental Design

### Go Environment

- Go version: 1.27.1
- Go operating system and architecture: darwin/arm64
- Go module: `github.com/gummybeardream/bootstrap-r-go`

R and Go will read the same saved observations from the shared dataset, `data/symmetric.csv`.

- Distribution: normal and symmetric
- Observation count: 1,000
- Data-generation seed: 2025
- Shared file: `data/symmetric.csv`
- Planned bootstrap repetitions: 10,000
- R version: 4.3.2
- `boot` package version: 1.3.28.1

## R Implementation

The program reads `data/symmetric.csv`. The `value` column contains 1,000 observations. The `boot` package performs 10,000 resamples with replacement while the defined statistic function calculates the median of each resample. Seed 2026 makes the bootstrap reproducible. The standard deviation of the bootstrap medians estimates the median's standard error.

### R Baseline Results

| Dataset | Observations | Bootstrap Seed | Bootstrap Repetitions | Original Median | Bootstrap Standard Error |
|---|---:|---:|---:|---:|---:|
| Symmetric normal | 1,000 | 2026 | 10,000 | -0.002755907 | 0.03938473 |

Using 10,000 bootstrap resamples, the estimated standard error of the sample median was 0.03938473. This represents the approximate variation of the median across repeated samples. Running the program twice with the same seed produced identical results, demonstrating reproducibility.

## Go Implementation

The sequential Go application uses the shared dataset, seed 2026, and 10,000 bootstrap repetitions. It creates each bootstrap sample with replacement, calculates its median, and estimates the median's standard error using the sample standard deviation of the 10,000 bootstrap medians. Two executions using the same seed produced identical results.

### Go Median Baseline

| Dataset | Observations | R Median | Go Median | Match |
|---|---:|---:|---:|---|
| Symmetric normal | 1,000 | -0.002755907 | -0.002755907 | Yes |

The initial Go application successfully read the shared CSV dataset and produced the same original median as R. This confirms that both implementations use the same observations and median definition.

## Testing

| Component | Test coverage | Tests | Result |
|---|---|---:|---|
| Median calculation | Odd count, even count, single value, empty input, input preservation | 5 | Passed |
| CSV reader | Valid input, reordered columns, missing column, invalid number, empty value, header only, missing file, empty file | 8 | Passed |
| Bootstrap sampling | Sample size, source values, empty input, nil generator, reproducibility | 5 | Passed |
| Sample standard deviation | Known result and insufficient input | 2 | Passed |
| Sequential bootstrap | Constant data, invalid repetitions, empty input, nil generator, reproducibility | 5 | Passed |
| Total | Current Go unit-test suite | 25 | Passed |

Tests were run with `go test -count=1 -v ./go`. The `-count=1` option disabled Go's test cache so that the saved output represented a fresh test execution.

## Benchmarking

### Initial Go Microbenchmarks

The Go benchmarks were run five times on an Apple M1 using Go 1.27.1 on `darwin/arm64`. Unit tests were skipped during measurement, and `-benchmem` reported allocation statistics.

| Component | Average Time per Operation | Bytes per Operation | Allocations per Operation |
|---|---:|---:|---:|
| CSV reading | 0.126 ms | 110,248 | 2,023 |
| Median calculation | 0.013 ms | 8,192 | 1 |
| Sequential bootstrap, 10,000 repetitions | 510.754 ms | 163,925,066 | 20,007 |

The sequential bootstrap was the most expensive component. It allocated approximately 164 MB across about 20,007 allocations per operation. These values represent total allocation activity, not peak process memory.

Run benchmarks with this command:

```bash
go test -run '^$' -bench . -benchmem -count 5 ./go
```


## Profiling and Logging

The Go application uses the standard-library `log` package to record major execution events. It logs application startup, the input file, the number of observations loaded, the bootstrap seed and repetition count, successful completion of the calculation, and application completion. Fatal input or calculation errors are also reported through the logger.

Individual bootstrap repetitions are not logged because 10,000 log messages would produce excessive output and distort performance measurements.

Logs are written to standard error, while statistical results are written to standard output. Both streams were captured in `results/logs/go_application.log` using:

```text
go run ./go 2>&1 | tee results/logs/go_application.log
```
### Baseline Go Profiles

CPU profiling showed that median sorting was the primary application-level CPU cost. The `calculateMedian` and sorting calls accounted for approximately 45.45% of cumulative CPU samples. Random-number index generation accounted for a much smaller share.

The allocation-space memory profile showed that `calculateMedian` allocated 79.62 MB and `bootstrapSample` allocated 77.60 MB. Together, these functions accounted for approximately 97.99% of allocated memory. The program created one slice for each bootstrap sample and a second copy when calculating each median.

### First Go Optimization

The baseline memory profile showed that each bootstrap repetition allocated one slice in `bootstrapSample` and a second slice when `calculateMedian` copied the sample before sorting it.

The optimized implementation introduced `calculateMedianInPlace`. The original `calculateMedian` function continues to copy its input, preserving its general-purpose behavior. The sequential bootstrap uses the in-place function because each bootstrap sample is newly created and is not needed after its median is calculated.

| Measurement | Before Optimization | After Optimization | Reduction |
|---|---:|---:|---:|
| Average runtime | 510.754 ms | 496.756 ms | 2.74% |
| Bytes allocated per operation | 163,925,066 | 82,004,763 | 49.97% |
| Allocations per operation | 20,007 | 10,004 | 50.00% |

The optimization removed approximately half of the allocation activity while preserving all statistical results. Runtime improved by approximately 2.74%, showing that sorting remained the main CPU cost even after the extra slice copy was removed.

## Results

### Sequential Go Bootstrap Results

| Dataset | Observations | Repetitions | Original Median | Bootstrap Standard Error |
|---|---:|---:|---:|---:|
| Symmetric normal | 1,000 | 10,000 | -0.002755907 | 0.038877677 |

### Initial R and Go Comparison

| Measurement | R | Sequential Go | Absolute Difference | Relative Difference |
|---|---:|---:|---:|---:|
| Original median | -0.002755907 | -0.002755907 | 0 | 0% |
| Bootstrap standard error | 0.039384730 | 0.038877677 | 0.000507053 | 1.29% |

The original medians match exactly because both programs use the same saved observations and median definition. The bootstrap standard errors differ by approximately 1.29%. This difference is expected because the R and Go random-number generators do not produce identical resampling sequences.

## Cloud Cost Analysis

## Recommendation

## Installation and Usage

## Building the Go Executable

## GenAI Tools

## References