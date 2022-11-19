package main

/*
func LastIndexFunc(s string, f func(rune) bool) int
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.LastIndexFunc("go 123", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("123 go", unicode.IsNumber))
	fmt.Println(strings.LastIndexFunc("go", unicode.IsNumber))
}

/*
$ go run LastIndexFunc.go
5
2
-1
*/
