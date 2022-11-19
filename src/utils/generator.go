package utils

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/Rayato159/clean-arch-cli/models"
)

type GeneratorContext string

const (
	GeneratorCall GeneratorContext = "GeneratorCall"
)

func ListDir(ls models.DefaultFileAndFolderListInterface) map[string]models.FileInfo {
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err.Error())
	}

	filesMap := ls.GetDefaultFileAndFolderList()
	for i := range files {
		if entry, ok := filesMap[files[i].Name()]; ok {
			entry.IsExist = true
			filesMap[files[i].Name()] = entry
		}
	}
	return filesMap
}

func CreateListInit(ctx context.Context, projectName string) {
	ctx = context.WithValue(ctx, GeneratorCall, time.Now().UnixMilli())
	log.Printf("called:\t%v", ctx)
	defer log.Printf("return:\t%v", ctx)

	ls := models.NewDefaultFileAndFolderList()
	ls.SetGetDefaultFileAndFolderList(projectName)
	list := ListDir(ls)

	// Creat project folder
	if err := os.Mkdir(projectName, os.FileMode(0777)); err != nil {
		log.Fatalf("error, can't create a dir with an error: %v", err.Error())
	}
	CreateDir(projectName, list)
	CreateFile(projectName, list)

	// Go mod init
	// cmd := exec.Command(fmt.Sprintf("go mod init %s", projectName))
	// cmd.Dir = "./" + projectName
	// if err := cmd.Run(); err != nil {
	// 	log.Fatal(err)
	// }
}

func CreateDir(projectName string, list map[string]models.FileInfo) {
	for i := range list {
		if !list[i].IsExist && list[i].IsDir {
			log.Printf("folder: %v has been created", i)

			switch i {
			case fmt.Sprintf("%s/assets", projectName):
				dir := []string{
					fmt.Sprintf("%s/logs", i),
					fmt.Sprintf("%s/screenshots", i),
				}
				for j := range dir {
					if err := os.MkdirAll(dir[j], os.FileMode(0777)); err != nil {
						log.Fatalf("error, can't create a dir with an error: %v", err.Error())
					}
				}
			case fmt.Sprintf("%s/app", projectName):
				if err := os.MkdirAll(i, os.FileMode(0777)); err != nil {
					log.Fatalf("error, can't create a dir with an error: %v", err.Error())
				}
				CreateMain(projectName)
			case fmt.Sprintf("%s/modules", projectName):
				destination := filepath.Join(projectName, "modules", "servers")
				if err := os.MkdirAll(destination, os.FileMode(0777)); err != nil {
					log.Fatalf("error, can't create a dir with an error: %v", err.Error())
				}
				CreateServerModule(destination)
			case fmt.Sprintf("%s/package", projectName):
				dir := []string{
					fmt.Sprintf("%s/databases", i),
					fmt.Sprintf("%s/middlewares", i),
					fmt.Sprintf("%s/utils", i),
				}
				for j := range dir {
					if err := os.MkdirAll(dir[j], os.FileMode(0777)); err != nil {
						log.Fatalf("error, can't create a dir with an error: %v", err.Error())
					}
				}
			default:
				if err := os.MkdirAll(i, os.FileMode(0777)); err != nil {
					log.Fatalf("error, can't create a dir with an error: %v", err.Error())
				}
			}
		}
	}
}

func CreateFile(projectName string, list map[string]models.FileInfo) {
	for i := range list {
		if !list[i].IsExist && !list[i].IsDir {
			switch i {
			case fmt.Sprintf("%s/.gitignore", projectName):
				CreateGitignore(projectName)
			case fmt.Sprintf("%s/.env.dev", projectName):
				CreateDotEnv(projectName, "dev")
			case fmt.Sprintf("%s/.env.test", projectName):
				CreateDotEnv(projectName, "test")
			case fmt.Sprintf("%s/.env.prod", projectName):
				CreateDotEnv(projectName, "prod")
			}
		}
	}
}

func GetProjectName() string {
	// Check wether pwd is already in project?
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err.Error())
	}
	if len(files) == 0 {
		log.Fatal("error, project not found")
	}
	var fileStack string
	for i := range files {
		if files[i].Name() != "main.exe" && files[i].IsDir() {
			fileStack = files[i].Name()
			break
		}
	}
	if fileStack == "" {
		log.Fatal("error, project must be a folder")
	}
	return fileStack
}

func BeforeCreateModuleChecK(projectName, name string) {
	// Check a module name
	path := fmt.Sprintf("./%s/modules", projectName)
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err.Error())
	}
	for i := range files {
		if files[i].Name() == name {
			log.Fatalf("error, module: %s is already exist", name)
		}
	}
}

func CreateModule(ctx context.Context, name string) {
	ctx = context.WithValue(ctx, GeneratorCall, time.Now().UnixMilli())
	log.Printf("called:\t%v", ctx)
	defer log.Printf("return:\t%v", ctx)

	projectName := GetProjectName()
	BeforeCreateModuleChecK(projectName, name)

	// Mkdir for module
	destination := filepath.Join(projectName, "modules", name)
	if err := os.Mkdir(destination, os.FileMode(0777)); err != nil {
		log.Fatalf("error, can't create a dir with an error: %v", err.Error())
	}

	// Nested mkdir for module
	destinationModuleInternal := []string{
		"entities",
		"repositories",
		"usecases",
		"controllers",
	}
	for i := range destinationModuleInternal {
		destination := filepath.Join(projectName, "modules", name, destinationModuleInternal[i])
		if err := os.Mkdir(destination, os.FileMode(0777)); err != nil {
			log.Fatalf("error, can't create a dir with an error: %v", err.Error())
		}
	}

	// Create module
	CreateEntitiy(destination, name)
	CreateRepository(destination, name)
	CreateUsecase(destination, name)
	CreateController(destination, name)
}
