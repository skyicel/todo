package add

import (
	"todo/models"
	"todo/storage"
)

func PrintList(list storage.List) {
	for _, elem := range list {
		models.PrintTask(elem)
	}
}
