package main

import (
	"fmt"
	"structs-example/user"
)

func main() {
	userFirstName := getUserData("Please enter your first name: ")
	userLastName := getUserData("Please enter your last name: ")
	userBirthDate := getUserData("Please enter your birthdate (DD/MM/YYYY): ")

	var appUser *user.User

	appUser, err := user.New(userFirstName, userLastName, userBirthDate)

	if err != nil {
		fmt.Println(err)
		return
	}

	admin := user.NewAdmin("yan@email.com", "pwd123")

	admin.OutputDetails()
	admin.ClearUserName()
	admin.OutputDetails()

	appUser.OutputDetails()
	appUser.ClearUserName()
	appUser.OutputDetails()
}

func getUserData(promptText string) string {
	fmt.Println(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
