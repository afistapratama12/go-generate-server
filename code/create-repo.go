package code

// example:
// 0. dirname => test-server

// 1. User
// 2. entity.User
// 3. entity.User
// 4. entity.UserInput
// 5. entity.UserInput
// ============= sampai last interface

// 6. user
// 7. User
// 8. user
// 9. user
// ============= sampai last new repo

// 10. user
// 11. entity.User
// 12. users
// 13. entity.User
// 14. users
// 15. users
// 16. users
// ============= sampai last GetAll()

// 17.user
// 18. entity.User
// 19 user
// 20 entity.User
// 21. user
// 22. user
// 23. user
// ============= sampai last GetById()

var repository = `package repository

import (
	"%s/entity" 

	"gorm.io/gorm"
)

type %sRepository interface {
	GetAll() (%s, error)
	GetById(id string) (%s, error)
	Create(input %s) (error)
	UpdateById(id string, input %s) (error)
	DeleteById(id string) (error)
}

type %sRepository struct {
	db *gorm.DB
}

func New%sRepository(db *gorm.DB) *%sRepository {
	return &%sRepository{db: db}
}

func (r *%sRepository) GetAll() (%s, error) {
	var %s []%s
	
	if err := r.db.Find(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}

func (r *%sRepository) GetById(id string) (%s, error) {
	var %s %s
	
	if err := r.db.Where("id = ?", id).Find(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}`

func AddRepository() string {
	return repository
}
