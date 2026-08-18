package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)

	if err != nil {
		return 0, errors.New("Failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 0, errors.New("Failed to parse stored balance values")
	}

	return balance, nil
}

func writeBalanceToFIle(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

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

	var accountBalance, err = getBalanceFromFile()

	if err != nil {
		fmt.Println("ERROR:", err)
		fmt.Println("------------")
	}

	fmt.Println("Welcome to devyank's bank:")

	for {

		fmt.Println("What do you want to do?")
		fmt.Println("1. Check balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")

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
			writeBalanceToFIle(accountBalance)
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
				writeBalanceToFIle(accountBalance)
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
