package code

import (
	"fmt"
	"go-gen-server/helper"
)

// 2 entity
// <Upper-name model> => ex : User
// <Upper-name model>Input => ex : UserInput

var modelUp = `package entity

type %s struct {`

var modelDown = "\n}"

func CreateModelStruct(name string, props []helper.PropsStruct) string {
	upperName, _ := helper.ChangeToUpper(name)

	// create Model
	var dataModel = fmt.Sprintf(modelUp, upperName)

	for _, p := range props {
		camelP := helper.ChangeSnaketoCamel(p.Name)

		dataModel += fmt.Sprintf("\n\t%s %s", camelP, p.Type)
	}

	dataModel += modelDown

	// create ModelInput
	dataModel += fmt.Sprintf("\ntype %sInput struct {", upperName)
	for _, p := range props {
		if p.Name == "id" || p.Name == "created_at" || p.Name == "updated_at" {
			continue
		} else {
			camelP := helper.ChangeSnaketoCamel(p.Name)

			dataModel += fmt.Sprintf("\n\t%s %s", camelP, p.Type)
		}
	}

	dataModel += modelDown

	return dataModel
}

func AddCreateModel(name string, props string, isDefault bool) string {
	var initProps string

	if isDefault {
		if props == "" {
			initProps += "id:int,created_at:string,updated_at:string"
		} else {
			initProps += "id:int,created_at:string,updated_at:string," + props
		}
	} else {
		initProps += props
	}

	propStructs := helper.GetProps(initProps)

	dataModel := CreateModelStruct(name, propStructs)

	return dataModel
}
