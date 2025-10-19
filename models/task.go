package models

import (
	"fmt"
	"os"
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

func CreateTask(text string) Task {
	newTask := Task{
		ID:          1,
		Text:        text,
		Status:      "process",
		CreatedTime: time.Now(),
	}

	return newTask
}

func PrintTask(task Task) {
	fmt.Println(task.ID, task.Text, task.Status, task.CreatedTime.Format(LAYOUT))
}

func WriteTaskToFile(task Task, filename string) {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		fmt.Print("error")
		return
	}

	strTask := []byte(strconv.Itoa(task.ID) + " " + task.Text +
		" " + task.Status + " " + task.CreatedTime.Format(time.DateTime) + "\n")

	file.Write(strTask)

	defer file.Close()

}

// добавить ошибки если в строке меньше параметров, а потом вообще на json поменять

func CreateTaskFromFileString(lineString string) Task {
	params := strings.Split(lineString, " ")

	idTask, _ := strconv.Atoi(params[0])

	timeTask, _ := time.Parse(LAYOUT, params[3]+" "+params[4])

	return Task{ID: idTask, Text: params[1], Status: params[2], CreatedTime: timeTask}
}
