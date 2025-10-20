package models

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int
	Text        string
	Status      string
	CreatedTime time.Time
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

// добавить ошибки если в строке меньше параметров, а потом вообще на json поменять

func CreateTaskFromFileString(lineString string) Task {
	if lineString == "\n" || lineString == "" {
		return Task{}
	}

	params := strings.Split(lineString, " ")

	idTask, _ := strconv.Atoi(params[0])

	timeTask, _ := time.Parse(LAYOUT, params[len(params)-2]+" "+params[len(params)-1])

	return Task{ID: idTask, Text: strings.Join(params[1:len(params)-3], " "), Status: params[len(params)-3], CreatedTime: timeTask}
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
