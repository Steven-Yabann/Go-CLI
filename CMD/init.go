package cmd

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func Init() {
	ok := GetUserApproval()
	if !ok {
		fmt.Print("You've declined to create the JSON file. You can always run \"init\" subcommand again if you change your mind.")	
		os.Exit(0)		
	}

	homeDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	filepath := filepath.Join(homeDir, "todos.json")

	// check if json file already exists
	_, err = os.Stat(filepath)
	if err != nil {
		if os.IsNotExist(err) {
			file, err := os.Create(filepath)
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			// Write an empty JSON array so Load() doesn't fail later
			_, err = file.WriteString("[]")
			if err != nil {
				file.Close()
				log.Fatal(err)
			}

			fmt.Println("Succefully create a \".todos.json\" file in your home directory.")
		} else {
			log.Fatal(err)
		}
	} else {
		fmt.Print(".todos.json file exists in your home directory already.")
	}
}