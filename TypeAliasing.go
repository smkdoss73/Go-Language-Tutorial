package main

import (
	"fmt"
	"strings"
)

type SMKString string
type SMKFloat float64

func main() {

	c := strings.ToUpper(string(SMKString("hello dear friend")))

	f := SMKFloat(100.0)

	fmt.Println(c)
	fmt.Println(f)
	fmt.Sprintln(f)
}
