package pt1

import (
	"errors"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func converter(c float64) float64 {
	var f float64

	f = (c*9)/5 + 32

	return f
}

func ex1() {

	fmt.Printf("Your degree is: %.2f", converter(36.6))
}

func totalTickets(quantity int) float64 {

	const ticketPrice float64 = 32.5

	payment := ticketPrice * float64(quantity)

	return payment
}

func ex() {

	fmt.Println(totalTickets(5))
}

func visitCard(name string, age int, height float64) string {

	return fmt.Sprintf(`===========
Name: %s
Age: %d,
Height: %.2f
===========`, name, age, height)
}

func ex3() {

	fmt.Println(visitCard("Yan", 26, 1.83))
}

func netSalary(salary float64) (float64, float64) {

	if salary > 0 && salary <= 2000 {
		return 0, salary
	} else if salary <= 5000 {
		return salary * 0.1, salary
	} else {
		return salary * 0.3, salary
	}
}

func ex4() {
	fmt.Println(netSalary(10000))
}

func calculator(a, b float64, op string) (float64, error) {

	var result float64

	switch op {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		if b == 0 {
			errors.New("You can't divide by 0.")
		} else {
			result = a / b
		}
	default:
		return 0, fmt.Errorf("Invalid value %s", op)
	}

	return result, nil
}

func ex5() {

	fmt.Println(calculator(10, 20, "+"))
	fmt.Println(calculator(10, 20, "-"))
	fmt.Println(calculator(10, 20, "*"))
	fmt.Println(calculator(10, 20, "/"))
}

func ex6() {

	var number = rand.Intn(100)
	var guessNum int
	var guesses int

	for {
		fmt.Print("Type your guess: ")
		fmt.Scan(&guessNum)
		if number > guessNum {
			fmt.Println("Number is higher than your guess")
			guesses++
		} else if number < guessNum {
			fmt.Println("Number is lower than your guess")
			guesses++
		} else {
			guesses++
			fmt.Println("YOU GOT IT!! Attempts:", guesses)
			continue
		}
	}
}

func validate(pwd string) error {

	if len(pwd) < 8 {
		return errors.New("Your password is too short")
	} else if pwd == "12345678" {
		return errors.New("Your password is too easy")
	} else {
		fmt.Println("Your password has been stored")
	}

	return nil
}

func ex7() {
	fmt.Println(validate("123456789"))
}

func loadExpenses() float64 {
	if _, err := os.Stat(documentTxt); os.IsNotExist(err) {
		fmt.Println("File doesn't found")
		return 0.0
	}

	content, err := os.ReadFile(documentTxt)

	if err != nil {
		fmt.Println("Error by reading information", err)
	}

	text := strings.TrimSpace(string(content))
	if text == "" {
		return 0.0
	}

	value, err := strconv.ParseFloat(text, 64)
	if err != nil {
		fmt.Println("Convert error.", err)
		return 0.0
	}

	return value
}

func saveExpense(value float64) {

	text := fmt.Sprintf("%.2f", value)

	err := os.WriteFile(documentTxt, []byte(text), 0644)
	if err != nil {
		fmt.Println("We couldn't save your expenses", err)
	}
}

const documentTxt = "document.txt"

func main() {
	const monthlyGoal float64 = 3000

	totalExpense := loadExpenses()

	var input int

	for {

		fmt.Println(`=== Expense Control ===
1. Record Expense
2. Check Total Expense
3. Check Total Credits
4. Check how much is left to reach the goal
5. Exit`)

		fmt.Println("Choose:")
		fmt.Scan(&input)

		switch input {
		case 1:
			fmt.Println("Type your expense value (R$):")
			var newExpense float64
			_, err := fmt.Scan(&newExpense)
			if err != nil || newExpense <= 0 {
				fmt.Println("You cannot put 0 as a value")
				continue
			}

			totalExpense += newExpense
			saveExpense(totalExpense)
			fmt.Println("Expense stored successfully")
		case 2:
			fmt.Println("Expense total at this moment:", totalExpense)
		case 3:
			fmt.Printf("Your total credits are: %.2f", monthlyGoal-totalExpense)
		case 4:
			if totalExpense <= monthlyGoal {
				fmt.Println("You have ", monthlyGoal-totalExpense, "credits")
				fmt.Printf("Your percentage is %.2f", (totalExpense/monthlyGoal)*100, "%")
			} else {
				fmt.Println("You beat your goal")
			}
		case 5:
			fmt.Println("Thank you, bye")
			return
		default:
			fmt.Println("Invalid input, try a new one")
		}
	}
}
