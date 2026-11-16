# Bootstrap Performance Comparison: R and Go
R and Go bootstrap sampling comparison

## Project Overview

This project compares R and Go implementations of the bootstrap estimation of the median's standard error. The study compares runtime and memory across 100 runs and evaluates projected cloud computing savings.

## Management Problem

A research consultancy would like to examine the performance advantages of Go compared to R, which is the software platform that they have used for years. The firm is considering moving much of their programming to Go because of growing cloud-computing costs.

## Selected Statistical Method

The selected statistical method is the bootstrap estimation of the median's standard error. The median is the middle value after sorting all observations. Standard error measures how much an estimate would vary across repeated samples. Bootstrap resampling approximates how an estimate can vary across repeated samples. Bootstrap resampling replicates this variation by repeatedly sampling from the observed data with replacement. This method is computationally intensive because it may repeat the calculation thousands of times.

## Bootstrap Methodology

R generated a symmetric dataset containing 1,000 observations from a normal distribution using seed 2025. The dataset was saved as `data/symmetric.csv` so that R and Go would process identical input values.

Each implementation performed 10,000 bootstrap repetitions. During each repetition, 1,000 observations were sampled from the original dataset with replacement. The median of each bootstrap sample was calculated and stored. The estimated standard error was the sample standard deviation of the 10,000 bootstrap medians.

R used the `boot` package with seed 2026. Go used the 64-bit Mersenne Twister package with the same numerical seed. The two generators belong to the same algorithm family but do not produce identical resampling sequences. Therefore, the original medians were expected to match exactly, while the bootstrap standard-error estimates were expected to be close rather than identical.

## R Package Research

The `boot` package was chosen because it is an established R package that handles repeated resampling while allowing the user to define the statistic being measured.

`boot`: https://cran.r-project.org/package=boot

`boot manual`: https://cran.r-project.org/web/packages/boot/boot.pdf

`R median`: https://stat.ethz.ch/R-manual/R-devel/library/stats/html/median.html

## Go Package Research

| Package | Type | Use |
|---|---|---|
| `encoding/csv` | Go standard library | Read the shared CSV dataset |
| `strconv` | Go standard library | Convert CSV text into numerical values |
| `sort` | Go standard library | Sort values when calculating the median |
| `math` | Go standard library | Calculate the square root used in standard deviation |
| `os` | Go standard library | Open the input file |
| `log` | Go standard library | Record application events and errors |
| `mt19937` | Third-party | Generate reproducible Mersenne Twister random numbers |

Standard-library packages are included with Go and were used for the project. The `mt19937` package was selected because R uses a Mersenne Twister generator by default. The Go package implements the 64-bit Mersenne Twister and provides reproducible random-number generation. Using generators from the same algorithm family improves the experimental comparison, but R and Go are not expected to produce identical resamples from the same seed.

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
- `results/logs/`: Contains saved program and test output.
- `results/logs/go_application.log`: Example Go application log and statistical output.
- `results/benchmarks/r_end_to_end_100_2026-11-15.txt`: Raw timing and memory output from 100 R runs.
- `results/benchmarks/go_end_to_end_100_2026-11-15.txt`: Raw timing and memory output from 100 compiled Go runs.

## Data and Experimental Design

### Go Environment

- Go version: 1.27.1
- Go operating system and architecture: darwin/arm64
- Go module: `github.com/gummybeardream/bootstrap-r-go`

R and Go read the same saved observations from the shared dataset, `data/symmetric.csv`.

- Distribution: normal and symmetric
- Observation count: 1,000
- Data-generation seed: 2025
- Shared file: `data/symmetric.csv`
- Bootstrap repetitions: 10,000
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
| Median calculation | Odd count, even count, single value, empty input, input preservation, in-place calculation | 6 | Passed |
| CSV reader | Valid input, reordered columns, missing column, invalid number, empty value, header only, missing file, empty file | 8 | Passed |
| Bootstrap sampling | Sample size, source values, empty input, nil generator, reproducibility | 5 | Passed |
| Sample standard deviation | Known result and insufficient input | 2 | Passed |
| Sequential bootstrap | Constant data, invalid repetitions, empty input, nil generator, reproducibility | 5 | Passed |
| Total | Current Go unit-test suite | 26 | Passed |

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
### Benchmarking Methodology

All end-to-end measurements were performed on the same Apple M1 computer with 16 GiB of memory, running macOS 26.6.2. The software environment used Go 1.27.1 for `darwin/arm64` and R 4.3.2. Both implementations processed the same 1,000 saved observations and performed 10,000 bootstrap repetitions.

Both programs were executed in batch mode from the command line. R was run using `Rscript`, while Go was run as a compiled executable. The macOS `/usr/bin/time -l` utility measured elapsed time and maximum resident set size using the same method for both implementations.

### End-to-End R and Go Comparison

Both implementations were executed 100 times in batch mode using the same computer, input dataset, and bootstrap repetition count. R was executed with `Rscript`, while Go was run as a compiled executable. The macOS `/usr/bin/time -l` utility measured elapsed time and maximum resident set size.

| Measurement | R | Go |
|---|---:|---:|
| Runs | 100 | 100 |
| Mean elapsed time | 0.5571 s | 0.5038 s |
| Median elapsed time | 0.5500 s | 0.5000 s |
| Runtime standard deviation | 0.0137 s | 0.0167 s |
| Minimum elapsed time | 0.5400 s | 0.4900 s |
| Maximum elapsed time | 0.6100 s | 0.5700 s |
| Mean maximum resident set size | 153.94 MiB | 10.44 MiB |
| Median maximum resident set size | 154.59 MiB | 10.42 MiB |

Based on mean elapsed time, compiled Go was approximately 1.106 times as fast as R, representing a 9.57% runtime reduction. Go used approximately 93.22% less resident memory, while R used approximately 14.74 times as much resident memory.

