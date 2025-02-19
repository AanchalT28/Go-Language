package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	warehouse := make(chan string)

	wg.Add(3)
	go producer(warehouse, "Producer 1", &wg)
	go producer(warehouse, "Producer 2", &wg)
	go producer(warehouse, "Producer 3", &wg)

	go consumer(warehouse, "Consumer 1")
	go consumer(warehouse, "Consumer 2")

	wg.Wait()

	close(warehouse)

	// Wait for consumers to finish processing
	// The consumers will keep consuming until the channel is closed.
}

func producer(ch chan string, name string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		item := fmt.Sprintf("%s Item %d", name, i)
		fmt.Println(name, "produced", item)
		ch <- item // Send the item to the channel
	}
}

func consumer(ch chan string, name string) {
	for item := range ch {
		fmt.Println(name, "consumed", item)
	}
}
