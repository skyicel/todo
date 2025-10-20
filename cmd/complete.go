package cmd

import (
	"todo/models"
)

func CompleteTask(ID int, TaskList *[]models.Task) {
	for index, elem := range *TaskList {
		if elem.ID == ID {
			(*TaskList)[index].Status = "completed"
			break
		}
	}
}
