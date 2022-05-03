package service

// 1. package name => test-server
// 2. package name => test-server
var serviceHead = `package service

import (
	"%s/repository"
	"%s/entity"
)`

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
var serviceInterface = `

type %sService interface {
	GetAll%s() ([]%s, error)
	Get%sById(id string) (%s, error)
	Create%s(input %s) (error)
	Update%sById(id string, input %s) (error)
	Delete%sById(id string) (error)
}`

// 1. user
// 2. user
// 3. UserRepository / interfaceName
// 4. User
// 5. user
// 6. UserRepository
// 7. user
// 8. user
// 9. user
// 10. user
var serviceInit = `

type %sService struct {
	%sRepository repository.%s
}

func New%sService(%sRepository repository.%s) *%sService {
	return &%sService{
		%sRepository: %sRepository,
	}
}`

// 1. user
// 2. User
// 3. entity.User
// 4. user
var serviceGetAll = `

func (s *%sService) GetAll%s() ([]%s, error) {
	return s.%sRepository.GetAll()
}`

// 0. user
// 1. User
// 2. entity.User
// 3. user
var serviceGetById = `

func (s *%sService) Get%sById(id string) (%s, error) {
	return s.%sRepository.GetById(id)
}`

// 1. user
// 2. User
// 3. entity.UserInput
// 4. user
var serviceCreate = `

func (s *%sService) Create%s(input %s) (error) {
	return s.%sRepository.Create(input)
}`

// 1. user
// 2. User
// 3. entity.UserInput
// 4. user
var serviceUpdate = `

func (s *%sService) Update%sById(id string, input %s) (error) {
	return s.%sRepository.UpdateById(id, input)
}`

// 1. user
// 2. User
// 3. user
var serviceDelete = `

func (s *%sService) Delete%sById(id string) (error) {
	return s.%sRepository.DeleteById(id)
}`
