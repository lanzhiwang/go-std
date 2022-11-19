package main

/*
func CanBackquote(s string) bool
*/

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println(strconv.CanBackquote("Fran & Freddie's Diner ☺"))
	fmt.Println(strconv.CanBackquote("`can't backquote this`"))

}

/*
$ go run CanBackquote.go
true
false
*/
