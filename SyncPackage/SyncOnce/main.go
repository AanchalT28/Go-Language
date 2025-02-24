//sync.Once:
//Ensures a function is executed only once, no matter how many times it is called from multiple goroutines.
//Useful for initialization tasks that should only be done once, even if multiple goroutines try to perform them.

package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func initialize() {
	fmt.Println("Initialization performed!")
}

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("Goroutine %d calling Do()\n", i)
			once.Do(initialize)
		}(i)
	}

	wg.Wait()
}
