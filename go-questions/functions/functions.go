package functions

import "fmt"

func add(x, y float64) float64 {
	return x + y
}

func main() {

	result := add(10, 5)

	fmt.Println(result)
}
