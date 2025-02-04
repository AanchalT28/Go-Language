package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	const (
		header = "Location,Size,Beds,Baths,Price"
		data   = `New York,100,2,1,100000
New York,150,3,2,200000
Paris,200,4,3,400000
Istanbul,500,10,5,1000000`
		separator = ","
	)

	//1.Split the header
	headercolumn := strings.Split(header, separator)
	for _, column := range headercolumn {
		fmt.Printf("%-15s", column)
	}
	fmt.Println()

	fmt.Println("===============================================================================")

	//2: Split data into rows
	rows := strings.Split(data, "\n")

	//3: Initialize slices for the columns
	var locations []string
	var sizes []int
	var beds []int
	var baths []int
	var prices []int

	//4:parse each row and load data into slices
	//this loop iterates over each row of data("New York,100,2,1,100000")
	for _, row := range rows {

		columns := strings.Split(row, separator)

		locations = append(locations, columns[0])

		//strconv.Atoi is used to convert the string integers to integers, Atoi is ASCII o integers
		size, _ := strconv.Atoi(columns[1])
		sizes = append(sizes, size)

		bed, _ := strconv.Atoi(columns[2])
		beds = append(beds, bed)

		bath, _ := strconv.Atoi(columns[3])
		baths = append(baths, bath)

		price, _ := strconv.Atoi(columns[4])
		prices = append(prices, price)
	}

	for i := 0; i < len(locations); i++ {
		fmt.Printf("%-15s%-15d%-15d%-15d%-15d\n", locations[i], sizes[i], beds[i], baths[i], prices[i])
	}
}
