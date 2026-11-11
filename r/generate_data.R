set.seed(2025)

observation_count <- 1000
values <- rnorm(observation_count)
dataset <- data.frame(value = values)

write.csv(dataset, "data/symmetric.csv", row.names = FALSE)
