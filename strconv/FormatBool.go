package main

/*
func FormatBool(b bool) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	v := true
	s := strconv.FormatBool(v)
	fmt.Printf("%T, %v\n", s, s)

}

/*
$ go run FormatBool.go
string, true
*/
