package code

// example:
// 1. User
// 2. entity.User
// 3. entity.User
// 4. entity.UserInput
// 5. entity.UserInput
// 6. user
// 7. User
// 8. user
// 9. user

var Repository = `package repository

type %sRepository interface {
	GetAll() (%s, error)
	GetById(id string) (%s, error)
	Create(input %s) (error)
	UpdateById(id string, input %s) (error)
	DeleteById(id string) (error)
}

type %sRepository struct {

}

func New%sService() *%sRepository {
	return &%sRepository{}
}

`
