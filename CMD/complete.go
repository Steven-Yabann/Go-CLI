package cmd

import (
	"flag"
	"fmt"
	"log"

	"github.com/todo"
)


func CompleteTask(todo *todo.Todos, args []string) {
	// Create the complete task command
	completeCmd := flag.NewFlagSet("complete", flag.ExitOnError)
	// completeBool = completeCmd.Bool("complete", false, "Bool to confirm task is complete")
	completeIdx := completeCmd.Int("index", -1, "Integer for completed task")

	// Parse arguments
	completeCmd.Parse(args)

	err := todo.Complete(*completeIdx)
	if err != nil {
		log.Fatal(err)
	}

	err = todo.Store(GetJsonFile())
	if err != nil {
		log.Fatal(err)
	}

	// print todo list
	fmt.Println("Completed task has been updated")

}