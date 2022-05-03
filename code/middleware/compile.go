package middleware

import (
	"fmt"
	"go-gen-server/helper"
)

func compileAuthService(params string) string {
	paramFunc, mapListParam := helper.GenerateToParam(params)

	// TODO: change data
	var changeToListParam string

	for key := range mapListParam {
		changeToListParam += "\n\t\t" + `"` + key + `": ` + helper.ChangeSnaketoCamel(key) + ","
	}

	changeToListParam += "\n\t\t"

	var (
		authStrHead     = authHead
		authStrInit     = fmt.Sprintf(authInit, paramFunc)
		authStrGenerate = fmt.Sprintf(authGenerateToken, paramFunc, changeToListParam)
		authStrValidate = authValidateToken
	)

	return authStrHead + authStrInit + authStrGenerate + authStrValidate
}

func compileMiddleware() string {
	return ""
}

func AddMiddleware() string {
	return compileMiddleware()
}

func AddAuthService(param string) string {
	return compileAuthService(param)
}
