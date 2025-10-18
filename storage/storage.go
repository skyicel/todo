package storage

import (
	"fmt"
	"todo/models"
)

type List []models.Task

var list List

func LoadTask() []models.Task {
	return list
}

func SaveTasks() {
	fmt.Print("saving...")
}
