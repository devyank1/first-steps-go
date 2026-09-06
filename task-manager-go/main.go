package main

import (
	"fmt"
	"time"
)

func main() {
	tm := &TaskManager{}

	taskOne, err := NewTask("Write report", "Finish the Q3 report", time.Now().AddDate(0, 0, 3))

	if err != nil {
		fmt.Println("error creating task", err)
		return
	}

	tm.AddTask(taskOne)

	taskTwo, err := NewTask("Buy groceries", "Milk, eggs, bread", time.Now().AddDate(0, 0, 1))

	if err != nil {
		fmt.Println("error creatin task", err)
		return
	}

	tm.AddTask(taskTwo)

	fmt.Println("--------ALL TASKS--------")
	tm.ListTasks()

	if err := tm.CompleteTask(taskOne.id); err != nil {
		fmt.Println("error completing task:", err)
	}

	fmt.Println("--------AFTER COMPLETING TASK ONE--------")
	tm.ListTasks()

	if err := tm.DeleteTask(taskTwo.id); err != nil {
		fmt.Println("error deleting task:", err)
	}

	fmt.Println("--------AFTER DELETING TASK TWO--------")
	tm.ListTasks()

	if err := tm.CompleteTask("does-not-exist"); err != nil {
		fmt.Println("expected error:", err)
	}
}
