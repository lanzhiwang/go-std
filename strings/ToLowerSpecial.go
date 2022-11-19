package main

/*
func ToLowerSpecial(c unicode.SpecialCase, s string) string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToLowerSpecial(unicode.TurkishCase, "Önnek İş"))
}

/*
$ go run ToLowerSpecial.go
önnek iş
*/
