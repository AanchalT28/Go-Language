//sync.RWMutex:
//Provides separate locks for reading and writing, enabling concurrent reads but exclusive writes.
//Used when you need multiple goroutines to read concurrently but only one to write at a time.

package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	mu   sync.RWMutex
	data int
}

func (c *Counter) Increment() {
	c.mu.Lock() //write lock
	defer c.mu.Unlock()
	c.data++ //update a shared resource
}

func (c *Counter) Read() int {
	c.mu.RLock() //read only lock
	defer c.mu.RUnlock()
	return c.data //read the shared resource
}

func main() {

	counter := &Counter{} //create a counter instance

	var wg sync.WaitGroup //add wait group to wait for goroutines to finish

	//start multiple goroutines to increment the counter
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()     //decrease the counter when goroutine finishes
			counter.Increment() //increment the counter
		}()
	}

	//start multiple routines to read the counter
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(counter.Read(), "Goroutine No:", i)
		}()
	}

	wg.Wait()
	fmt.Println("Final counter: ", counter.Read())
}
