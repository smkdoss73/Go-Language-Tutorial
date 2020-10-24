package main

import (
	"fmt"
	"log"
	"strconv"
)

func main() {

	fmt.Println("Working with Convert data type")

	// float to int & String

	var priceRange float64 = 10.45

	fmt.Println(int(priceRange))
	fmt.Println(strconv.Itoa(int(priceRange)) + " Kalidoss " + fmt.Sprint(priceRange) + " Testing ")

	//String to number

	var salary string = "10008953464687"

	salaryInt, error := strconv.Atoi(salary)

	if error == nil {
		fmt.Println(salaryInt)
	} else {
		log.Fatal(error)
	}

}
