package routes

// di main ditambah 1 route

// 1. "pakcage1/repository"
// 2. "pakcage1/service"
// 3. "pakcage1/controller"
var routesHead = `package routes

import (
	"%s/repository"
	"%s/service"
	"%s/controller"

	"github.com/gin-gonic/gin"
)`

// User
// User
// User
// User
// User
// User
// User
// User
var routesInit = `

var (
	%sRepository = repository.New%sRepository(DB)
	%sService = service.New%sService(%sRepository)
	%sController = controller.New%sController(%sService)
)`

// User
// users
// routeGetAll
// users
// routerGetId
// users
// routeCreate
// users
// routeUpdate
// users
// routeDelete
var routesRoute = `

func %sRoute(r *gin.Engine) {
	v1 := r.Group("/v1")
	{
		v1.GET("/%s/all", %s)
		v1.GET("%s/:id", %s)
		v1.POST("/%s", %s)
		v1.PUT("/%s/:id", %s)
		v1.DELETE("/%s/:id", %s)
	}
}
`

// route without middleware
var (
	routeGetAll = `%sController.GetAll%s`
	routeGetId  = `%sController.Get%sById`
	routeCreate = `%sController.Create%s`
	routeUpdate = `%sController.Update%sById`
	routeDelete = `%sController.Delete%sById`
)

// route with middleware
var (
	routeGetAllMw = `MainMiddleware, %sController.GetAll%s`
	routeGetIdMw  = `MainMiddleware, %sController.Get%sById`
	routeCreateMw = `MainMiddleware, %sController.Create%s`
	routeUpdateMw = `MainMiddleware, %sController.Update%sById`
	routeDeleteMw = `MainMiddleware, %sController.Delete%sById`
)
