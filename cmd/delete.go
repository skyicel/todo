package cmd

import (
	"slices"
	"todo/models"
)

func DeleteTask(ID int, TaskList []models.Task) []models.Task {
	for index, elem := range TaskList {
		if elem.ID == ID {
			TaskList = slices.Delete(TaskList, index, index+1)
			break
		}
	}

	return TaskList
}
