package config

import (
	"fmt"
	"os"
)

const (
	RepoDir       = "repository"
	ServiceDir    = "service"
	ControllerDir = "controller"
	RouteDir      = "routes"
)

func FailError(err error) {
	if err != nil {
		fmt.Println("error config-os", err.Error())
		return
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
		} else {
			FailError(err)
		}
	}

	fmt.Println("create dir")
}

func CreateOrOpenFile(name string, layer string) {
	dir, err := os.Getwd()
	FailError(err)

	fileLayer := fmt.Sprintf("%s/%s/%s.go", dir, layer, name)
	_, err = os.OpenFile(fileLayer, os.O_RDWR|os.O_CREATE, os.ModePerm)
	FailError(err)

	fmt.Println("create file")
}
