package cmd

import (
	"flag"

	"github.com/Steven-Yabann/gtodo/todo"
)

func ListTasks(todo *todo.Todos, args []string) {
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)

	showDone := listCmd.Bool("done", false, "Show completed tasks only")

	listCmd.Parse(args)

	if *showDone {
		todo.Print()
	} else {
		todo.Print()
	}
}