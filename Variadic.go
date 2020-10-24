package main

import "fmt"

func DisplayNames(names ...string) {
	for _, n := range names {
		fmt.Printf("%s\n", n)
	}
}

func main() {
	fmt.Println("Working with Variadic function")
	DisplayNames("Kali")
	DisplayNames("Kali", "Doss")
	DisplayNames("Kali", "Doss", "Malairaj", "Akash")
	DisplayNames("KD")
}
