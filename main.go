package main

import (
	"errors"
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func validateAmount(amount float64) error {

	if amount < 0 {
		return errors.New("value cannot be negative")
	}

	if amount == 0 {
		return errors.New("amount cannot be 0")
	}

	return nil
}

func main() {

	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("------------")
	}

	fmt.Println("Welcome to devyank's bank:")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())

	for {
		presentOptions()

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			fmt.Println("Your balance is:", accountBalance)
		}

		if choice == 2 {
			fmt.Println("How much you want to deposit: ")
			var deposit float64
			fmt.Scan(&deposit)

			err := validateAmount(deposit)

			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}

			accountBalance += deposit
			fmt.Println("Your balance now is:", accountBalance)
			fileops.WriteFloatToFIle(accountBalance, accountBalanceFile)
		}

		if choice == 3 {
			fmt.Println("How much you want to withdraw: ")
			var withdraw float64
			fmt.Scan(&withdraw)

			err := validateAmount(withdraw)
			if err != nil {
				fmt.Println("ERROR:", err)
				continue
			}

			if withdraw > accountBalance {
				fmt.Println("You can't withdraw this value!")
			} else {
				accountBalance -= withdraw
				fmt.Println("You withdrew:", withdraw, "and now your balance is:", accountBalance)
				fileops.WriteFloatToFIle(accountBalance, accountBalanceFile)
			}

		}

		if choice == 4 {
			fmt.Println("Goodbye!")
			break
		}

		if choice > 4 {
			fmt.Println("Please choose a real option!")
		}
	}

	fmt.Println("Thanks for choosing our bank")
}
