package utils

import "fmt"

var utilsListErrorConst = `package utils

const (
	Error4xx               = "error 4xx"
	ErrorUnauthorizeUser   = "error 401 unauthorize user"
	ErrorInternalServer    = "error 500 internal server error"
	ErrorBadRequest        = "error 400 bad request"
	ErrorNotFound 		   = "error 404 parameter not found"
)`

var utilsErrorInit = `

type errMessageModel struct {
	Status  string %s
	Message string %s
	Errors  string %s
}

func ErrorMessages(errMessage string, err error) *errMessageModel {
	return &errMessageModel{
		Status:  "error",
		Message: errMessage,
		Errors:  err.Error(),
	}
}`

var utilsErrorInitStr = fmt.Sprintf(utilsErrorInit,
	"`json:\"status\"`",
	"`json:\"message\"`",
	"`json:\"errors\"`",
)
