package main

/*
func ToUpperSpecial(c unicode.SpecialCase, s string) string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToUpperSpecial(unicode.TurkishCase, "örnek iş"))
}

/*
$ go run ToUpperSpecial.go
ÖRNEK İŞ
*/
