package internal

import (
	"fmt"
	"go-gen-server/code"
	"go-gen-server/code/config"
	"os"
	"os/exec"
	"strings"
)

func GetPackageName() string {
	dir, err := os.Getwd()
	FailError(err)

	rootDirService := strings.Split(dir, "/")

	return rootDirService[len(rootDirService)-1] + "-service"
}

func ProcessInitServer(projectName string, driver string) {
	dir, err := os.Getwd()
	FailError(err)

	if projectName == "." {
		ProcessInitAndGetPackage(dir, driver)
	} else {
		// create folder server
		CreateOrOpenDir(projectName)

		os.Chdir(dir + "/" + projectName)
		newDir, err := os.Getwd()
		FailError(err)

		ProcessInitAndGetPackage(newDir, driver)

		fmt.Println("success create directory")
	}
}

func ProcessInitMiddleware() {
	goExc, err := exec.LookPath("go")
	if err != nil {
		fmt.Println(err.Error())
	}

	// running install jwt
	initGo := &exec.Cmd{
		Path:   goExc,
		Args:   []string{goExc, "get", "-u", "github.com/dgrijalva/jwt-go"},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := initGo.Run(); err != nil {
		FailError(err)
	}
}

func ProcessInitAndGetPackage(dir string, driver string) {
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

	var execDriver string

	switch driver {
	case "mysql":
		execDriver = config.DriverMysql
		break
	case "postgres":
		execDriver = config.DriverPostgres
		break
	case "sqlite":
		execDriver = config.DriverSqlite
		break
	default:
		execDriver = config.DriverMysql
		break
	}

	// default install mysql
	dbGo := &exec.Cmd{
		Path:   goExc,
		Args:   []string{goExc, "get", "-u", execDriver},
		Stdout: os.Stdout,
		Stderr: os.Stdout,
	}

	if err := dbGo.Run(); err != nil {
		FailError(err)
	}

	createAnotherFileInit()
}

func createAnotherFileInit() {

	var errs = []error{}

	f, err := os.Create(".env")
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] create file .env error: %s", err))
	}

	_, err = f.WriteString(code.DotEnv)
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] write file .env error: %s", err))
	}

	f.Close()

	f2, err := os.Create("example.env")
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] create file example.env error: %s", err))
	}

	_, err = f2.WriteString(code.ExampleEnv)
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] write file example.env error: %s", err))
	}

	f2.Close()

	f3, err := os.Create(".gitignore")
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] create file .gitignore error: %s", err))
	}

	_, err = f3.WriteString(code.Gitignore)
	if err != nil {
		errs = append(errs, fmt.Errorf("[!] write file .gitignore error: %s", err))
	}

	f3.Close()

	if len(errs) > 0 {
		var errMsg string

		for _, e := range errs {
			errMsg += e.Error() + "\n"
		}

		fmt.Println(errMsg)
	}
}
