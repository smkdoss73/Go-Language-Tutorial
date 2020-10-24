package main

import (
	"fmt"
)

type User struct {
	FirstName, LastName string
}

type Namer interface {
	Name() string
}

func (u *User) Name() string {
	return fmt.Sprintf("Customer fullname is %s %s", u.FirstName, u.LastName)
}

func Greet(n Namer) string {
	return fmt.Sprintf(n.Name())
}

func main() {
	fmt.Println("Working with Interfaces")

	userInfo := &User{"Kali", "Doss"}
	userInfo.Name()
	fmt.Println(userInfo.Name()) // call via methods
	fmt.Println(Greet(userInfo)) // call via interface

}
