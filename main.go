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
	data, err := storage.LoadTasks()

	if err != nil {
		fmt.Print(err)
		return
	}

	switch command[1] {
	case ADD:
		if len(os.Args) < 2 {
			fmt.Println("error, you didn't enter text of task")
			showhelp()
			return
		} else {
			nTask := models.CreateTask(command[2:], data)
			data = cmd.AddTask(nTask, data)
		}
	case LIST:
		if len(data) != 0 {
			cmd.PrintList(data)
		} else {
			fmt.Println("data is empty")
		}
	case COMPLETE:
		ID, err := strconv.Atoi(command[2])

		if err != nil {
			fmt.Print("error! id is not int!")
			return
		}

		cmd.CompleteTask(ID, data)
	case DELETE:
		ID, err := strconv.Atoi(command[2])

		if err != nil {
			fmt.Print("error! id is not int!")
			return
		}

		data = cmd.DeleteTask(ID, data)

	case HELP:
		showhelp()
	default:
		fmt.Println("unknown command")
		showhelp()
	}

	storage.SaveTasks(data)
}

func showhelp() {
	fmt.Println("---------------")
	fmt.Println("TODO COMMANDS")
	fmt.Println("task{id, text, status, created time}")
	fmt.Println("todo add [text] / add task to list")
	fmt.Println("todo list / show list of tasks")
	fmt.Println("todo complete [id] / make task complete")
	fmt.Println("todo delete [id]")
	fmt.Println("---------------")
}
