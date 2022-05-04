package repository

// example
// 1. dirname => test-server
var repositoryHead = `package repository

import (
	"%s/entity" 

	"gorm.io/gorm"
)`

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
var repositoryInit = `

type %sRepository interface {
	GetAll() ([]%s, error)
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
}`

// example
// 1. user
// 2. entity.User
// 3. users
// 4. entity.User
// 5. users
// 6. users
// 7. users
var repositoryFuncGetAll = `

func (r *%sRepository) GetAll() ([]%s, error) {
	var %s []%s
	
	if err := r.db.Find(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}`

// example
// 1. user
// 2. entity.User
// 3. user
// 4. entity.User
// 5. user
// 6. user
// 7. user
var repositoryFuncGetById = `

func (r *%sRepository) GetById(id string) (%s, error) {
	var %s %s
	
	if err := r.db.Where("id = ?", id).Find(&%s).Error; err != nil {
		return %s, err
	}

	return %s, nil
}`

// example
// 1. user
// 2. entity.UserInput
var repositoryFuncCreate = `

func (r *%sRepository) Create(input %s) (error) {
	if err := r.db.Create(&input).Error ; err != nil {
		return err
	}

	return nil
}`

// example
// 1. user
// 2. entity.UserInput
var repositoryFuncUpdate = `

// TODO: update soon
func (r *%sRepository) UpdateById(id string, input %s) (error) {
	return nil
}`

// example
// user
// entity.User
var repositoryFuncDelete = `

func (r *%sRepository) DeleteById(id string) (error) {
	if err := r.db.Delete(&%s{}, id).Error; err != nil {
		return err
	}

	return nil
}`
