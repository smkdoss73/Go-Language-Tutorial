package main

import "fmt"

func init() {
	fmt.Print("Init called \n")
}
func init() {
	fmt.Println("Second init")
}
func init() {
	fmt.Println("Third init")
}
func main() {
	fmt.Println("Working with multiple init()")
}
