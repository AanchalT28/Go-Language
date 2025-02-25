package main

import (
	"context"
	"fmt"
	"time"
)

func longRunningTask(ctx context.Context) {
	select {
	case <-time.After(5 * time.Second): //simulate a long running task
		fmt.Println("Task completed successfully.")
	case <-ctx.Done(): //If the context is cancelled, exit
		fmt.Println("Task cancelled: ", ctx.Err())
	}
}

func main() {
	//create a context with 3 second timeout
	ctx, cancel := context.WithTimeout(context.Background(), 9*time.Second)
	defer cancel() //cancel the context to avoid leaking resources

	//call long running task with the context
	fmt.Println("Starting task..")
	go longRunningTask(ctx)

	//wait fot task to complete or timeout
	select {
	case <-ctx.Done():
		fmt.Println("Main recieved context cancellation: ", ctx.Err())

	}
}
