package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // Continuously listen for jobs in jobs channel
		// Simulate work with sleep
		time.Sleep(time.Second)
		// Print the worker and the job it's processing
		fmt.Printf("Worker %d processed job %d\n", id, job)
		// Send result back to the results channel
		results <- job * 2 // Result is the job number multiplied by 2
	}
}

func main() {
	// Create channels
	jobs := make(chan int, 5)    // jobs channel for sending jobs to workers
	results := make(chan int, 5) // results channel for collecting results

	// Start 3 worker goroutines
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs to the jobs channel
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs) // Close the jobs channel to signal workers

	// Collect results from the results channel (fan-in pattern)
	for i := 1; i <= 5; i++ {
		fmt.Println("Received:", <-results) // Read results from results channel
	}

	close(results) // Close results channel when done
}
