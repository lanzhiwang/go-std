package main

/*
func TrimFunc(s string, f func(rune) bool) string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Print(strings.TrimFunc("¡¡¡Hello, Gophers!!!", func(r rune) bool {
		return !unicode.IsLetter(r) && !unicode.IsNumber(r)
	}))
}

/*
$ go run TrimFunc.go
Hello, Gophers
*/
