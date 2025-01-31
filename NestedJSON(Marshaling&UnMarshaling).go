package main

import (
	"encoding/json"
	"fmt"
)

type Address struct {
	Street  string `json:"street"`
	City    string `json:"city"`
	ZipCode string `json:"zip_code"`
}

type Person struct {
	FirstName string  `json:"first_name"`
	LastName  string  `json:"last_name"`
	Age       int     `json:"age"`
	Address   Address `json:"address"` // Nested struct
}

func main() {
	// Unmarshalling JSON into Go struct
	jsonStr := `{
		"first_name": "Aanchal,",
		"last_name": "T,",
		"age": 21,
		"address": {
			"street": "123 Pl St,",
			"city": "Pune,",
			"zip_code": "41107"
		}
	}`

	var p Person
	// Unmarshalling JSON into the Person struct
	err := json.Unmarshal([]byte(jsonStr), &p)
	if err != nil {
		fmt.Println("Error unmarshalling:", err)
		return
	}

	fmt.Println("Unmarshalled Go struct:", p)

	// Marshalling Go struct back to JSON
	jsonData, err := json.Marshal(p)
	if err != nil {
		fmt.Println("Error marshalling:", err)
		return
	}

	fmt.Println("Marshalled JSON:", string(jsonData))
}
