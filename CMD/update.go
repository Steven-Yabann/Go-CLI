package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/todo"
)

func UpdateTask(todo *todo.Todos, args []string) {
	// Create update command
	updateCmd := flag.NewFlagSet("update", flag.ExitOnError)

	// update args
	updateIdx := updateCmd.Int("index", -1, "Update index")
	updateTitle := updateCmd.String("title", "", "Update title")
	updateTask := updateCmd.String("task", "", "update task")

	// parse argument
	updateCmd.Parse(args)

	err := todo.Update(*updateIdx, *updateTitle, *updateTask)
	if err != nil {
		log.Fatal(err)
	}

	err = todo.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	// print todo
	fmt.Println("Todo item updated successfully.")
}