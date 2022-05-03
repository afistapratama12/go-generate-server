package cmd

import (
	"fmt"
	"go-gen-server/code/controller"
	"go-gen-server/code/model"
	"go-gen-server/code/repository"
	"go-gen-server/code/routes"
	"go-gen-server/code/service"
	"go-gen-server/internal"
	optionmodel "go-gen-server/option-model"
	"os"

	"github.com/spf13/cobra"
)

var createOpts optionmodel.CreateListOpts

var createCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c", "cr"},
	Short:   "create new service to layer",
	Long: `create service with layer entity, repository, service and controller

example : go-gen create users name:string,email:string,password:string,role:string,status:int

\033[33m NOTE: make sure name entity is sigular not plural, ex : user not users \033[33m
	`,
	Run: func(cmd *cobra.Command, args []string) {
		// create <entity-name> <props>
		// entity-name, ex : users (jamak)
		// props, ex : id:int,name:string,description:string,username:string,password:string
		// default : id:int,created_at:string,updated_at:string
		// can using --without-default-model (tanpa default props) | alias: --wdm
		// helper custom
		// --with-wm / --with-middleware  (create api with middleware, must set the middleware first)
		// --template-user (add login, register, and custom save data to db)
		// --query-order-by-desc / --q-order-by-desc (get all desc with default param: created_at)
		// --query-order-by-asc / --q-order-by-asc (get all asc with default param: created_at)

		if len(args) < 1 {
			fmt.Println(`
\033[31m ERROR: create command must have entity name \033[31m

go-gen create [ENTITY-NAME] [PROPERTIES]

example : go-gen create user name:string,email:string,password:string`)
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
	serverName := internal.GetPackageName()

	// check command init must call first
	err := internal.CheckInitCreated()
	if err != nil {
		fmt.Println("error created :", err)
		return
	}

	// entity
	internal.ProcessCreateOrOpenFile(name, internal.ModelDir, model.AddCreateModel(name, props, isDefault))

	// repo
	internal.ProcessCreateOrOpenFile(name, internal.RepoDir, repository.AddRepository(name, serverName))

	// service
	internal.ProcessCreateOrOpenFile(name, internal.ServiceDir, service.AddService(name, serverName))

	// controller
	internal.ProcessCreateOrOpenFile(name, internal.ControllerDir, controller.AddController(name, serverName))

	// routes
	// TODO: if using middleware, must create first
	internal.ProcessCreateOrOpenFile(name, internal.RouteDir, routes.AddRoutes(name, serverName, createOpts.WithMiddlware))

	// change main.go
	internal.AddRoutesInMain(name, serverName)
}

func init() {
	// TODO: if use this "--with-mw", middleware must create first
	createCmd.Flags().BoolVar(&createOpts.WithMiddlware, "with-mw", false, "create command: with middleware in route generate")

	createCmd.Flags().BoolVar(&createOpts.QueryOrderByASC, "q-order-by-asc", false, "create command: query in repository generate getAll order by ASC")
	createCmd.Flags().BoolVar(&createOpts.QueryOrderByDESC, "q-order-by-desc", false, "create command: query in repository generate getAll order by DESC")

	createCmd.Flags().BoolVar(&createOpts.TemplateUser, "template-user", false, "create command: generate code add login, register")

	rootCmd.AddCommand(createCmd)
}
