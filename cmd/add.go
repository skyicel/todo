package add

import (
	"todo/models"
	"todo/storage"
)

func AddTask(task models.Task, list storage.List) storage.List {
	list = append(list, task)
	return list
}
