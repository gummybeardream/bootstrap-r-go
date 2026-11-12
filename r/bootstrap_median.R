library(boot)

dataset <- read.csv("data/symmetric.csv")

# Validate that the value column exists 
if (!"value" %in% names(dataset)) {
    stop("The input dataset must contain a column named 'value'.")
}
dataset_values <- dataset$value

median_function <- function(values, indices) {
    return(median(values[indices]))
}

bootstrap_seed <- 2026
bootstrap_reps <- 10000

set.seed(bootstrap_seed)
results <- boot(
    data = dataset_values,
    statistic = median_function,
    R = bootstrap_reps
)

bootstrap_standard_error <- sd(results$t[, 1])

cat("Observation count:", length(dataset_values), "\n")
cat("Bootstrap seed:", bootstrap_seed, "\n")
cat("Bootstrap repetitions:", bootstrap_reps, "\n")
cat("Original median:", results$t0, "\n")
cat("Bootstrap standard error:", bootstrap_standard_error, "\n")

