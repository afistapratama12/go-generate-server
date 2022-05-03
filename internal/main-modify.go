package internal

import (
	"fmt"
	"go-gen-server/helper"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func AddRoutesInMain(name string, serviceName string) {
	dir, err := os.Getwd()
	FailError(err)

	var fileLayer = fmt.Sprintf("%s/main.go", dir)

	file, err := os.OpenFile(fileLayer, os.O_RDWR|os.O_CREATE, os.ModePerm)
	FailError(err)

	defer file.Close()

	body, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	read := string(body)

	split := strings.Split(read, "\n")

	allNameStr, _ := helper.ConverterName(name, "entity")

	addImport := CheckImportRoutes(serviceName, split)
	newSplit := AddRoutesMain(allNameStr.NameUpper, addImport)

	newRead := strings.Join(newSplit, "\n")

	err = os.WriteFile(fileLayer, []byte(newRead), 0644)
	FailError(err)

	fmt.Printf("\n[+] success insert routes code to file : %s.go", name)
}

func CheckImportRoutes(serviceName string, split []string) []string {
	var (
		newSplit            []string
		checkRoutesImported = false
		flag                = false
	)

	imported := []string{}

	for _, sp := range split {
		if sp == ")" {
			flag = false
		}

		if flag {
			imported = append(imported, sp)
		}

		if sp == "import (" {
			flag = true
		}
	}

	for _, im := range imported {
		splited := strings.Split(im, "\t")

		for _, s := range splited {

			if s == "" || s == " " {
				continue
			} else {
				if s == fmt.Sprintf(`"%s/routes"`, serviceName) {
					checkRoutesImported = true
					break
				}
			}

		}
	}

	var addImport = make([]string, 4)

	if !checkRoutesImported {
		copy(addImport, split[:3])

		addImport[3] = "\t" + fmt.Sprintf(`"%s/routes"`, serviceName)

		newSplit = append(addImport, split[3:]...)
	} else {
		newSplit = split
	}

	return newSplit
}

func AddRoutesMain(nameUpper string, split []string) []string {
	var (
		newSplit []string
		start    = make([]string, len(split))
	)

	copy(start, split[:len(split)-4])

	start = append(start, fmt.Sprintf("\troutes.%sRoute(r)", nameUpper))

	newSplit = append(start, split[len(split)-5:]...)

	return newSplit
}
