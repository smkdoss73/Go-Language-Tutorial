package main

import (
	"fmt"
	"time"
)

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	fmt.Println("Working with Concurrency...!")
	go say("World")
	say("Hello")

}