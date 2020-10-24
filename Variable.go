package main
import "fmt"

func main() {

	/*
	Type 1
	*/
var customerName string = "Kalidoss"

fmt.Printf("%T",customerName)
fmt.Println(customerName)
fmt.Println(add(10,20))

/*
Type 2
*/

name, location := "Prince Oberyn", "Dorne"
age := 32
fmt.Printf("%s (%d) of %s\n", name, age, location)

/*
Type 3
*/

var(
	studentName string
	studentAge int
	) 

studentName = "Kali doss"
studentAge = 31
displayMessage := fmt.Sprintf("Student age %d and Name is %s",studentAge,studentName)
fmt.Println(displayMessage)

/*
Type 4
*/
var (
	cardNumber, cardName, cardDate = 123456, "Credit Card", "22/05/1991"
)
fmt.Println(cardNumber,cardName,cardDate)
}

func add(firstNumber int,secondNumber int) int {

return firstNumber + secondNumber

}




