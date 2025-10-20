package main

import (
	"fmt"
	"os"
	"strconv"
	"todo/cmd"
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
		nTask := models.CreateTask(command[2], data)
		data = cmd.AddTask(nTask, data)
	case LIST:
		cmd.PrintList(data)
	case COMPLETE:
		ID, _ := strconv.Atoi(command[2])

		cmd.CompleteTask(ID, &data)
	case DELETE:
		ID, _ := strconv.Atoi(command[2])

		fmt.Print("\n\n\n")

		data = cmd.DeleteTask(ID, data)

	case HELP:
		fmt.Print(HELP)
	default:
		fmt.Print(HELP)
	}

	storage.SaveTasks(data)
}

func showhelp() {
	fmt.Print("no arguement bro!")
}
