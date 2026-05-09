package main

import (
	"os"
	"github.com/Steven-Yabann/gtodo/CMD"  // Your updated path
	"github.com/Steven-Yabann/gtodo/todo" // Your updated path
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
	case "list":
		cmd.RemindInit(todos)
		// Implement your List function here
	case "update":
		cmd.RemindInit(todos)
		cmd.UpdateTask(todos, os.Args[2:])
	case "delete":
		cmd.RemindInit(todos)
		cmd.DeleteTask(todos, os.Args[2:])
	case "complete":
		cmd.RemindInit(todos)
		cmd.CompleteTask(todos, os.Args[2:])
	default:
		cmd.Help()
	}
}