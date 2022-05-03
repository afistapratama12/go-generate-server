package service

// 1. User
// 2. User
// 3. entity.User
// 4. User
// 5. entity.User
// 6. User
// 7. entity.UserInput
// 8. User
// 9. entity.UserInput
// 10. User
var serviceUserInterface = `

type %sService interface {
	GetAll%s() ([]%s, error)
	Get%sById(id string) (%s, error)
	Register(input %s) (error)
	Login(input %s) 
	Update%sById(id string, input %s) (error)
	Delete%sById(id string) (error)
}`
