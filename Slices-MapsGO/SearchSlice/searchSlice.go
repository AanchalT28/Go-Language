package main

import (
	"fmt"
)

func main() {

	stringSlice := []string{"a", "b", "c", "d", "cc", "ab", "bd"}
	stringMap := map[string]bool{
		"a":  true,
		"b":  true,
		"c":  true,
		"d":  true,
		"cc": false,
		"ab": true,
		"bd": false,
	}

	var val string
	fmt.Println("Enter the value you want to search:")
	fmt.Scanln(&val)

	//searching in slice
	foundInSlice := false //initialized flag to check whether true or false
	for i, v := range stringSlice {
		if v == val {
			foundInSlice = true
			fmt.Println("Element", val, "found in slice at pos:", i)
			break
		}
	}

	if foundInSlice == false {
		fmt.Println("Not found in slice:", val)
	}

	//searching in map
	_, foundInMap := stringMap[val]

	if foundInMap {
		fmt.Println("Found in map:", val)
	} else {
		fmt.Println("Not found in map:", val)
	}

}
