package helper

import (
	"fmt"
	"os"
	"strings"
)

// user_id:int,role:string => (userId int, role string)
// user_id:int,role:string => map[user_id: true, role: true]
func GenerateToParam(input string) (string, map[string]bool) {
	var output string
	var listNameParam = map[string]bool{}

	splitter := strings.Split(input, ",")
	for i, s := range splitter {
		nameType := strings.Split(s, ":")

		listNameParam[nameType[0]] = true

		if ok := CheckDataTypeInParam(nameType[1], i+1); ok {
			output += ChangeSnaketoCamel(nameType[0]) + " " + nameType[1]

		} else {
			fmt.Println()
			os.Exit(1)
		}

		if i+1 != len(splitter) {
			output += ", "
		}
	}

	return output, listNameParam
}
