/*

create go files in below path

└── $GOPATH
    └── src
        └── github.com
            └── gopherguides
                └── example
                    └── example.go



	Program(example.go):

	package example
	import "fmt"

	func SayHello() {
		fmt.Println("Hello! Good Morning!")
	}


	//Main Project( using example package)

	package main

	import(
		"github.com/gopherguides/example"
	)

	func main(){
		example.SayHello()
	}



*/