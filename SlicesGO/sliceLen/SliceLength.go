package main

import "fmt"

func main() {
	// Define a slice with some initial elements
	slice := []string{"Apple", "Banana", "Cherry", "Date"}

	// Check the length of the slice
	fmt.Println("The length of the slice is:", len(slice))
	fmt.Println(cap(slice))
}