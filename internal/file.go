package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// ProcessCreateOrOpenFile is function to running create or open dir, then create or open file
func ProcessCreateOrOpenFile(name string, layer string, data string) {
	// directory of layer
	if layer != "" {
		CreateOrOpenDir(layer)
	}

	// create file in directory layer
	CreateOrOpenFile(name, layer, data)
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

func OpenFile(fileName string) string {
	dir, err := os.Getwd()
	FailError(err)

	var fileLayer = fmt.Sprintf("%s/%s", dir, fileName)

	file, err := os.OpenFile(fileLayer, os.O_RDWR|os.O_CREATE, os.ModePerm)
	FailError(err)

	defer file.Close()

	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	read := string(body)

	return read
}
