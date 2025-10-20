package storage

import (
	"encoding/json"
	"fmt"
	"os"
	"todo/models"
)

var FILELISTNAME string = "todo.json"

func LoadTasks() ([]models.Task, error) {
	var tasks []models.Task

	file, err := os.OpenFile(FILELISTNAME, os.O_RDONLY|os.O_CREATE, 0666)

	if err != nil {
		return []models.Task{}, fmt.Errorf("LoadTask failed. File was not opened")
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&tasks)

	if err != nil {
		return []models.Task{}, nil
	}

	return tasks, nil
}

func SaveTasks(TaskList []models.Task) error {
	file, err := os.OpenFile(FILELISTNAME, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	encoder.SetIndent("", " ")
	return encoder.Encode(TaskList)
}
