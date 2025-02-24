//sync/atomic:
//Provides atomic operations like AddInt64, LoadInt64, and StoreInt64 for safely modifying variables without the need for locking.
//Ideal for simple types and operations, where fine-grained control over atomicity is needed.

package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func main() {
	var counter int64 = 0 // Use int64 for atomic operations

	var wg sync.WaitGroup

	// Start multiple goroutines to increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1) // Atomically increment counter
		}()
	}

	wg.Wait()
	fmt.Println("Final Counter:", counter) // Print the final counter value
}
