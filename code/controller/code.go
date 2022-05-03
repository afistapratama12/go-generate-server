package controller

// 1. name pacakge = test-service
// 2. name pacakge = test-service
// 3. name pacakge = test-service
var controllerHead = `package controller

import (
	"errors"
	"%s/service"
	"%s/entity"
	"%s/utils"

	"github.com/gin-gonic/gin"
)`

// 1 user
// 2 user
// 3 User

// 4 User
// 5 user
// 6 User
// 7 user
// 8 user
// 9 user
// 10 user
var controllerInit = `

type %sController struct {
	%sService service.%sService
}

func New%sController(%sService service.%sService) *%sController {
	return &%sController{
		%sService: %sService,
	}
}`

// user
// User
// user
// User
var controllerGetAll = `

func (cs *%sController) GetAll%s(c *gin.Context) {
	datas, err := cs.%sService.GetAll%s()
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, datas)
}`

// user
// User
// user
// User
var controllerGetById = `

func (cs *%sController) Get%sById(c *gin.Context) {
	param := c.Param("id")

	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	data, err := cs.%sService.Get%sById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, data)
}`

// user
// User
// entity.UserInput
// user
// User
// user
var controllerCreate = `

func (cs *%sController) Create%s(c *gin.Context) {
	var input %s

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.%sService.Create%s(input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(201, gin.H{
		"message": "success create %s",
	})
}`

// user
// User
// entity.UserInput
// user
// User
// user
var controllerUpdate = `

func (cs *%sController) Update%sById(c *gin.Context) {
	var input %s

	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, err))
		return
	}

	err := cs.%sService.Update%sById(param, input)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success update %s id: " + param,
	})
}`

// user
// User
// user
// User
// user
var controllerDelete = `

func (cs *%sController) Delete%sById(c *gin.Context) {
	param := c.Param("id")
	if param == "" {
		c.JSON(400, utils.ErrorMessages(utils.ErrorBadRequest, errors.New("parameter not valid")))
		return
	}

	err := cs.%sService.Delete%sById(param)
	if err != nil {
		c.JSON(500, utils.ErrorMessages(utils.ErrorInternalServer, err))
		return
	}

	c.JSON(200, gin.H{
		"message": "success delete %s id: " + param,
	})
}`
