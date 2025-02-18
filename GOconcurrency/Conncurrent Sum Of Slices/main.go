package main

import (
	"fmt"
	"sync"
)

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	numRoutines := 2

	chunkSize := (len(numbers) + numRoutines - 1) / numRoutines
	var partialSums []int

	var wg sync.WaitGroup
	//split the numbers into chunks and process each chunk concurrently
	for i := 0; i < numRoutines; i++ {
		start := i * chunkSize
		end := start + chunkSize
		if end > len(numbers) {
			end = len(numbers)
		}
		chunk := numbers[start:end]

		wg.Add(1)
	
		go func(chunk []int) {
			defer wg.Done()

			sum := 0
			for _, num := range chunk {
				sum += num
			}

			partialSums = append(partialSums, sum)
		}(chunk)
	}

	wg.Wait()
	
	//compute the final sum by combining the partial sums
	finalSum := 0
	for _, sum := range partialSums {
		finalSum += sum
	}

	fmt.Printf("The total sum is: %d\n", finalSum)
}
