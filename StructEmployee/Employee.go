package main

import "fmt"

type Employee struct {
	name        string
	CompanyName string
	Salary      float64
}

// func updateSalary(emp *Employee, newslary float64) {
// 	emp.Salary = newslary
// }

func main() {
	emp1 := Employee{"Nina", "Siemens", 10000}

	fmt.Println("Employee1:", emp1)

	// updateSalary(&emp1, 20000)
	// fmt.Println("Employee1 :", emp1)

}
