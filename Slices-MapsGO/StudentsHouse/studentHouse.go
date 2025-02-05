package main

import (
	"fmt"
	"sort"
)

func main() {

	houses := make(map[string][]string)

	//adding data to map
	houses["gryffindor"] = []string{"weasley", "hagrid", "dumbledore", "lupin"}
	houses["hufflepuf"] = []string{"wenlock", "scamander", "helga", "diggory"}
	houses["ravenclaw"] = []string{"flitwick", "bagnold", "wildsmith", "montmorency"}
	houses["slytherin"] = []string{"horace", "nigellus", "higgs", "scorpius"}
	houses["bobo"] = []string{"wizardry", "unwanted"}

	//retrive students by house
	var houseName string
	fmt.Println("Enter the house name:")
	fmt.Scanln(&houseName)

	if houseName == "" {
		fmt.Println("House name cannot be empty.")
		return
	}
	//sort students by name
	if students, found := houses[houseName]; found {
		studentsCopy := append([]string{}, students...) //not modify but copy the slice
		sort.Strings(studentsCopy)

		fmt.Printf("Students in house %s:\n", houseName)
		for _, student := range studentsCopy {
			fmt.Println(student)
		}
	} else {
		fmt.Println("House not found!")
	}

	fmt.Println("Bobo will be deleted as it is not a part of houses")
	fmt.Println("\n")
	delete(houses, "bobo")

	fmt.Println("\nUpdated list of houses and students:")
	for house, students := range houses {

		studentsCopy := append([]string{}, students...)
		sort.Strings(studentsCopy)

		fmt.Printf("%s: %v\n", house, studentsCopy)

	}

}
