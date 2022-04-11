package code

// example:
// 1. user
// 2. User
// 3. user
// 4. user

var controller = `package controller

type %sController struct {

}

func New%sController() *%sController {
	return &%sController{}
}

func (c *%sController) GetAll%s(c *gin.Context) {
	
}

func (c *%sController) Get%sById(c *gin.Context) {
	
}

func (c *%sController) Update%s(c *gin.Context) {
	
}

func (c *%sController) Create%s(c *gin.Context) {
	
}

func (c *%s) Delete%s(c *gin.Context) {
	
}
`

func AddController() string {
	return controller
}
