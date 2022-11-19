package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func CreateEntitiy(destination string, name string) {
	data := strings.ReplaceAll(`package entities

import (
	"context"
)

type _Context string

const (
	_Con _Context = "_Controller"
	_Use _Context = "_Usecase"
	_Rep _Context = "_Repository"
)

type _Repository interface {}
type _Usecase interface {}
`, "_", strings.ToTitle(name))

	fileName := fmt.Sprintf("module/%s/%s", destination, fmt.Sprintf("%s_entity.go", name))
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateRepository(destination string, name string) {

}

func CreateUsecase(destination string, name string) {

}

func CreateController(destination string, name string) {

}
