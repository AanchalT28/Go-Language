package main

import (
	"fmt"
	"sync"
	//"time"
)

func main() {

	var wg sync.WaitGroup

	//create buffered channel with cap 3
	warehouse := make(chan string, 3)

	//start 3 producer goroutines
	wg.Add(3)
	go producer(warehouse, "Producer 1", &wg)
	go producer(warehouse, "Producer 2", &wg)
	go producer(warehouse, "Producer 3", &wg)

	//start 2 consumer goroutines
	wg.Add(2)
	go consumer(warehouse, "Consumer 1", &wg)
	go consumer(warehouse, "Consumer 2", &wg)

	//allow goroutines to work for a while
	//time.Sleep(5 * time.Second)

	wg.Wait()
}

func producer(ch chan string, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 5; i++ {
		item := fmt.Sprintf("%s Item %d", name, i)
		fmt.Println(name, "produced", item)

		//add timer to produce the items
		//time.Sleep(1 * time.Second)

		ch <- item
	}

	if name == "Producer 3" {
		close(ch)
	}
}

func consumer(ch chan string, name string, wg *sync.WaitGroup) {
	defer wg.Done()

	//recieve item from warehouse
	for item := range ch {
		fmt.Println(name, "consumed", item)
	}

	//add timer so consumer to take time to consume item
	//time.Sleep(1 * time.Second)
}
