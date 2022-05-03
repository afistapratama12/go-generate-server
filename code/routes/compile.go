package routes

import (
	"fmt"
	"go-gen-server/helper"
)

func compileMainRoutes(packageName string) string {

	return fmt.Sprintf(routesMain, packageName)
}

func compileRoutes(name, packageName, nameUpper, namePlural, nameEntity, nameEntityInput string, isWithMiddleware bool) string {
	var (
		// package
		// package
		// package
		routesStrHead = fmt.Sprintf(routesHead, packageName, packageName, packageName)

		// User
		// User
		// User
		// User
		// User
		// User
		// User
		// User
		routesStrInit = fmt.Sprintf(routesInit, nameUpper, nameUpper, nameUpper, nameUpper,
			nameUpper, nameUpper, nameUpper, nameUpper,
		)

		routesStrRoute = ""
	)

	// User
	// users
	// routeGetAll
	// users
	// routerGetId
	// users
	// routeCreate
	// users
	// routeUpdate
	// users
	// routeDelete
	if isWithMiddleware {
		routesStrRoute = fmt.Sprintf(routesRoute, nameUpper,
			namePlural,
			fmt.Sprintf(routeGetAllMw, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeGetIdMw, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeCreateMw, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeUpdateMw, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeDeleteMw, nameUpper, nameUpper),
		)
	} else {
		routesStrRoute = fmt.Sprintf(routesRoute, nameUpper,
			namePlural,
			fmt.Sprintf(routeGetAll, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeGetId, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeCreate, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeUpdate, nameUpper, nameUpper),
			namePlural,
			fmt.Sprintf(routeDelete, nameUpper, nameUpper),
		)
	}

	return routesStrHead + routesStrInit + routesStrRoute
}

func AddMainRoutes(packageName string) string {
	return compileMainRoutes(packageName)
}

func AddRoutes(name, packageName string, isWithMiddleware bool) string {
	allConvertName, err := helper.ConverterName(name, "entity")
	helper.HandleErrorCmd(err)

	return compileRoutes(name,
		packageName,
		allConvertName.NameUpper,
		allConvertName.NamePlural,
		allConvertName.NameEntity,
		allConvertName.NameEntityInput,
		isWithMiddleware,
	)
}
