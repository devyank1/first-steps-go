package functions

import "fmt"

type CType interface {
	int | int32 | uint | uint32
}

func sum1[T CType](A T, B T) T {
	return A + B
}

func sum2[T CType](Numbers ...T) T {
	var res T

	for _, num := range Numbers {
		res += num
	}

	return res
}

func main() {
	numbers := []uint{1, 2, 3, 4, 5}

	fmt.Println(sum2(numbers...))
}
