package storage

import (
	"bufio"
	"fmt"
	"os"
	"todo/models"
)

var FILELISTNAME string = "list.txt"

type List []models.Task

func LoadTasks() []models.Task {
	var list List

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

func SaveTasks(list List) {
	models.WriteTaskToFile(list[len(list)-1], FILELISTNAME)
}
