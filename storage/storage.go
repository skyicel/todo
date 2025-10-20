package storage

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"todo/models"
)

var FILELISTNAME string = "list.txt"

func LoadTasks() []models.Task {
	var list []models.Task

	file, err := os.OpenFile(FILELISTNAME, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		fmt.Print("TODO LIST cant open file error")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		list = append(list, models.CreateTaskFromFileString(scanner.Text()))
	}

	return list
}

func SaveTasks(TaskList []models.Task) {
	file, err := os.OpenFile(FILELISTNAME, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	for _, elem := range TaskList {
		fmt.Println(elem.ID, elem.Text, elem.Status)
	}

	if err != nil {
		fmt.Print("error")
		return
	}

	defer file.Close()

	for _, task := range TaskList {
		strTask := []byte(strconv.Itoa(task.ID) + " " + task.Text +
			" " + task.Status + " " + task.CreatedTime.Format(time.DateTime) + "\n")

		file.Write(strTask)

		fmt.Print("we are here")
	}
}
