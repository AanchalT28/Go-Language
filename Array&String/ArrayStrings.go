package main

import "fmt"

func main() {

	// var n int
	// fmt.Println("Enter the number of input: ")
	// fmt.Scan(&n)

	// //create slice to input the values
	// var fruits = make([]string,n)

	// //loop to get user input for each string
	// for i:=0; i<n; i++{
	// 	fmt.Printf("Enter the string %d:",i+1)
	// 	fmt.Scan(&fruits[i])
	// }

	// fmt.Println("The array of fruits: ",fruits)

	var fruits = [5]string{"Apple", "Orange", "Pineapple", "Blueberry", "Banana"}

	//Print the length of the array
	fmt.Println("The length of the array is: ", len(fruits))
}
