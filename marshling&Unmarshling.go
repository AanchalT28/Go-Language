package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
}

func main() {
	jsonStr := `{
		"first_name": "Aanchal",
		"last_name": "T",
		"age": 21
	}`

	var p Person
	// Unmarshalling JSON into the Person struct
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}
	fmt.Println("Unmarshalled:", p)

	// marshalling struct back to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	fmt.Println("Marshalled JSON:", string(jsonData))
}
