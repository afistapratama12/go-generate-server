package main

import (
	"fmt"
	"regexp"
)

// TODO: planning

// v0.0.2
// finishing logic on service (create, update, delete)
// finishing logic on controller (create, update, delete)
// finishing logic on repository (create, update, delete)

// v0.0.3
// create with user template
// login, register, forgot password, reset password, change password (service, controller, routes)
// add new model / entities (userRegisterResponse, userRegisterRequest, userLoginResponse, userLoginRequest)

// v0.1.0
// stable condition
// handle many bug to ready implement

// v1.0.0
// more stable condition
// create download

func main() {
	//cmd.Execute()

	var data = `var (
		DB = config.ConnectDB()
	)`

	regex, _ := regexp.Compile("DB = config[.]ConnectDB[(][)]")

	newData := regex.ReplaceAllString(data, "DB = config.ConnectDB()\n\tauthService = auth.NewAuthService()\n\tMainMiddleware = middleware.Middleware(authService)")

	fmt.Println(newData)

}
