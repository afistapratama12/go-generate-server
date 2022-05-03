package internal

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

// remove file in any directory (repo, service, controller, route)
// remove installing route in main.go

func ProcessRemove(nameEntity string) {
	if nameEntity == "." {
		// remove all entity service in current directory
		RemoveAllService()
	} else {
		RemoveSpecificEntity(nameEntity)
	}
}

func RemoveAllService() {
	var errs = []error{}

	var entities []string

	// delete all repo
	files, err := ioutil.ReadDir("/repository")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		entities = append(entities, strings.Split(file.Name(), ".")[0])

		err := os.Remove(fmt.Sprintf("controller/%s", file.Name()))
		if err != nil {
			errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
		}
	}

	// delete all controller
	files, err = ioutil.ReadDir("/controller")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := os.Remove(fmt.Sprintf("controller/%s", file.Name()))
		if err != nil {
			errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
		}
	}

	// delete all service
	files, err = ioutil.ReadDir("/service")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := os.Remove(fmt.Sprintf("service/%s", file.Name()))
		if err != nil {
			errs = append(errs, fmt.Errorf("[!] error remove service : %s", err))
		}
	}

	// remove all entity
	files, err = ioutil.ReadDir("/entity")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		err := os.Remove(fmt.Sprintf("entity/%s", file.Name()))
		if err != nil {
			errs = append(errs, fmt.Errorf("[!] error remove entity : %s", err))
		}
	}

	// remove all routes
	files, err = ioutil.ReadDir("/routes")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		if file.Name() == "root.go" {
			continue
		} else {
			err := os.Remove(fmt.Sprintf("routes/%s", file.Name()))
			if err != nil {
				errs = append(errs, fmt.Errorf("[!] error remove routes : %s", err))
			}
		}
	}

	for _, ent := range entities {
		RemoveRoutesInMain(ent)
	}

	if len(errs) > 0 {
		errMsg := ""
		for _, e := range errs {
			errMsg += e.Error() + "\n"
		}

		fmt.Println(errMsg)
	}

	fmt.Printf("\n[+] success remove all service \n")

}

func RemoveSpecificEntity(nameEntity string) {
	var errs = []error{}

	err := os.Remove(fmt.Sprintf("controller/%s.go", nameEntity))
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
	}

	err = os.Remove(fmt.Sprintf("service/%s.go", nameEntity))
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
	}

	err = os.Remove(fmt.Sprintf("repository/%s.go", nameEntity))
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
	}

	err = os.Remove(fmt.Sprintf("entity/%s.go", nameEntity))
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
	}

	err = os.Remove(fmt.Sprintf("routes/%s.go", nameEntity))
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
	}

	RemoveRoutesInMain(nameEntity)

	if len(errs) > 0 {
		errMsg := ""
		for _, e := range errs {
			errMsg += e.Error() + "\n"
		}

		fmt.Println(errMsg)
	}

	fmt.Printf("\n[+] success remove service entity : %s\n", nameEntity)
}
