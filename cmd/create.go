package cmd

import (
	"fmt"
	"go-gen-server/code"
	"go-gen-server/config"
	"os"

	"github.com/spf13/cobra"
)

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "cr"},
	Short:   "create new service to layer",
	Long:    "create service with layer entity, repository, service and controller",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Println("please insert entity name, ex : go-gen-server init user")
			os.Exit(1)
		} else {
			if len(args) == 1 {
				ProcessCreate(args[0], "", true)
			} else {
				ProcessCreate(args[0], args[1], true)
			}
		}
	},
}

func ProcessCreate(name string, props string, isDefault bool) {
	// entity
	config.ProcessCreateOrOpenFile(name, config.ModelDir, code.AddCreateModel(name, props, isDefault))

	// repo
	// config.ProcessCreateOrOpenFile(name, config.RepoDir, "")

	// service
	// config.ProcessCreateOrOpenFile(name, config.ServiceDir, "")

	// controller
	// config.ProcessCreateOrOpenFile(name, config.ControllerDir, "")

	// routes
	// config.ProcessCreateOrOpenFile(name, config.RouteDir, "")

	// change main.go, add routes

}

func init() {
	rootCmd.AddCommand(createCmd)
}
