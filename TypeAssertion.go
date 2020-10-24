

package main

import (
		"fmt"
		"time"
)


func timeMap(y interface{}) {
	z, ok := y.(map[string]interface{})
	if ok {
		z["updated_time"] = time.Now()
		z["customer_type"] = "Pro Member"
	}

}


type Stringer interface {
	String() string
}

type fakeString struct {
	content string
}

func (s *fakeString) String() string {
	return s.content
}

func printString(value interface{}) {
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())		
	}

}


func main() {

	

	fmt.Println("Type Assertion Example")

	foo := map[string]interface{} {
		"HIMS_Customer_Name": "Kalidoss",
		"HIMS_Customer_Age": 29,
	}

	timeMap(foo)
	fmt.Println(foo)

fmt.Println("*************************")

s := &fakeString{"Life is an empty dream"}
printString(s)
printString("Hello world")

}


