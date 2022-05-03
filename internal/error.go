package internal

import "fmt"

func FailError(err error) {
	if err != nil {
		fmt.Println("error internal : ", err)
		return
	}
}
