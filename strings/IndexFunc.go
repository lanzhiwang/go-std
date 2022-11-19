package main

/*
func IndexFunc(s string, f func(rune) bool) int
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	f := func(c rune) bool {
		return unicode.Is(unicode.Han, c)
	}
	fmt.Println(strings.IndexFunc("Hello, 世界", f))
	fmt.Println(strings.IndexFunc("Hello, world", f))
}

/*
$ go run IndexFunc.go
7
-1
*/
