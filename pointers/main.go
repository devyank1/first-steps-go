package main

import "fmt"

func main() {

	number := 100
	resetToZero(&number)
	fmt.Println(number)

	value := 19.99
	duplicateValue(&value)
	fmt.Println(value)

	totalOne := 100.0
	applyDiscount(&totalOne, true)
	fmt.Println(totalOne)

	totalTwo := 100.0
	applyDiscount(&totalTwo, false)
	fmt.Println(totalTwo)
}

func resetToZero(n *int) {
	*n = 0
}

func duplicateValue(n *float64) {

	*n = *n * 2
}

func applyDiscount(price *float64, isMember bool) {

	if isMember {
		*price = *price * 0.9
	}
}
