package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Rayato159/clean-arch-cli/utils"
)

func main() {
	defer fmt.Scanln()
	// Create directory if not exist
	// os.Args[2] = project_name
	if len(os.Args) == 1 {
		log.Fatalf("error, command not found")
	}
	commands := map[string]string{
		"init": os.Args[1],
	}
	if commands["init"] != "" {
		switch os.Args[2] {
		case "":
			log.Fatalf("error, project name must not empty")
		default:
			utils.CreateListInit(os.Args[2])
		}
	}
}
