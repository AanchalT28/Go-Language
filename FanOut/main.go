package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	for job := range jobs { // Continuously listen on jobs channel
		// Simulate some work by sleeping for 1 second
		time.Sleep(time.Second)
		fmt.Printf("Worker %d processed job %d\n", id, job)
		results <- job * 2 // The result is the job multiplied by 2
	}
}

func main() {
	jobs := make(chan int, 10)    // jobs channel that will hold the jobs
	results := make(chan int, 10) // results channel that will hold the results

	// Start 3 worker goroutines to process jobs
	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results)
	}

	// Send jobs to the jobs channel
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	// Collect results from the results channel
	for i := 1; i <= 5; i++ {
		<-results
	}
	close(results)
}
