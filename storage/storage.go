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

func LoadTasks() ([]models.Task, error) {
	var list []models.Task

	file, err := os.OpenFile(FILELISTNAME, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		return nil, fmt.Errorf("LoadTask failed. File was not opened")
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		list = append(list, models.CreateTaskFromFileString(scanner.Text()))
	}

	return list, nil
}

func SaveTasks(TaskList []models.Task) {
	file, err := os.OpenFile(FILELISTNAME, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		fmt.Print("error")
		return
	}

	defer file.Close()

	for _, task := range TaskList {
		strTask := []byte(strconv.Itoa(task.ID) + " " + task.Text +
			" " + task.Status + " " + task.CreatedTime.Format(time.DateTime) + "\n")

		file.Write(strTask)
	}
}
