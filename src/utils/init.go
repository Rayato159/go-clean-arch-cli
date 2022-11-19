package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/Rayato159/clean-arch-cli/models"
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

func CreateListInit(projectName string) {
	ls := models.NewDefaultFileAndFolderList()
	ls.SetGetDefaultFileAndFolderList(projectName)
	list := ListDir(ls)

	// Creat project folder
	if err := os.Mkdir(projectName, os.FileMode(0777)); err != nil {
		log.Fatalf("error, can't create a dir with an error: %v", err.Error())
	}
	CreateDir(projectName, list)
	CreateFile(projectName, list)
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
				GitignoreCreator(projectName)
			case fmt.Sprintf("%s/.env.dev", projectName):
				DotEnvCreator(projectName, "dev")
			case fmt.Sprintf("%s/.env.test", projectName):
				DotEnvCreator(projectName, "test")
			case fmt.Sprintf("%s/.env.prod", projectName):
				DotEnvCreator(projectName, "prod")
			}
		}
	}
}
