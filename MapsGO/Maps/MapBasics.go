package main

import "fmt"

func main() {

	//var a= map[keytype]valuetype{key:value}
	var cars = map[string]string{"Brand": "BMW", "Price": "50Lkh"}
	fmt.Println(cars)

	//using make func
	var fruits = make(map[string]string) //empty map initialized
	fruits["Yellow"] = "Banana"
	fruits["Red"] = "Apple"
	fmt.Println(fruits)

	//access map elements
	fmt.Println(fruits["Red"])

	//update or add elements
	fruits["Red"] = "strawberry"
	fruits["Purple"] = "Grapes"
	fmt.Println(fruits)

	//removing an element from map
	delete(fruits, "Yellow")
	fmt.Println(fruits)

	//check for an existing key and value
	val1, ok1 := fruits["Red"]
	//check for only existing key
	_, ok2 := fruits["Purple"]

	fmt.Println(val1, ok1)
	fmt.Println(ok2)
}
