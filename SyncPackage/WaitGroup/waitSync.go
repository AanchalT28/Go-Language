package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go helloworld(&wg)
	go bye(&wg)
	wg.Wait()
}

func helloworld(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("HelloW World")
}

func bye(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Goodbye")
}
