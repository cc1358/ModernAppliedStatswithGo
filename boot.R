# Install and load the boot package
install.packages("boot")
library(boot)

# Generate some example data
set.seed(123)
data <- rnorm(100)

# Define a function to calculate the statistic of interest (e.g., mean)
mean_func <- function(data, indices) {
  return(mean(data[indices]))
}

# Start measuring CPU time
start_time <- system.time()

# Perform bootstrapping
boot_results <- boot(data, mean_func, R = 100)

# Stop measuring CPU time and print the elapsed time
end_time <- system.time({})
cpu_time <- end_time - start_time
print(cpu_time)

# View bootstrap results
print(boot_results)

# Convert times to microseconds
cpu_time_micro <- cpu_time * 1e6
print(cpu_time_micro)
