package main

import (
	"fmt"
	"practices/data"
	"practices/inputs"
)

func main() {
	title, content, err := inputs.GetUserInput()

	if err != nil {
		fmt.Println(err)
		return
	}

	data.GetNoteData()
}
