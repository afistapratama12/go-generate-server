package internal

import (
	"fmt"
	"os"
)

// remove file in any directory (repo, service, controller, route)
// remove installing route in main.go

func ProcessRemove(nameEntity string) {
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

	err = RemoteRouteInjectInMain(nameEntity)
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] error remove controller : %s", err))
	}

	if len(errs) > 0 {
		errMsg := ""
		for _, e := range errs {
			errMsg += e.Error() + "\n"
		}

		fmt.Println(errMsg)
	} else {
		fmt.Printf("\n[+] success remove service entity : %s\n", nameEntity)
	}
}

func RemoteRouteInjectInMain(nameEntity string) error {
	// delete route in main.go
	return nil
}
