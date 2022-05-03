package cmd

import (
	"fmt"
	"go-gen-server/code"
	"go-gen-server/code/config"
	"go-gen-server/code/routes"
	"go-gen-server/code/utils"
	"go-gen-server/internal"
	optionmodel "go-gen-server/option-model"
	"os"

	"github.com/spf13/cobra"
)

var initOpts optionmodel.InitListOpts

var initCmd = &cobra.Command{
	Use:     "init",
	Aliases: []string{"i", "in"},
	Short:   "initialize new service",
	Long: `Initialize service with install default package gin-gonic and GORM
	
example : go-gen init my-example
	
note : use '.' if you want create service in recent directory
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running init")
		// get now directory, create directory if null
		// get name folder directory
		// - go mod init name sesuai nama folder <nama folder>
		// - go get -u github.com/gin-gonic/gin
		// - go get -u gorm.io/gorm
		// create config, connection with mysql (default)
		// add helper custom
		// db default using mysql
		// --db-mysql
		// --db-postgres / --db-postgresql / --db-pg
		// --db-sqlite

		if len(args) < 1 {
			fmt.Println("\033[31mERROR: init command must have project name\033[0m")

			fmt.Println(`
go-gen init [PROJECT_NAME]
			
example : go-gen init my-go-app

if you want to generate service code in this directory please using "." , ex : go-gen init .`)
			os.Exit(1)
		} else if len(args) > 1 {
			fmt.Println("\033[31m ERROR: cannot using the option flag, please try again\033[31m")
			os.Exit(1)
		} else {
			ProcessInit(args[0])
		}
	},
}

func ProcessInit(projectName string) {
	var driver string

	if initOpts.ConfigDBGenerate != "" {
		if initOpts.ConfigDBGenerate != "mysql" && initOpts.ConfigDBGenerate != "postgres" && initOpts.ConfigDBGenerate != "sqlite" {
			fmt.Println("error: invalid db driver connection, please try again")
			os.Exit(1)
		} else {
			driver = initOpts.ConfigDBGenerate
		}
	} else {
		driver = "mysql"
	}

	internal.ProcessInitServer(projectName, driver)

	internal.ProcessCreateOrOpenFile(internal.MainFile, "", code.AddInitMain())

	internal.ProcessCreateOrOpenFile(internal.ErrorFile, internal.UtilsDir, utils.AddUtilsError())

	var serverName string
	if projectName != "." {
		serverName += projectName + "-service"
	} else {
		serverName = internal.GetPackageName()
	}

	internal.ProcessCreateOrOpenFile(internal.RootFile, internal.RouteDir, routes.AddMainRoutes(serverName))

	configDBStr, err := config.AddInitConfig(driver)
	if err != nil {
		fmt.Println("[!] error :", err)
		os.Exit(1)
	}

	internal.ProcessCreateOrOpenFile(internal.ConfigFile, internal.ConfigDir, configDBStr)
}

func init() {
	initCmd.Flags().StringVarP(&initOpts.ConfigDBGenerate, "db", "d", "", "set db connection with custom driver, list (mysql | postgres | sqlite)")

	rootCmd.AddCommand(initCmd)
}
