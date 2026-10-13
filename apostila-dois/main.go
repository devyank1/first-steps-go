package main

import (
	"errors"
	"fmt"
)

//func double(number *int) {
//	*number *= 2
//}

//func main() {
//	dedup := 20
//	double(&dedup)
//	fmt.Println(dedup)
//}

//func switchValues(x, y *int) {
//	var c int
//
//	c = *x
//	*x = *y
//	*y = c
//}
//
//func main() {
//	x := 10
//	y := 2
//	switchValues(&x, &y)
//
//	fmt.Println(x, y)
//}

//func discount(price *float64, percentage float64) error {
//
//	if *price < 0 {
//		return errors.New("you should input a real value")
//	} else if percentage < 0 || percentage > 100 {
//		return errors.New("you should input a real percentage")
//	}
//
//	*price = *price - (*price * (percentage / 100))
//
//	return nil
//}
//
//func main() {
//
//	price := 100.0
//	discountPercentage := 100.0
//
//	if err := discount(&price, discountPercentage); err != nil {
//		fmt.Println("Error", err)
//		return
//	}
//
//	fmt.Println(price)
//}

func cart() {

}

func addProduct(total *float64, items *int, price float64, quantity int) error {

	if quantity <= 0 || price <= 0 {
		return errors.New("Invalid price or quantity")
	}

	*total += price * float64(quantity)
	*items += quantity

	return nil
}

func applyCoupon(used bool, total *float64, code string) error {

	if used == true {
		errors.New("This coupon was already used")
	}

	if code != "GOPHER10" {
		errors.New("You should put a valid coupon")
	}

	if used == false {
		*total = *total * 0.1
	}

	return nil
}

func main() {
	var total float64
	var items int
	var couponUsed bool

	for {
		fmt.Println(`
Welcome to Devyank's Supermarket
1. Add Product
2. Apply Coupon
3. Check Cart
4. Proceed to Checkout`)

		var option int
		fmt.Print("Choose: ")
		fmt.Scan(&option)

		switch option {
		case 1:
			var price float64
			var qty int

			fmt.Print("Price: ")
			fmt.Scan(&price)
			fmt.Print("Quantity: ")
			fmt.Scan(&qty)

			if err := addProduct(&total, &items, price, qty); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Product added!")
			}

		case 2:
			var code string
			fmt.Print("Coupon code: ")
			fmt.Scan(&code)

			if err := applyCoupon(&total, &couponUsed, code); err != nil {
				fmt.Println("Error:", err)
			} else {
				fmt.Println("Coupon applied!")
			}

		case 3:
			status := "não"
			if couponUsed {
				status = "sim"
			}
			fmt.Printf("%d itens · Total: R$ %.2f · Cupom: %s\n", items, total, status)

		case 4:
			if total == 0 {
				fmt.Println("O carrinho está vazio")
				continue
			}
			fmt.Printf("Total final: R$ %.2f\n", total)
			return

		default:
			fmt.Println("Invalid option")
		}
	}

}
