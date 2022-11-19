package main

/*
type CaseRange struct {
	Lo    uint32
	Hi    uint32
	Delta d
}

type SpecialCase []CaseRange

var AzeriCase SpecialCase = _TurkishCase
var TurkishCase SpecialCase = _TurkishCase

func (special SpecialCase) ToLower(r rune) rune
func (special SpecialCase) ToTitle(r rune) rune
func (special SpecialCase) ToUpper(r rune) rune

*/

import (
	"fmt"
	"unicode"
)

func main() {
	t := unicode.TurkishCase

	const lci = 'i'
	fmt.Printf("%#U\n", t.ToLower(lci))
	fmt.Printf("%#U\n", t.ToTitle(lci))
	fmt.Printf("%#U\n", t.ToUpper(lci))

	const uci = 'İ'
	fmt.Printf("%#U\n", t.ToLower(uci))
	fmt.Printf("%#U\n", t.ToTitle(uci))
	fmt.Printf("%#U\n", t.ToUpper(uci))

}

/*
$ go run SpecialCase.go
U+0069 'i'
U+0130 'İ'
U+0130 'İ'
U+0069 'i'
U+0130 'İ'
U+0130 'İ'
*/
