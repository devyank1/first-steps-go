package inputs

import (
	"errors"
	"fmt"
)

func GetUserInput(input string) (string, error) {

	fmt.Println(input)
	var valueText string
	fmt.Scanln(valueText)

	if valueText == "" {
		return "", errors.New("You should input a text to proceed")
	}

	return valueText, nil
}
