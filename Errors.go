package main

import (
	"errors"
	"fmt"
)

type CustomError struct {
	Errorcode int
	Error     error
}

func (err *CustomError) ShowError() string {
	return fmt.Sprintf("Error Code: %d, Error Details: %s", err.Errorcode, err.Error)
}

func getDemoError() error {
	return errors.New("Test Error")
}

func main() {
	fmt.Println("Working with Error")
	fmt.Println(getDemoError())
	ce := &CustomError{
		Errorcode: 123,
		Error:     errors.New("Invalid username or password!"),
	}
	fmt.Println(ce.ShowError())
}
