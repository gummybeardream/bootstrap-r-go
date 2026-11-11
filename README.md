# Bootstrap Performance Comparison: R and Go
Go and R bootstrap sampling comparison

## Project Overview

## Management Problem

## Selected Statistical Method

The selected statistical method is the bootstrap estimation of the median's standard error. The median is the middle value after sorting all observations. Standard error measures how much an estimate would vary across repeated samples. Bootstrap resampling approximates how an estimate can vary across repeated samples. Bootstrap resampling replicates this variation by repeatedly sampling from the observed data with replacement. This method is computationally intensive because it may repeat the calculation thousands of times. 

## Bootstrap Methodology

## R Package Research

`boot`: https://cran.r-project.org/package=boot

`boot manual`: https://cran.r-project.org/web/packages/boot/boot.pdf

`R median`: https://stat.ethz.ch/R-manual/R-devel/library/stats/html/median.html

## Go Package Research

## Repository Structure

## Data and Experimental Design

R and go will read the same saved observations from the shared dataset, `data/symmetric.csv`

- Distribution: normal and symmetric
- Observation count: 1,000
- Data-generation seed: 2025
- Shared file: data/symmetric.csv
- Planned bootstrap repetitions: 10,000
- R version: 4.3.2
- `boot` package version 1.3.28.1


## R Implementation

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