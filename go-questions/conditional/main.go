package conditional

import (
	"fmt"
	"math/rand"
)

func main() {

	randNumber := rand.Intn(101)

	if randNumber == 50 {
		fmt.Println("It's 50!!")
		fmt.Println(randNumber)

	} else if randNumber > 50 && randNumber%2 == 0 {
		fmt.Println("It's closer to 100 and even")
		fmt.Println(randNumber)

	} else if randNumber > 50 && randNumber%2 != 0 {
		fmt.Println("It's closer to 100 and odd")
		fmt.Println(randNumber)

	} else if randNumber < 50 && randNumber%2 == 0 {
		fmt.Println("It's closer to 0 and even")
		fmt.Println(randNumber)

	} else {
		fmt.Println("It's closer to 0 and odd")
		fmt.Println(randNumber)
	}
}
