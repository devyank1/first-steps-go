package continue1

import "fmt"

func main() {

	sum := 0

	for i := range 100 {
		if i == 50 {
			continue
		}

		sum += i
	}

	fmt.Println(sum)
}
