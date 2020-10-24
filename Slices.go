package main

import "fmt"

func main() {

	fmt.Println("Working with Slice")

	mySlice := []int{10,20,30,40,60,70,80}

	myArray := []int{1,2,3,4,5,6,7,8,9,10}

	myArray = append(myArray,10,23,2334,3545,46546,4646)
	mySlice = append(mySlice,100,200,300,400,500,600)

	
	fmt.Println(mySlice[1:])
	// fmt.Println(myArray[1:5])
	// fmt.Println(myArray[:4])
	// fmt.Println(myArray[4:])
	fmt.Println(len(myArray))
	fmt.Println(cap(myArray))

	fmt.Println(cap([]int{10,20}))




	//Making Slice

	cities := make([]string,10)
	cities[0] = "Chennai"
	cities[1] = "Madurai"
	cities[2] = "Ramnad"

//Append



	fmt.Println(cities)
}