package main

/*
func TrimLeftFunc(s string, f func(rune) bool) string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimLeftFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}

/*
$ go run TrimLeftFunc.go
Hello, Gophers!!!
*/
