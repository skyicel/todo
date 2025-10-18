package models

import (
	"fmt"
	"time"
)

type Task struct {
	ID          int
	Text        string
	Status      string
	CreatedTime time.Time
}

func CreateTask(text string) Task {
	newTask := Task{
		ID:          1,
		Text:        text,
		Status:      "in process",
		CreatedTime: time.Now(),
	}

	return newTask
}

func PrintTask(task Task) {
	fmt.Println(task.ID, task.Text, task.Status, task.CreatedTime)
}
