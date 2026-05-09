package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/Steven-Yabann/gtodo/todo"
)

// Get JSON file
func GetJsonFile() string{
	homeDir, err := os.Getwd()

	if err != nil {
		log.Fatal(err)
	}

	return filepath.Join(homeDir, "todos.json")
}

// function to get user approval to initialize a list
func GetUserApproval() bool {
	confirmMessage := "Need to create an empty \".todos.json\" file in your home directory to store your todo items, continue? (y/n): "

	r := bufio.NewReader(os.Stdin)
	var s string

	fmt.Println(confirmMessage)
	s, _ = r.ReadString('\n')
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	for {
		if s == "y" || s == "yes" {
			return true
		}

		if s == "n" || s == "no" {
			return false
		}
	}
}

func RemindInit(todos *todo.Todos) {
	// check if .todos.json already exists in user home directory
	_, err := os.Stat(GetJsonFile())
	if err != nil {
		fmt.Println("Please run \"init\" subcommand to create an JSON file to store your todo items.")
		os.Exit(1)
	} else {
		if err := todos.Load(GetJsonFile()); err != nil {
			log.Fatal(err)
		}
	}
}

