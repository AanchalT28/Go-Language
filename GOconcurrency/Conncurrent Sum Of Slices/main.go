package main

import (
	"fmt"
	"sync"
)

func main() {
	// Input slice of integers
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// Number of Go routines
	numRoutines := 2 // You can choose to change this value based on the input size or any other condition

	// Calculate chunk size
	chunkSize := (len(numbers) + numRoutines - 1) / numRoutines // This ensures even distribution

	// To store the partial sums from each Go routine
	var partialSums []int

	// WaitGroup to wait for all Go routines to complete
	var wg sync.WaitGroup

	// Split the numbers into chunks and process each chunk concurrently
	for i := 0; i < numRoutines; i++ {
		// Calculate the start and end indices of the current chunk
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}

		// Ensure that the chunk is properly sliced
		chunk := numbers[start:end]

		// Increment the WaitGroup counter
		wg.Add(1)

		// Launch a Go routine to process the current chunk
		go func(chunk []int) {
			defer wg.Done() // Decrement the WaitGroup counter when done

			// Compute the sum of the chunk
			sum := 0
			for _, num := range chunk {
				sum += num
			}

			// Append the sum to the partialSums slice
			partialSums = append(partialSums, sum)
		}(chunk)
	}

	// Wait for all Go routines to complete
	wg.Wait()

	// Compute the final sum by combining the partial sums
	finalSum := 0
	for _, sum := range partialSums {
		finalSum += sum
	}

	// Output the result
	fmt.Printf("The total sum is: %d\n", finalSum)
}
