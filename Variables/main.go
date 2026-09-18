package main

import "fmt"

func main() {

	myCustomTask()

}

func myCustomTask() {

	//Calculate the total number of liters of fuel required for the entire trip and the cost.

	var distance float64 = 250.5
	var consumption = 8.5
	pricePerLiter := 60
	var currency string = "RUB"

	var requiredFuelQuantity float64 = (consumption / 100) * distance
	fmt.Printf("Required Fuel Quantity: %.1f liters\n", requiredFuelQuantity)

	var totalFuelCost float64 = float64(pricePerLiter) * requiredFuelQuantity
	fmt.Printf("Total fuel cost: %.2f %s\n", totalFuelCost, currency)

}
