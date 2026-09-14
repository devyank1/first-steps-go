package lists

import "fmt"

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have - OK
// 		Output (print) that array in the command line. - OK
// 2) Also output more data about that array: - OK
//		- The first element (standalone) - OK
//		- The second and third element combined as a new list - OK
// 3) Create a slice based on the first element that contains
//		the first and second elements. - OK
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals) - OK
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array - OK
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products). - OK
//		Then add a third product to the existing list of products.

type Product struct {
	id    int64
	title string
	price float64
}

func main() {
	hobbies := [3]string{"run", "basketball", "play guitar"}
	fmt.Println(hobbies)

	fmt.Println("Standalone:", hobbies[0])
	fmt.Println("Combined as a new list:", hobbies[1:3])

	newSlice := hobbies[:2]
	fmt.Println(newSlice)

	fmt.Println(cap(newSlice))
	newSlice = newSlice[1:3]
	fmt.Println(newSlice)

	courseGoals := []string{"Learn GO", "Work with GO"}
	fmt.Println(courseGoals)

	courseGoals[1] = "Work and save money with GO"
	anotherGoal := append(courseGoals, "Live with my family close and happy :)")
	fmt.Println(anotherGoal)

	productList := []Product{
		{
			01,
			"Laptop",
			1399.90,
		},
		{
			02,
			"Mouse Wireless",
			599.99,
		},
	}

	fmt.Println(productList)

	newProduct := []Product{
		{
			01,
			"Mousepad",
			89.99,
		},
	}

	fmt.Println(newProduct)

	productList = append(productList, newProduct...)

	fmt.Println(productList)
}

//func main() {
//	prices := []float64{10, 8.99}
//	fmt.Println(prices)
//
//	prices[1] = 9.99
//
//	updatedPrices := append(prices, 5.99)
//	fmt.Println(prices, updatedPrices)
//}

//func main() {
//
//	var productNames = [4]string{"A Book"}
//	prices := [5]float64{10.99, 9.99, 45.99, 20.00, 4.49}
//
//	fmt.Println(prices[3])
//
//	fmt.Println(prices)
//	fmt.Println(productNames)
//
//	featuredPrices := prices[1:3]
//	fmt.Println(featuredPrices)
//
//	fmt.Println(cap(featuredPrices))
//
//}
