package internal

import (
	"fmt"
	"go-gen-server/helper"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"
)

func OpenFileMain() string {
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

	return read
}

func WriteNewMain(name string, newRead string, fileLayer string) {

	err := os.WriteFile(fileLayer, []byte(newRead), 0644)
	FailError(err)

	fmt.Printf("\n[+] success insert new code file : %s.go", name)
}

func AddRoutesInMain(name string, serviceName string) {
	read := OpenFileMain()

	split := strings.Split(read, "\n")

	allNameStr, _ := helper.ConverterName(name, "entity")

	addImport := CheckImportRoutes(serviceName, split)

	readImport := strings.Join(addImport, "\n")

	sepAddRoute := `r[.]Run[(][)]\n`

	regex, err := regexp.Compile(sepAddRoute)
	FailError(err)

	newRoute := fmt.Sprintf("routes.%sRoute(r)", allNameStr.NameUpper)

	newRead := regex.ReplaceAllString(readImport, newRoute+"\n\t"+"r.Run()"+"\n")

	WriteNewMain(name, newRead, "main.go")
}

func RemoveRoutesInMain(name string) {
	read := OpenFileMain()

	sep := fmt.Sprintf(`routes[.]%sRoute[(]r[)]\n`, name)

	regex, err := regexp.Compile(sep)
	if err != nil {
		log.Fatal(err)
		return
	}

	newRead := regex.ReplaceAllString(read, "")

	WriteNewMain(name, newRead, "main.go")

	fmt.Printf("\n[+] success remove routes code from file : %s.go", name)

	if err != nil {
		log.Fatal(err)
		return
	}

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
