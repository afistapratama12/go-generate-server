package service

import (
	"fmt"
	"go-gen-server/helper"
)

func compileCodeService(name, packageName, nameUpper, namePlural, nameEntity, nameEntityInput string) string {
	var (
		// 1. package name => test-server
		// 2. package name => test-server
		serviceStrHead = fmt.Sprintf(serviceHead, packageName, packageName)

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
		serviceStrInterface = fmt.Sprintf(serviceInterface, nameUpper, nameUpper, nameEntity, nameUpper, nameEntity,
			nameUpper, nameEntityInput, nameUpper, nameEntityInput, nameUpper,
		)

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
		serviceStrInit = fmt.Sprintf(serviceInit, name, name, nameUpper+"Repository", nameUpper, name,
			nameUpper+"Repository", name, name, name, name,
		)

		// 1. user
		// 2. User
		// 3. entity.User
		// 4. user
		serviceStrGetAll = fmt.Sprintf(serviceGetAll, name, nameUpper, nameEntity, name)

		// o. user
		// 1. User
		// 2. entity.User
		// 3. user
		serviceStrGetById = fmt.Sprintf(serviceGetById, name, nameUpper, nameEntity, name)

		// 1. user
		// 2. User
		// 3. entity.UserInput
		// 4. user
		serviceStrCreate = fmt.Sprintf(serviceCreate, name, nameUpper, nameEntityInput, name)

		// 1. user
		// 2. User
		// 3. entity.UserInput
		// 4. user
		serviceStrUpdate = fmt.Sprintf(serviceUpdate, name, nameUpper, nameEntityInput, name)

		// 1. user
		// 2. User
		// 3. user
		serviceStrDelete = fmt.Sprintf(serviceDelete, name, nameUpper, name)
	)

	return serviceStrHead + serviceStrInterface + serviceStrInit + serviceStrGetAll + serviceStrGetById + serviceStrCreate + serviceStrUpdate + serviceStrDelete
}

func AddService(name string, packageName string) string {
	allConvertName, err := helper.ConverterName(name, "service")
	helper.HandleErrorCmd(err)

	serviceStrCompile := compileCodeService(name, packageName, allConvertName.NameUpper, allConvertName.NamePlural, allConvertName.NameEntity, allConvertName.NameEntityInput)

	return serviceStrCompile
}
