package main

import (
	"fmt"
	"sync"
	"time"
)

// Job represents a task to be processed
type Job struct {
	ID int
}

// Worker function that processes jobs
func Worker(id int, jobs <-chan Job, wg *sync.WaitGroup) {
	defer wg.Done() // Mark as done when the worker is finished
	for job := range jobs {
		// Simulate processing time
		fmt.Printf("Worker %d is processing job %d\n", id, job.ID)
		time.Sleep(1 * time.Second) // Simulating time taken to process a job
	}
}

func main() {
	// Number of workers and jobs
	workerCount := 3
	jobCount := 5

	// Channel to send jobs to workers
	jobQueue := make(chan Job, jobCount)

	// WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Start worker goroutines
	for i := 1; i <= workerCount; i++ {
		wg.Add(1)
		go Worker(i, jobQueue, &wg) // Start a worker
	}

	// Generate and send jobs to the jobQueue
	for i := 1; i <= jobCount; i++ {
		jobQueue <- Job{ID: i}
	}

	// Close the jobQueue channel once all jobs are sent
	close(jobQueue)

	// Wait for all workers to finish processing jobs
	wg.Wait()

	fmt.Println("All jobs have been processed!")
}
