package main

import "fmt"

func main() {

	number1 := 10
	text1 := "hello"
	float1 := 1.1
	boolean1 := true

	fmt.Println("number1:", number1)
	fmt.Println("text1:", text1)
	fmt.Println("float1", float1)
	fmt.Println("boolean1:", boolean1)

	var number2 int = 20
	var text2 string = "world"
	var float2 float64 = 2.2
	var boolean2 bool = false

	fmt.Println("number2:", number2)
	fmt.Println("text2:", text2)
	fmt.Println("float2", float2)
	fmt.Println("boolean2:", boolean2)

	MyCustomTask()

}

func MyCustomTask() {

	//Calculate the total number of liters of fuel required for the entire trip and the cost.

	var distance float64 = 250.5
	var consumption float64 = 8.5
	var pricePerLiter int = 60
	var currency string = "RUB"

	var requiredFuelQuantity float64 = (consumption / 100) * distance
	fmt.Printf("Required Fuel Quantity: %.1f liters\n", requiredFuelQuantity)

	var totalFuelCost float64 = float64(pricePerLiter) * requiredFuelQuantity
	fmt.Printf("Total fuel cost: %.2f %s\n", totalFuelCost, currency)

}
