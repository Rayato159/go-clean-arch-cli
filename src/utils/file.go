package utils

import (
	"fmt"
	"io/ioutil"
	"log"
)

func CreateGitignore(projectName string) {
	data := `# Binaries for programs and plugins
*.exe
*.exe~
*.dll
*.so
*.dylib
*.log

# Test binary
*.test
*.tmp

# Output of the go coverage tool, specifically when used with LiteIDE
*.out

# IDE files
.vscode
.DS_Store
.idea

# Misc
*.fiber.gz
*.fasthttp.gz
*.pprof
*.workspace

# Dependencies
/vendor/
vendor/
vendor
/Godeps/

# Customs
*.env*
.air*
/tmp`

	fileName := fmt.Sprintf("%s/.gitignore", projectName)
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateDotEnv(projectName string, state string) {
	data := fmt.Sprintf(`STAGE=%s
APP_HOST=
APP_PORT=
APP_REQUEST_TIMEOUT=

DB_HOST=
DB_PORT=
DB_DATABASE=
DB_USERNAME=
DB_PASSWORD=

FILE_LOG_PATH=./assets/logs`, state)

	fileName := fmt.Sprintf("%s/.env.%s", projectName, state)
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}

func CreateMain(projectName string) {
	data := `package main

import "fmt"

func main() {
	fmt.Println("hello, clean arch cli")
}
`

	fileName := fmt.Sprintf("%s/app/main.go", projectName)
	err := ioutil.WriteFile(fileName, []byte(data), 0777)
	if err != nil {
		log.Fatalf("error, can't create a file: %v, with an error: %v", fileName, err.Error())
	}
}
