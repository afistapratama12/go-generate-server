package cmd

import (
	"fmt"
	"go-gen-server/code/middleware"
	"go-gen-server/internal"
	"os"

	"github.com/spf13/cobra"
)

var initMwCmd = &cobra.Command{
	Use:     "init-middleware",
	Aliases: []string{"init-mw", "im"},
	Short:   "init-middleware to initialize middleware",
	Long:    "init-middleware to initialize middleware with auth function and middlware function",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running init middleware")

		// install jwt-go github.com/dgrijalva/jwt-go
		if len(args) < 1 {

			fmt.Println("\033[31mERROR: init-middleware command must have param\033[0m")
			fmt.Println(`
go-gen init-mw [PARAM]
			
example : go-gen init-mw user_id:int,role:string`)
			os.Exit(1)
		} else {
			ProcessInitMiddlware(args[0])
		}
	},
}

func ProcessInitMiddlware(param string) {
	//serverName := config.GetPackageName()

	// TODO: check init must be call first

	//config.ProcessInitMiddleware()

	// auth
	internal.ProcessCreateOrOpenFile("config", internal.AuthDir, middleware.AddAuthService(param))

	// modify routes root.go
	internal.ModifyRoutesRoot()

}

func init() {
	rootCmd.AddCommand(initMwCmd)
}
