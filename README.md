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

Standard-library packages are included with Go and will be used for the project. `mt19937` is being considered because R uses Mersenne Twister by default to generate random numbers. Using the same random-number algorithm improves comparability, although it does not guarantee identical resamples.

- Go standard library: https://pkg.go.dev/std
- Mersenne Twister package: https://github.com/seehuhn/mt19937

## Repository Structure

## Data and Experimental Design

### Go Environment:

- Go version: 1.27.1
- Go operating system and architecture: darwin/arm64
- Go module: `github.com/gummybeardream/bootstrap-r-go`

R and Go will read the same saved observations from the shared dataset, `data/symmetric.csv`.

- Distribution: normal and symmetric
- Observation count: 1,000
- Data-generation seed: 2025
- Shared file: data/symmetric.csv
- Planned bootstrap repetitions: 10,000
- R version: 4.3.2
- `boot` package version 1.3.28.1


## R Implementation

The program reads `data/symmetric.csv`. The `value` column contains 1,000 observations. The `boot` package performs 10,000 resamples with replacement while the defined statistic function calculates the median of each resample. Seed 2026 makes the bootstrap reproducible. The standard deviation of the bootstrap medians estimates the median's standard error. 

### R Baseline Results

| Dataset | Observations | Bootstrap Seed | Bootstrap Repetitions | Original Median | Bootstrap Standard Error |
|---|---:|---:|---:|---:|---:|
| Symmetric normal | 1,000 | 2026 | 10,000 | -0.002755907 | 0.03938473 |

Using 10,000 bootstrap resamples, the estimated standard error of the sample median was 0.03938473. This represents the approximate variation of the median across repeated samples. Running the program twice with the same seed produced identical results, demonstrating reproducibility.

## Go Implementation

## Testing

## Benchmarking

## Profiling and Logging

## Results

## Cloud Cost Analysis

## Recommendation

## Installation and Usage

## Building the Go Executable

## GenAI Tools

## References