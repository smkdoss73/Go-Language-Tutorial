package main

import "fmt"

func Sum(values []int, c chan int) {
	sum := 0
	for _, v := range values {
		sum += v
	}
	c <- sum
}
func main() {
	fmt.Println("Working with Channels")

	// a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// c := make(chan int, 0)

	first := make(chan int, 2)
	first <- 10
	first <- 20
	// go Sum(a, c)
	// x := <-c
	// fmt.Println(x)
	fmt.Println(first)
	fmt.Println(first)

	///Fibbonacci

	cc := make(chan int, 10)
	go fibonacci(cap(cc), cc)
	for i := range cc {
		fmt.Println(i)
	}

}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}
