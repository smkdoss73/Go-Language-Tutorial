package main

import (
	"fmt"
	"strings"
)

func main() {
	customerInfo := map[string]string{
		"name":       "Kalidoss",
		"department": "computer science and engg",
		"title":      "CSE",
	}

	//fmt.Printf("%#v\n", customerInfo)
	var key string = "   name  "
	key = strings.TrimSpace(key)
	if name, status := customerInfo[key]; status {
		fmt.Println(name)
	} else {
		fmt.Println("Key is not available")
	}
}
