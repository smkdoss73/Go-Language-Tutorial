package main

import "fmt"

type Bootcamp struct {
	Lat float64
	Long float64
}
func main() {
	
	x := new(Bootcamp)
	x.Lat = 10.00
	x.Long = 93.999
	
	y := &Bootcamp{10.00,93.999}


	fmt.Println(*x == *y)
}




