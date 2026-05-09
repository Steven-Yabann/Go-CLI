package cmd

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/todo"
)

func DeleteTask(todo *todo.Todos, args []string) {
	// setup delete command
	deleteCmd := flag.NewFlagSet("delete", flag.ExitOnError)
	index := deleteCmd.Int("delete index", -1, "delete task with index")

	// Parse the arguments
	deleteCmd.Parse(args)

	// delete the task
	err := todo.Delete(*index)
	if err != nil {
		log.Fatal(err)
	}

	// store the updated task list
	err = todo.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)		
	} 

	// print todo list
	fmt.Println("Todo item deleted successfully.")
}