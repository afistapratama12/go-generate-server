package controller

import (
	"fmt"
	"go-gen-server/helper"
)

func compileCodeController(name, packageName, nameUpper, namePlural, nameEntity, nameEntityInput string) string {
	var (
		// 1 pkg name
		// 2 pkg name
		// 3 pkg name
		controllerStrHead = fmt.Sprintf(controllerHead, packageName, packageName, packageName)

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
		controllerStrInit = fmt.Sprintf(controllerInit, name, name, nameUpper, nameUpper, name,
			nameUpper, name, name, name, name,
		)

		// user
		// User
		// user
		// User
		controllerStrGetAll = fmt.Sprintf(controllerGetAll, name, nameUpper, name, nameUpper)

		// user
		// User
		// user
		// User
		controllerStrGetById = fmt.Sprintf(controllerGetById, name, nameUpper, name, nameUpper)

		// user
		// User
		// entity.UserInput
		// user
		// User
		// user
		controllerStrCreate = fmt.Sprintf(controllerCreate, name, nameUpper, nameEntityInput, name, nameUpper, name)

		// user
		// User
		// entity.UserInput
		// user
		// User
		// user
		controllerStrUpdate = fmt.Sprintf(controllerUpdate, name, nameUpper, nameEntityInput, name, nameUpper, name)

		// user
		// User
		// user
		// User
		// user
		controllerStrDelete = fmt.Sprintf(controllerDelete, name, nameUpper, name, nameUpper, name)
	)

	return controllerStrHead + controllerStrInit + controllerStrGetAll + controllerStrGetById + controllerStrCreate + controllerStrUpdate + controllerStrDelete
}

func AddController(name string, packageName string) string {
	allConvertName, err := helper.ConverterName(name, "controller")
	helper.HandleErrorCmd(err)

	controllerStrCompile := compileCodeController(name,
		packageName,
		allConvertName.NameUpper,
		allConvertName.NamePlural,
		allConvertName.NameEntity,
		allConvertName.NameEntityInput,
	)

	return controllerStrCompile
}
