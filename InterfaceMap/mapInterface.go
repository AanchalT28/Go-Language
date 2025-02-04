package main

import "fmt"

func main() {

	data := make(map[string]interface{})

	//adding key value to the map
	data["name"] = "Sam"
	data["age"] = 30
	data["Student"] = false
	data["grades"] = []int{30, 29, 26, 28}

	fmt.Println(data)

	//type assertion if we want to print individual values
	name := data["name"].(string) //type assertion to string
	//age := data["age"].(int)
	Student := data["Student"].(bool)
	grades := data["grades"].([]int)

	fmt.Println(name)
	//fmt.Println(age)
	fmt.Println(Student)
	fmt.Println(grades)

	//safe type assertion incase the value does not match and not to cause panic
	if age, ok := data["age"].(int); ok {
		fmt.Println("Age:", age)
	} else {
		fmt.Println("Age is not int ")
	}

}
