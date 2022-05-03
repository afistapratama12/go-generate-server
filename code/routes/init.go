package routes

var routesMain = `package routes

import (
	"%s/config"
)

var (
	DB = config.ConnectDB()
)
`
