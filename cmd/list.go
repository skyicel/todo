package cmd

import (
	"todo/models"
)

func PrintList(list []models.Task) {
	for _, elem := range list {
		models.PrintTask(elem)
	}
}
