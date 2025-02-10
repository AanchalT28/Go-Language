package main

import "fmt"

func causePanic() {
	panic("Something went wrong!")
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	fmt.Println("Before panic")
	causePanic()               //this will trigger a panic
	fmt.Println("After panic") //this won't be executed
}
