package main

import "fmt"

var customerName string

func main() {
	fmt.Println("Working with Global Variable")
	customerName = "Raja"

	fmt.Println(getlocalName())

	fmt.Println("This is Global Variable", customerName)
}

func getlocalName() string {
	name := "Kalidoss - Local variable"

	customerName = "Raja raja solan"

	return name
}
