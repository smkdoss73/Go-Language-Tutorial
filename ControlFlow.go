package main

import (
	"fmt"
)

func main() {

	 a := 10
	 b := 20
	if a > b {
		fmt.Println("A is Big")
	} else {
		fmt.Println("B is Big")
	}

	if c := true; c {
		fmt.Println("C is Positive")
	} else {
		fmt.Println("C is Negative")
	}

	//For Loop

	for i :=0;i<10;i++ {
		fmt.Println(i)
	}

	//Switch
cases := 3

	switch cases {
	case 1:

		fmt.Println("First Case")
		break
	case 2:
		fmt.Println("Second Case")
		break
	default:
		fmt.Println("Default cases")
		break
	}
}