The first isolated R run took 0.83 seconds, which was substantially slower than the repeated-run mean. This likely reflected cold-start package or filesystem effects. The 100-run results were used for the primary comparison to avoid basing conclusions on one potentially unusual execution.

## Profiling and Logging

The Go application uses the standard-library `log` package to record major execution events. It logs application startup, the input file, the number of observations loaded, the bootstrap seed and repetition count, successful completion of the calculation, and application completion. Fatal input or calculation errors are also reported through the logger.

Individual bootstrap repetitions are not logged because 10,000 log messages would produce excessive output and distort performance measurements.

Logs are written to standard error, while statistical results are written to standard output. Both streams were captured in `results/logs/go_application.log` using:

```bash
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

### After-Optimization Profiles

The after-optimization allocation profile confirmed that the copy previously created by `calculateMedian` had been removed from the sequential bootstrap path. Before optimization, `calculateMedian` allocated 79.62 MB and `bootstrapSample` allocated 77.60 MB. After optimization, `calculateMedianInPlace` had no significant allocation, while `bootstrapSample` allocated 78.11 MB.

The after-optimization CPU profile showed that sorting accounted for approximately 72.50% of cumulative CPU samples. Bootstrap sample generation accounted for approximately 10%. This confirmed that sorting remained the primary CPU bottleneck after the redundant slice copy was removed.

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

Google Cloud Compute Engine was selected as the infrastructure-as-a-service provider. The hypothetical deployment uses an `e2-standard-2` virtual machine in the Iowa `us-central1` region. This machine provides 2 virtual CPUs and 8 GiB of memory. Its default on-demand price was $0.06701142 per hour when accessed on November 15, 2026.

Source: https://cloud.google.com/products/compute/pricing/general-purpose

The cost estimate applies the locally measured mean runtimes to a continuous batch workload of one million executions.

| Measurement | R | Go |
|---|---:|---:|
| Mean runtime per execution | 0.5571 s | 0.5038 s |
| Runtime for 1,000,000 executions | 154.75 hours | 139.94 hours |
| Estimated compute cost | $10.37 | $9.38 |
| Estimated savings with Go | — | $0.99 |
| Percentage savings | — | 9.57% |

This estimate assumes that the relative performance measured on the Apple M1 transfers to the selected cloud VM. It is a projection rather than a cloud-hosted benchmark. The estimate excludes persistent storage, network transfer, taxes, discounts, engineering labor, and maintenance costs. It also assumes the executions are processed continuously on an already-running VM. Launching a separate VM for every subsecond execution would invoke Google Cloud's minimum billing period and would not produce the same cost estimate.

Go also used 93.22% less resident memory locally. This may allow denser workloads or smaller VM configurations, but that potential benefit was not converted directly into cost savings because both implementations fit within the selected VM's memory.

## Recommendation

The consultancy should not replace R with Go for all statistical work. R remains preferable for exploratory analysis, rapid model development, visualization, and methods already supported by its extensive package ecosystem.

Go should be considered for stable, frequently executed production workloads where the statistical method is already understood and runtime or memory use creates meaningful infrastructure costs. For this bootstrap experiment, Go produced a 9.57% mean runtime reduction and used 93.22% less resident memory than R. The projected compute-cost reduction was therefore approximately 9.57% for a continuous batch workload on the same VM configuration.

The financial benefit is small for infrequent executions. At one million executions, the projected savings were approximately $0.99 on the selected VM. A migration would make sense only when the workload is much larger, memory pressure is important, or the Go implementation can be reused across production systems.

This study evaluated one statistical method, one symmetric dataset, one computer, and a sequential Go implementation. It did not directly benchmark cloud hardware or multi-core scaling. The consultancy should conduct a cloud-hosted pilot using representative production data before making a broader migration decision.

## Installation and Usage

### Prerequisites

- Go 1.27.1 or compatible
- R 4.3.2 or compatible
- R package `boot`

Clone and enter the repository:

```bash
git clone https://github.com/gummybeardream/bootstrap-r-go.git
cd bootstrap-r-go
```

Run the R implementation:

```bash
Rscript r/bootstrap_median.R
```

Run the Go implementation:

```bash
go run ./go
```

Run the Go tests:

```bash
go test -count=1 -v ./go
```

Run the Go benchmarks:

```bash
go test -run '^$' -bench . -benchmem -count 5 ./go
```

Install the R package if needed:

```bash
Rscript -e 'install.packages("boot", repos="https://cloud.r-project.org")'
```

## Building the Go Executable

From the repository root, compile the application with:

```bash
go build -o bootstrap-median ./go
```

Run the compiled executable using:

```bash
./bootstrap-median
```

## GenAI Tools

I used OpenAI Codex as a programming tutor and review assistant during this project. It helped to:

- break the assignment into milestones
- explain R and Go packages for bootstrap estimation
- review and debug code

I reviewed, entered, tested, and revised the project code and documentation. AI-generated suggestions were validated with Go formatting, unit tests, vetting, benchmarks, and manual application runs.

## References

- Gelman, A., & Vehtari, A. (2021). *What are the most important statistical ideas of the past 50 years?* https://arxiv.org/abs/2012.00174
- R Core Team. *R: A Language and Environment for Statistical Computing.* https://www.r-project.org/
- Canty, A., & Ripley, B. *boot: Bootstrap Functions.* https://cran.r-project.org/package=boot
- Voss, J. *The Mersenne Twister in Go.* https://github.com/seehuhn/mt19937
- Google Cloud. *General Purpose VM Pricing.* https://cloud.google.com/products/compute/pricing/general-purpose
- Go Project. *Go Standard Library.* https://pkg.go.dev/std