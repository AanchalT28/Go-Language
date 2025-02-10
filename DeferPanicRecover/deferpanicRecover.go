package main

import "fmt"

func causePanic() {
	panic("Something went wrong!")
}

func main() {
	// Defer the recovery logic directly within the main function
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	causePanic()               // This will trigger a panic
	fmt.Println("After panic") // This won't be executed
}
