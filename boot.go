package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

func main() {
	// Create a new random generator with a time-based seed
	randomGenerator := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate example data
	n := 100
	data := make([]float64, n)
	for i := 0; i < n; i++ {
		data[i] = randomGenerator.NormFloat64()
	}

	// Start measuring CPU time
	startTime := time.Now()

	// Define the number of bootstrap resamples
	numResamples := 100

	// Perform bootstrapping with parallelization
	resampledMeans := bootstrapParallel(data, numResamples, randomGenerator)

	// Calculate confidence interval
	sort.Float64s(resampledMeans)
	alpha := 0.05
	lowerIdx := int(float64(numResamples) * alpha / 2)
	upperIdx := int(float64(numResamples) * (1 - alpha/2))
	if upperIdx >= numResamples {
		upperIdx = numResamples - 1
	}
	confidenceInterval := [2]float64{resampledMeans[lowerIdx], resampledMeans[upperIdx]}

	// Stop measuring CPU time
	endTime := time.Now()

	// Calculate and print CPU time
	cpuTime := endTime.Sub(startTime)
	fmt.Printf("CPU Time: %v\n", cpuTime)

	// Print the results
	fmt.Printf("Bootstrap Confidence Interval (95%%): [%f, %f]\n", confidenceInterval[0], confidenceInterval[1])
}

// Function to perform bootstrapping with parallelization
func bootstrapParallel(data []float64, numResamples int, randomGenerator *rand.Rand) []float64 {
	var wg sync.WaitGroup
	resampledMeans := make([]float64, numResamples)
	for i := 0; i < numResamples; i++ {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resampledData := resample(data, randomGenerator)
			resampledMeans[index] = mean(resampledData)
		}(i)
	}
	wg.Wait()
	return resampledMeans
}

// Function to perform resampling
func resample(data []float64, randomGenerator *rand.Rand) []float64 {
	n := len(data)
	resampledData := make([]float64, n)
	for i := 0; i < n; i++ {
		resampledData[i] = data[randomGenerator.Intn(n)]
	}
	return resampledData
}

// Function to calculate the mean
func mean(data []float64) float64 {
	sum := 0.0
	for _, v := range data {
		sum += v
	}
	return sum / float64(len(data))
}
