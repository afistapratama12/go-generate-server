package code

// di main ditambah 1 route

// example:
// 1. "pakcage1/repository"
// 2. "pakcage1/service"
// 3. "pakcage1/controller"
// 4.

var routes = `package routes

import (
	"%s"
	"%s"
	"%s"

	"github.com/gin-gonic/gin"
)

var (
	%sRepository = repository.New%sRepository(DB)
	%sService = service.New%sService(%sRepository)
	%sController = controller.New%sController(%sService)
)

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

func AddRoutes() string {
	return routes
}
