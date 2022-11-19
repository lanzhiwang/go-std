package main

/*
func FieldsFunc(s string, f func(rune) bool) []string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return !unicode.IsLetter(c) && !unicode.IsNumber(c)
	}
	fmt.Printf("Fields are: %q", strings.FieldsFunc("  foo1;bar2,baz3...", f))
}

/*
$ go run FieldsFunc.go
Fields are: ["foo1" "bar2" "baz3"]
*/
