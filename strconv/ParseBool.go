package main

/*
func ParseBool(str string) (bool, error)
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := "true"
	if s, err := strconv.ParseBool(v); err == nil {
		fmt.Printf("%T, %v\n", s, s)
	}

}

/*
$ go run ParseBool.go
bool, true
*/
