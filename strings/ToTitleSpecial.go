package main

/*
func ToTitleSpecial(c unicode.SpecialCase, s string) string
*/

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	fmt.Println(strings.ToTitleSpecial(unicode.TurkishCase, "dünyanın ilk borsa yapısı Aizonai kabul edilir"))
}

/*
$ go run ToTitleSpecial.go
DÜNYANIN İLK BORSA YAPISI AİZONAİ KABUL EDİLİR
*/
