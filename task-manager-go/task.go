package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
)

type Task struct {
	id          string
	name        string
	description string
	dueDate     time.Time
	completed   bool
}

type TaskManager struct {
	tasks []Task
}

func (tm *TaskManager) AddTask(t Task) {
	tm.tasks = append(tm.tasks, t)
}

func NewTask(name string, description string, dueDate time.Time) (Task, error) {

	if name == "" || description == "" {
		return Task{}, errors.New("invalid input")
	}

	newID := uuid.New().String()

	return Task{
		id:          newID,
		name:        name,
		description: description,
		dueDate:     dueDate,
		completed:   false,
	}, nil
}

func (tm *TaskManager) ListTasks() {
	for _, t := range tm.tasks {
		fmt.Println("name:", t.name)
		fmt.Println("description:", t.description)
		fmt.Println("due date:", t.dueDate)
		fmt.Println("task status:", t.completed)
	}
}

func (tm *TaskManager) CompleteTask(id string) error {

	for i := range tm.tasks {
		if tm.tasks[i].id == id {
			tm.tasks[i].completed = true
			return nil
		}
	}
	return errors.New("task not found")
}

func (tm *TaskManager) DeleteTask(id string) error {

	for i := range tm.tasks {
		if tm.tasks[i].id == id {
			tm.tasks = append(tm.tasks[:i], tm.tasks[i+1:]...)
			return nil
		}
	}
	return errors.New("we cannot delete your task")
}

//func (task *Task) FindTask(id string)  error {
//
//		if task.id == id {
//			fmt.Println(task.id)
//			fmt.Println(task.name)
//			fmt.Println(task.description)
//			fmt.Println(task.dueDate)
//			fmt.Println(task.completed)
//
//			return nil
//		}
//	return errors.New("task not found")
//}
