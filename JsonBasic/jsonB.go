package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name  string `json:"Full_name"` //key to place for Name value
	Age   int
	Hobby string
}

func main() {

	me := Person{
		Name:  "Aanchal",
		Age:   20,
		Hobby: "Skating",
	}

	me2 := `{
		"Full_name":"Aanchal",   		
		"Age":21,
		"Hobby":"Sketching"
	 }` //instead of name i used Full_name as for json key value which we provided
	fmt.Println(me2)
	var me2info Person

	err := json.Unmarshal([]byte(me2), &me2info) //changes json to a struct
	if err != nil {
		fmt.Println("Unable to UnMarshal")
	}
	fmt.Println(me2info)

	structJson, err := json.Marshal(me) //changes struct to json
	if err != nil {
		fmt.Println("Unable to json")
	}
	fmt.Println(string(structJson))
}
