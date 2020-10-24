package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Enter your name")

	var customerName string

	fmt.Scanln(&customerName)
	fmt.Printf("Hello %s\n", strings.ToUpper(customerName))
	//scount := strings.Count(customerName, "k")
	sLength := len(customerName)
	fmt.Println(sLength)
}
