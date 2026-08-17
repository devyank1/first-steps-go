package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFIle(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile("balance.txt", []byte(balanceText), 0644)
}

func main() {

	fmt.Println("Welcome to devyank's bank:")

	for {

		var accountBalance float64 = 1000

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

			if deposit <= 0 {
				fmt.Println("Your value should be greater than 0")
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
