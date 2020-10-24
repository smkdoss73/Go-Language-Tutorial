package main

import (
	"fmt"
)

func main(){

	numbers := []int{10,20,30,4,6,7,8,9,5,4,6,7,8}
	for i, v := range numbers {
		fmt.Println(i,v)
	}

	marks := make([]int, 10)

	fmt.Println("***********")
	for i,_ := range marks {
		marks[i] = i
	}
	fmt.Println(marks)

	//Map

	cities := map[string]int {
		"Madurai" : 100,
		"Ramnad" : 200,
		"Chennai": 300,
	}

	for key, value := range cities {
		fmt.Println(key,value)
	}
}