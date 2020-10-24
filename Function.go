package main

import "fmt"



func main() {
	customername, mobileNumber := getCustomerInfo()
	fmt.Println(customername,mobileNumber)
	printWelcomeMessage()
	printMessageWithCustomMessage("Life is an empty dream!")
	fmt.Println(getStudentInfomationsByID(90708132032))
}


func getCustomerInfo() (customername string, mobile int) {

	var cname = "Kalidoss"
	var cMobile = 9566241073

	return cname, cMobile
}

func printWelcomeMessage() {
	fmt.Println("Welcome to all!")
}

func printMessageWithCustomMessage(message string) {
	fmt.Println(message)
}

func getStudentInfomationsByID(studentID int) (string) {
	 studentName := fmt.Sprintf("Student Name is Kalidoss id is %d",studentID)
	 return studentName
}