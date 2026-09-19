package main

import (
	"fmt"
)

func main() {

	message := greet("Martin")
	fmt.Println(message)

	divide, multiplication := stats(3.0, 5.0)
	fmt.Printf("Result of divide: %.2f.\n", divide)
	fmt.Printf("Result of multiplication: %.2f.\n", multiplication)

}

func greet(name string) string {

	return fmt.Sprintf("Hello, %s!", name)

}

func stats(a, b float64) (float64, float64) {

	return a / b, a * b

}
