package helper

import "fmt"

var checkDataType = map[string]bool{
	"string":      true,
	"int":         true,
	"float32":     true,
	"float64":     true,
	"int32":       true,
	"int64":       true,
	"bool":        true,
	"time.Time":   true,
	"interface{}": true,
}

func CheckDataTypeInParam(param string, paramNo int) bool {
	if _, ok := checkDataType[param]; ok {
		return true
	} else {
		fmt.Println("did not find an existing support generate data type")
		fmt.Println("list : string, int, int32, int64, float32, float64, bool, time.Time, interface{}")

		switch param {
		case "interface":
			fmt.Println("do you mean this data type: interface{}, in param number", paramNo)
		case "str":
			fmt.Println("do you mean this data type: string, in param number", paramNo)
		case "number":
			fmt.Println("do you mean this data type: int | int32 | int64, in param number", paramNo)
		case "numb":
			fmt.Println("do you mean this data type: int | int32 | int64, in param number", paramNo)
		case "integer":
			fmt.Println("do you mean this data type: int | int32 | int64, in param number", paramNo)
		case "float":
			fmt.Println("please specificly float data type: float32 | float64, in param number", paramNo)
		case "date":
			fmt.Println("do you mean this data type: time.Time, in param number", paramNo)
		case "datetime":
			fmt.Println("do you mean this data type: time.Time, in param number", paramNo)
		default:
			TypoDataType(param, paramNo)
		}

		return false
	}
}

func TypoDataType(param string, paramNo int) {
	TypoIntDataType(param, paramNo)
	TypoFloatDataType(param, paramNo)
	TypoStringDataType(param, paramNo)
}

func TypoIntDataType(param string, paramNo int) {
	var stringTypo = []string{"in", "it", "intr", "numbr", "inter"}

	for _, s := range stringTypo {
		if s == param {
			fmt.Println("typo data type, do you mean this data type: int | int32 | int64. In param number", paramNo)
			break
		}
	}
}

func TypoFloatDataType(param string, paramNo int) {
	var stringTypo = []string{"folat32", "folat64", "flaot32", "flaot32", "flat32", "flat64", "flat", "folat", "flaot", "floa32", "floa64", "foat32", "foat64", "foat", "floa", "load32", "load64"}

	for _, s := range stringTypo {
		if s == param {
			fmt.Println("typo data type, do you mean this data type: int | int32 | int64. In param number", paramNo)
			break
		}
	}
}

func TypoStringDataType(param string, paramNo int) {
	var stringTypo = []string{"strung", "strong", "streng", "strang", "srt", "sr", "st", "strng", "sting"}

	for _, s := range stringTypo {
		if s == param {
			fmt.Println("typo data type, do you mean this data type: string. In param number", paramNo)
			break
		}
	}
}
