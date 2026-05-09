package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Steven-Yabann/gtodo/todo"
)

func AddTask(todo *todo.Todos, args []string) {
	// Define add command to add todo item
	addCmd := flag.NewFlagSet("add", flag.ExitOnError)
	addTitle := addCmd.String("title", "", "Title of new todo item")
	addTask := addCmd.String("task", "", "Content of new todo item")

	// Parse the arguments
	addCmd.Parse(args)

	// Check if the todo text was provided
	if *addTask == "" {
		fmt.Println("Error: the --task flag is required")
		os.Exit(1)
	}

	todo.AddTask(*addTitle, *addTask)
	err := todo.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	// print todo
	fmt.Println("Todo item added successfully.")
}