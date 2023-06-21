package utils

import "fmt"

func HandleError(err any) {
	if err != nil {
		fmt.Println("An error have ocurred: ", err)
	}
}
