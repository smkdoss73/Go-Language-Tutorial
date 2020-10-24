package main

import (
	"fmt"
	"log"
)

func PrintNames(names []string) {

	fmt.Println("Fav Names:%s", names)
	//Panic
	fmt.Println("Fav Names:%s index: %d", names[len(names)])
	//Output
	/*
			panic: runtime error: index out of range [4] with length 4

		goroutine 1 [running]:
		main.PrintNames(0xc000070040, 0x4, 0x4)
		        /Users/admin/Documents/Workspace/Go/oct23/Panics.go:8 +0xc6
		main.main()
		        /Users/admin/Documents/Workspace/Go/oct23/Panics.go:19 +0x9c
		exit status 2
	*/
}

func CustomPanic(panicMessage string) {

	panic(panicMessage)

	//Output
	/*
			panic: Oh no!!!!!!!!

		goroutine 1 [running]:
		main.CustomPanic(...)
		        /Users/admin/Documents/Workspace/Go/oct23/panics.go:24
		main.main()
		        /Users/admin/Documents/Workspace/Go/oct23/panics.go:37 +0x50
		exit status 2
	*/
}

func HandleWithPanic() {
	fmt.Println("Handling with panic start-->")
	defer func() {
		if err := recover(); err != nil {
			log.Println("Panic accured")
		}
	}()
	CustomPanic("123")
}

func main() {

	// fmt.Println("Working with panic")
	// //Type 1
	// names := []string{
	// 	"Kaidoss",
	// 	"Malairaj",
	// 	"Akash",
	// 	"Nisha",
	// }
	// PrintNames(names)

	// //Type 2
	// CustomPanic("Oh no!!!!!!!!")

	println("How to handle with panic")

	HandleWithPanic()

	//Output
	/*
		How to handle with panic
		Handling with panic start-->
		2020/10/23 09:58:33 Panic accured
	*/

}
