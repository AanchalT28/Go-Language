package main

import (
	"fmt"
)

func main() {
	numbers := []int{1, 2, 3, 4, 2, 5, 6, 3, 7, 8, 8}

	//create a map to store the frequency of each number
	countMap := make(map[int]int)

	//count the frequency of each number in the slice
	for _, num := range numbers {
		countMap[num]++
	}

	fmt.Println("Unique numbers:")
	for num, count := range countMap {
		if count == 1 {
			fmt.Println(num)
		}
	}

	fmt.Println("\nDuplicate numbers:")
	for num, count := range countMap {
		if count > 1 {
			fmt.Println(num)
		}
	}
}
