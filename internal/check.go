package internal

import (
	"fmt"
	"os"
)

var errCheckInit = `[!] service init not created
    make sure running on right directory 
    or please run command : go-get init [SERVICE_NAME]`

func CheckInitCreated() error {
	// check any go.mod
	_, err := os.OpenFile("go.mod", os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(errCheckInit)
			return err
		}
	}
	// check any go.sum
	_, err = os.OpenFile("go.sum", os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(errCheckInit)
			return err
		}
	}

	// check main.go
	_, err = os.OpenFile("main.go", os.O_RDONLY, os.ModePerm)
	if err != nil {
		if os.IsNotExist(err) {
			fmt.Println(errCheckInit)
			return err
		}
	}

	return nil
}
