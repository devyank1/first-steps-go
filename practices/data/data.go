package data

import (
	"fmt"
	"practices/inputs"
)

func GetNoteData() (string, string, error) {

	title, err := inputs.GetUserInput("Note Title:")

	if err != nil {
		fmt.Println("ERROR:", err)
		return "", "", err
	}

	content, err := inputs.GetUserInput("Note Content:")

	if err != nil {
		fmt.Println("ERROR:", err)
		return "", "", err
	}

	return title, content, nil
}
