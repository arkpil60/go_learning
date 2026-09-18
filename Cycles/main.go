package main

import (
	"fmt"
	"math/rand"
)

func main() {

	defaultIfElse(10)
	myCustomTask()
	sumTask()
	hp(100)

}

// A simple if-else example
func defaultIfElse(score int) {
	if score == 10 {
		fmt.Println("=10")
	} else if score > 10 {
		fmt.Println(">10")
	} else {
		fmt.Println("<10")
	}

}

func myCustomTask() {

	//Calculate how long it will take to accumulate a certain amount.

	target := 150000
	currentSavings := 20000
	monthlyInvest := 15000
	monthCount := 0

	for currentSavings < 150000 {
		currentSavings = currentSavings + monthlyInvest
		monthCount++
		if currentSavings >= target {
			fmt.Println("Good, You've saved up", currentSavings)
		} else {
			continue
		}
	}

}

// First Task
func sumTask() {
	sum := 0
	for i := 1; i <= 50; i++ {
		sum = sum + i
	}
	fmt.Println("Answer:", sum)
}

// Second Task
func hp(hp int) {

	for hp > 0 {
		damage := rand.Intn(15)
		hp = hp - damage

		fmt.Printf("HP lost: %d. Remaining HP: %d\n", damage, hp)

		if hp <= 0 {
			fmt.Println("You died.")
		}
	}

}
