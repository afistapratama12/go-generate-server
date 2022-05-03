package helper

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func HandleErrorCmd(err error) {
	if err != nil {
		fmt.Println("error :", err.Error())
		fmt.Println()
		os.Exit(1)
	}
}

// hello => Hello
func ChangeToUpper(name string) (string, bool) {
	if len(name) < 1 {
		return "", false
	}

	return strings.ToUpper(string(name[0])) + string(name[1:]), true
}

// user => users
// category => categories
func ChangetoPlural(name string) (string, bool) {
	if len(name) < 1 {
		return "", false
	}

	if string(name[len(name)-1]) == "y" {
		return name[:len(name)-1] + "ies", true
	} else {
		return name + "s", true
	}
}

// created_at => CreatedAt
func ChangeSnaketoCamel(name string) string {
	var camelCase string

	split := strings.Split(name, "_")

	for _, s := range split {
		v, _ := ChangeToUpper(s)
		camelCase += v
	}

	return camelCase
}

type PropsStruct struct {
	Name string
	Type string
}

// id:int,created_at:string,updated_at:string
func GetProps(props string) []PropsStruct {
	var propStructs []PropsStruct

	splitProps := strings.Split(props, ",")

	for _, p := range splitProps {

		sNameType := strings.Split(p, ":")

		propStructs = append(propStructs, PropsStruct{
			Name: sNameType[0],
			Type: sNameType[1],
		})
	}

	return propStructs
}

type AllNameConverter struct {
	Name            string
	NameUpper       string
	NamePlural      string
	NameEntity      string
	NameEntityInput string
}

func ConverterName(name string, optionLayer string) (AllNameConverter, error) {
	allNameConverter := AllNameConverter{
		Name: name,
	}

	nameUpper, ok := ChangeToUpper(name)
	if !ok {
		return allNameConverter, errors.New("error happen using helper ChangeToUpper in " + optionLayer)
	}

	allNameConverter.NameUpper = nameUpper

	namePlural, ok := ChangetoPlural(name)
	if !ok {
		return allNameConverter, errors.New("error happen using helper ChangetoPlural in " + optionLayer)
	}
	allNameConverter.NamePlural = namePlural

	nameEntity := "entity." + nameUpper
	nameEntityInput := "entity." + nameUpper + "Input"

	allNameConverter.NameEntity = nameEntity
	allNameConverter.NameEntityInput = nameEntityInput

	return allNameConverter, nil
}
