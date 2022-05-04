package repository

import (
	"fmt"
	"go-gen-server/helper"
)

func compileRepoCodeString(name, packageName, nameUpper, namePlural, nameEntity, nameEntityInput string) string {
	// compile code string
	var (
		repoStrHead = fmt.Sprintf(repositoryHead, packageName)
		repoStrInit = fmt.Sprintf(repositoryInit,
			nameUpper,
			nameEntity,
			nameEntity,
			nameEntityInput,
			nameEntityInput,
			name,
			nameUpper,
			name,
			name,
		)

		repoStrGetAll = fmt.Sprintf(repositoryFuncGetAll,
			name,
			nameEntity,
			namePlural,
			nameEntity,
			namePlural,
			namePlural,
			namePlural,
		)

		repoStrGetById = fmt.Sprintf(repositoryFuncGetById,
			name,
			nameEntity,
			name,
			nameEntity,
			name,
			name,
			name,
		)

		repoStrCreate = fmt.Sprintf(repositoryFuncCreate, name, nameEntityInput)

		repoStrUpdate = fmt.Sprintf(repositoryFuncUpdate, name, nameEntityInput)
		repoStrDelete = fmt.Sprintf(repositoryFuncDelete, name, nameEntity)
	)

	return repoStrHead + repoStrInit + repoStrGetAll + repoStrGetById + repoStrCreate + repoStrUpdate + repoStrDelete
}

func AddRepository(name string, packageName string) string {
	allConvertName, err := helper.ConverterName(name, "entity")
	helper.HandleErrorCmd(err)

	codeRepoStr := compileRepoCodeString(name,
		packageName,
		allConvertName.NameUpper,
		allConvertName.NamePlural,
		allConvertName.NameEntity,
		allConvertName.NameEntityInput,
	)

	return codeRepoStr
}
