package main

import "fmt"

func main() {

	process()

}

func process() {

	defer fmt.Println("Start")

	defer fmt.Println("Step 1")

	defer fmt.Println("Step 2")

	fmt.Println("Execution of the function...")

}
