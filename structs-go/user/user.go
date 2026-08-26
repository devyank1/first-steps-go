package user

import (
	"errors"
	"fmt"
	"time"
)

// data model like a "class" we can define the shape of a struct with the details
type User struct {
	firstName   string
	lastName    string
	birthdate   string
	isCreatedAt time.Time
}

type Admin struct {
	admEmail    string
	admPassword string
	User
}

func (u *User) OutputDetails() {
	fmt.Println(u.firstName, u.lastName, u.birthdate)
}

func (u *User) ClearUserName() {
	u.firstName = ""
	u.lastName = ""
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		admEmail:    email,
		admPassword: password,
		User: User{
			firstName:   "ADMIN",
			lastName:    "ADMIN",
			birthdate:   "---",
			isCreatedAt: time.Now(),
		},
	}
}

// you can call this constructor many times you want, if you want to create a user with newUser
func New(firstName, lastName, birthdate string) (*User, error) {

	if firstName == "" || lastName == "" || birthdate == "" {
		return nil, errors.New("All fields are required")
	}

	return &User{
		firstName:   firstName,
		lastName:    lastName,
		birthdate:   birthdate,
		isCreatedAt: time.Now(),
	}, nil
}
