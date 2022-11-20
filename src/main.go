package main

import (
	"context"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Rayato159/clean-arch-cli/utils"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
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
		utils.CreateListInit(ctx, os.Args[2])
	case "i":
		if os.Args[2] == "" {
			log.Fatalf("error, project name must not empty")
		}
		utils.CreateListInit(ctx, os.Args[2])
	case "-module":
		if os.Args[2] == "" {
			log.Fatalf("error, missing project name")
		}
		if os.Args[3] == "" {
			log.Fatalf("error, missing module name")
		}
		moduleName := strings.ToLower(os.Args[3])
		projectName := os.Args[2]
		utils.CreateModule(ctx, projectName, moduleName)
	case "-m":
		if os.Args[2] == "" {
			log.Fatalf("error, missing project name")
		}
		if os.Args[3] == "" {
			log.Fatalf("error, missing module name")
		}
		moduleName := strings.ToLower(os.Args[3])
		projectName := os.Args[2]
		utils.CreateModule(ctx, projectName, moduleName)
	default:
		log.Fatalf("error, unknown command: %v", os.Args[1])
	}
}
