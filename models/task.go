package models

import (
	"fmt"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Text        string    `json:"text"`
	Status      string    `json:"status"`
	CreatedTime time.Time `json:"created_time"`
}

var LAYOUT = "2006-01-02 15:04:05"

func CreateTask(text []string, TaskList []Task) Task {
	newTask := Task{
		ID:          GetNextID(TaskList),
		Text:        strings.Join(text, " "),
		Status:      "process",
		CreatedTime: time.Now(),
	}

	return newTask
}

func PrintTask(task Task) {
	fmt.Println(task.ID, task.Text, task.Status, task.CreatedTime.Format(LAYOUT))
}

func GetNextID(TaskList []Task) int {
	if len(TaskList) == 0 {
		return 1
	} else {
		max := -1

		for _, elem := range TaskList {
			if elem.ID > max {
				max = elem.ID
			}
		}

		return max + 1
	}

}
