package main

import (
	"fmt"
)

func main() {
	fmt.Println("Working with Collections")

	var customerNames[3] string
	customerNames[0] = "Kalidoss"
	customerNames[1] = "Malairaj"
	customerNames[2] = "Raja"

	fmt.Printf("Customer Names: %s\n",customerNames)

	customerAge := []int{10,20,30,4} 

	fmt.Printf("Customer Ages: %q\n",customerAge)
	customerAge = append(customerAge,40)

	fmt.Println(customerAge)
	
	customerLocations := [...]string{"Chennai","Madurai","Ramnad","Rameshwaram"}
	fmt.Printf("%q",customerLocations)

//2D Array

var a[2][10]string

for i := 0;i<2;i++{
	for j :=0; j<10;j++ {
		a[i][j] = fmt.Sprintf("%d,%d",i,j)
	}
}

fmt.Printf("%q\n",a)




}