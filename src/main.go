package main

import (
	"log"
	"os"
	"strings"

	"github.com/Rayato159/clean-arch-cli/utils"
)

func main() {
	// Create directory if not exist
	// os.Args[2] = project_name
	if len(os.Args) == 1 {
		log.Fatalf("error, command not found")
	}
	switch os.Args[1] {
	case "init":
		if os.Args[2] == "" {
			log.Fatalf("error, project name must not empty")
		}
		utils.CreateListInit(os.Args[2])
	case "i":
		if os.Args[2] == "" {
			log.Fatalf("error, project name must not empty")
		}
		utils.CreateListInit(os.Args[2])
	case "-module":
		moduleName := strings.ToLower(os.Args[2])
		utils.CreateModule(moduleName)
	case "-m":
		moduleName := strings.ToLower(os.Args[2])
		utils.CreateModule(moduleName)
	default:
		log.Fatalf("error, unknown command: %v", os.Args[1])
	}
}
