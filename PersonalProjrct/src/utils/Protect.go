package utils

import "fmt"

func Protect(err... interface{}) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	} ()
	panic(err)
}
