package main

import "fmt"

func main() {

	//for slices the index starts from 0 but when we slice it it only till slices end-1
	stringSlice := []string{"Yellow", "Red", "Blue", "Black", "Purple", "Orange"}
	fmt.Println("Original Slice: ", stringSlice)

	var start, end int

	fmt.Println("Enter the start value(index) starting from 0:")
	_, err := fmt.Scanln(&start)
	if err != nil {
		fmt.Println("Error value", err)
		return
	}

	fmt.Println("Enter the end value(index) from 0-", len(stringSlice), ":")
	_, err1 := fmt.Scanln(&end)
	if err1 != nil {
		fmt.Println("Error value: ", err)
		return
	}

	//check if valid index
	if start < 0 || end > len(stringSlice) || start > end {
		fmt.Println("Invalid index range")
		return
	}

	sliced := stringSlice[start:end]
	fmt.Println("The sliced string is :", sliced)

}
