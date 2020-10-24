package main

import "fmt"


type User struct {
	FirstName, LastName, Email string
	Location *UserLocation
}

func (user User) ShowWelcomeMessage() string {
	return fmt.Sprintf("Welcome Mr. %s",user.FirstName)
} 

func (user *User) ShowMessage() string {
	return fmt.Sprintf("Customer Location is  %s", user.Location.State)
}


type UserLocation struct {
	City string
	State string
	Country string
}

func newUser(firstname, lastname, email, city,state,country string) *User {
	return &User{FirstName: firstname, LastName: lastname,Email: email,
			Location: &UserLocation{City: city, State: state, Country: country},}
}

func main() {
	user := newUser("Kalidoss","S","smkdoss73@gmail.com","Chennai","Tamilnadu","India")
	fmt.Println(user.ShowWelcomeMessage())
	fmt.Println(user.ShowMessage())
}