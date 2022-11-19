package utils

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func CreateEntitiy(destination string, name string) {
	defer log.Printf("entity of module: %s has been created on %s", name, destination)
	data := strings.ReplaceAll(strings.ReplaceAll(`package entities
	
type -Context string

const (
	_Con _Context = "_Controller"
	_Use _Context = "_Usecase"
	_Rep _Context = "_Repository"
)

type _Repository interface {}
type _Usecase interface {}`, "_", strings.Title(name)), "-", name)

	fileName := fmt.Sprintf("%s/%s", destination, fmt.Sprintf("%s_entity.go", name))
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateRepository(destination string, name string) {
	defer log.Printf("repository of module: %s has been created on %s", name, destination)
	data := strings.ReplaceAll(strings.ReplaceAll(`package entities

package repositories

type -Repo struct {}

func New_Repository() entities._Repository {
	return &-Repo{}
}`, "_", strings.Title(name)), "-", name)

	fileName := fmt.Sprintf("%s/repositories/%s", destination, fmt.Sprintf("%s_repository.go", name))
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateUsecase(destination string, name string) {
	defer log.Printf("usecase of module: %s has been created on %s", name, destination)
	data := strings.ReplaceAll(strings.ReplaceAll(`package usecases

type -Use struct {}

func New_Usecase() entities._Usecase {
	return &-Use{}
}`, "_", strings.Title(name)), "-", name)

	fileName := fmt.Sprintf("%s/usecases/%s", destination, fmt.Sprintf("%s_usecase.go", name))
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateController(destination string, name string) {
	defer log.Printf("usecase of module: %s has been created on %s", name, destination)
	data := strings.ReplaceAll(strings.ReplaceAll(`package controllers
	
type -Con struct {}

func New_Controller() {
	controller := &-Con{}
}`, "_", strings.Title(name)), "-", name)

	fileName := fmt.Sprintf("%s/controllers/%s", destination, fmt.Sprintf("%s_controller.go", name))
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateServerModule(destination string) {
	defer log.Printf("server module has been created on %s", destination)
	data := `package servers

type server struct {}

func NewServer() *server {
	return &server{}
}

func (s *server) Start() {
	// Map all routes

	// Server config

	// Start server
}

func (s *server) mapHandlers() error {
	// Middleware list

	// Endpoint list

	// End point not found error response
	return nil
}`

	fileName := fmt.Sprintf("%s/server.go", destination)
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}
