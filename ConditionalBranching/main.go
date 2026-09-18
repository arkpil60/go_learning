package main

import "fmt"

var score int = 5
var year int = 400

func main() {

	whatsGrade()
	leapYear()

}

// First Task
func whatsGrade() {

	switch score {
	case 1:
		fmt.Println("Very Bad")
	case 2:
		fmt.Println("Bad")
	case 3:
		fmt.Println("Okay")
	case 4:
		fmt.Println("Good")
	case 5:
		fmt.Println("Very Good")
	default:
		fmt.Println("No Grade")
	}

}

// Second Task
func leapYear() {

	if year%4 == 0 && year%100 != 0 {
		fmt.Println("Leap Year")
	} else if year%400 == 0 && year%100 == 0 {
		fmt.Println("Leap Year")
	} else {
		fmt.Println("Not Leap Year")
	}

}
