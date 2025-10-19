package main

import (
	"fmt"
	"os"
	add "todo/cmd"
	list "todo/cmd"
	"todo/models"
	"todo/storage"
)

var ADD = "add"
var LIST = "list"
var COMPLETE = "complete"
var DELETE = "delete"
var HELP = "help"

func main() {
	if len(os.Args) == 1 {
		showhelp()
		return
	}

	command := os.Args
	data := storage.LoadTasks()

	switch command[1] {
	case ADD:
		nTask := models.CreateTask(command[2])
		data = add.AddTask(nTask, data)
		storage.SaveTasks(data)
	case LIST:
		list.PrintList(data)
	case COMPLETE:
		fmt.Print(COMPLETE)
	case DELETE:
		fmt.Print(DELETE)
	case HELP:
		fmt.Print(HELP)
	default:
		fmt.Print(HELP)
	}
}

func showhelp() {
	fmt.Print("no arguement bro!")
}
