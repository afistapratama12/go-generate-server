package helper

import (
	"strings"
)

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
