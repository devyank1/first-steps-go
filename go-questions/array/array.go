package array

import "fmt"

func array() {

	emptyArray := [10]int{}

	fmt.Println(emptyArray)

	stringArray := [4]string{"one", "two", "three", "four"}

	fmt.Println(stringArray)

	movies := [4]string{"movie1", "movie2", "movie3", "movie4"}

	fmt.Println(movies)
	fmt.Println(movies[1])
	fmt.Println(len(movies))
}
