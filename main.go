package main

import (
	"os"
	"github.com/Steven-Yabann/gtodo/cmd"  
	"github.com/Steven-Yabann/gtodo/todo" 
)

func main() {
	todos := &todo.Todos{}

	if len(os.Args) < 2 {
		cmd.Help()
		return
	}

	switch os.Args[1] {
	case "init":
		cmd.Init()
	case "add":
		cmd.RemindInit(todos)
		cmd.AddTask(todos, os.Args[2:])
	case "update":
		cmd.RemindInit(todos)
		cmd.UpdateTask(todos, os.Args[2:])
	case "delete":
		cmd.RemindInit(todos)
		cmd.DeleteTask(todos, os.Args[2:])
	case "complete":
		cmd.RemindInit(todos)
		cmd.CompleteTask(todos, os.Args[2:])
	case "list":
		cmd.RemindInit(todos)
		cmd.ListTasks(todos, os.Args[2:])
	default:
		cmd.Help()
	}
}