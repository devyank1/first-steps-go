package pt01

import (
	"fmt"
)

//func avg(grade []float64) float64 {
//
//	if len(grade) == 0 {
//		return 0
//	}
//
//	sum := 0.0
//
//	for _, g := range grade {
//		sum += g
//	}
//
//	return sum / float64(len(grade))
//}
//
//func main() {
//	fmt.Println(avg([]float64{7, 8, 9}))
//	fmt.Println(avg([]float64{10, 5}))
//	fmt.Println(avg([]float64{}))
//}

//func evenFilter(numbers []int) []int {
//
//	var evenValues []int
//
//	for _, n := range numbers {
//
//		if n%2 == 0 {
//			evenValues = append(evenValues, n)
//		}
//	}
//
//	return evenValues
//}
//
//func main() {
//
//	fmt.Println(evenFilter([]int{1, 2, 3, 4, 5, 6, 7, 8}))
//	fmt.Println(evenFilter([]int{-2}))
//
//}

func invertValue(value []int) {

	start := 0
	end := len(value) - 1

	for start < end {
		value[start], value[end] = value[end], value[start]
		start++
		end--
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	invertValue(nums)
	fmt.Println(nums)
}
