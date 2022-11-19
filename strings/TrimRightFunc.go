package main

/*
func TrimRightFunc(s string, f func(rune) bool) string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimRightFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}

/*
$ go run TrimRightFunc.go
¡¡¡Hello, Gophers
*/
