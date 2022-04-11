package code

var service = `package service

type %sService interface {

}

type %sService struct {

}`

func AddService() string {
	return service
}
