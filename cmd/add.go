package cmd

import (
	"todo/models"
)

func AddTask(task models.Task, list []models.Task) []models.Task {
	list = append(list, task)
	return list
}
