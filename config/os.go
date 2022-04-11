package config

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const (
	RepoDir       = "repository"
	ServiceDir    = "service"
	ControllerDir = "controller"
	RouteDir      = "routes"
	ModelDir      = "entity"
	MainFile      = "main"
)

func FailError(err error) {
	if err != nil {
		fmt.Println("error config-os : ", err)
		return
	}
}

// ProcessCreateOrOpenFile is function to running create or open dir, then create or open file
func ProcessCreateOrOpenFile(name string, layer string, data string) {
	// directory of layer
	if layer != "" {
		CreateOrOpenDir(layer)
	}

	// craete file in directory layer
	CreateOrOpenFile(name, layer, data)
}

func ProcessInitServer() {
	dir, err := os.Getwd()
	FailError(err)

	rootDirService := strings.Split(dir, "/")

	nameService := rootDirService[len(rootDirService)-1] + "-service"

	goExc, err := exec.LookPath("go")
	if err != nil {
		fmt.Println(err.Error())
	}

	// running go mod init
	initGo := &exec.Cmd{
		Path:   goExc,
		Args:   []string{goExc, "mod", "init", nameService},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := initGo.Run(); err != nil {
		FailError(err)
	}

	// install gorm
	gormGo := &exec.Cmd{
		Path:   goExc,
		Args:   []string{goExc, "get", "-u", "gorm.io/gorm"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := gormGo.Run(); err != nil {
		FailError(err)
	}

	// install gin-gonic
	ginGo := &exec.Cmd{
		Path:   goExc,
		Args:   []string{goExc, "get", "-u", "github.com/gin-gonic/gin"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := ginGo.Run(); err != nil {
		FailError(err)
	}

	// default install mysql
	dbGo := &exec.Cmd{
		Path:   goExc,
		Args:   []string{goExc, "get", "-u", "gorm.io/driver/mysql"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := dbGo.Run(); err != nil {
		FailError(err)
	}
}

func CreateOrOpenDir(layer string) {
	dir, err := os.Getwd()
	FailError(err)

	dirLayer := fmt.Sprintf("%s/%s", dir, layer)

	//create directory
	_, err = os.Open(dirLayer)
	if err != nil {
		if os.IsNotExist(err) {
			err = os.Mkdir(dirLayer, os.ModePerm)
			FailError(err)

			fmt.Printf("\n[+] success create directory layer : %s", layer)
			return
		} else {
			FailError(err)
		}
	}

	fmt.Printf("\n[+] success open directory layer : %s", layer)
}

func CreateOrOpenFile(name string, layer string, data string) {
	dir, err := os.Getwd()
	FailError(err)

	var fileLayer string
	if layer == "" {
		// case for main.go
		fileLayer += fmt.Sprintf("%s/%s.go", dir, name)
	} else {
		fileLayer += fmt.Sprintf("%s/%s/%s.go", dir, layer, name)
	}

	file, err := os.OpenFile(fileLayer, os.O_RDWR|os.O_CREATE, os.ModePerm)
	FailError(err)

	fmt.Printf("\n[+] success create file name : %s.go", name)

	defer file.Close()

	_, err = file.Write([]byte(data))
	FailError(err)

	fmt.Printf("\n[+] success insert code to file : %s.go", name)
}
