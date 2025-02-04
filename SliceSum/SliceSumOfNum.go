package main

import "fmt"

func calcSum(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total = total + number
	}
	return total
}

func main() {

	sum1 := calcSum(1, 2, 3, 4)
	fmt.Println("Sum1:", sum1)
	nums := []int{20, 30, 50}
	sum2 := calcSum(nums...)
	fmt.Println("Sum2:", sum2)
}
