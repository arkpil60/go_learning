package main

import (
	"fmt"
)

var count = 10

func main() {

	reset(&count)
	fmt.Println(count)

	a := 5
	b := 10
	swapPointers(&a, &b)
	fmt.Println(a, b)

}

func reset(n *int) {
	*n = 0
}

func swapPointers(x, y *int) {
	*x, *y = *y, *x
}
