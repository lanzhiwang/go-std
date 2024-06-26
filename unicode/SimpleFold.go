package main

/*
func SimpleFold(r rune) rune
*/

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Printf("%#U\n", unicode.SimpleFold('A'))      // 'a'
	fmt.Printf("%#U\n", unicode.SimpleFold('a'))      // 'A'
	fmt.Printf("%#U\n", unicode.SimpleFold('K'))      // 'k'
	fmt.Printf("%#U\n", unicode.SimpleFold('k'))      // '\u212A' (Kelvin symbol, K)
	fmt.Printf("%#U\n", unicode.SimpleFold('\u212A')) // 'K'
	fmt.Printf("%#U\n", unicode.SimpleFold('1'))      // '1'

}

/*
$ go run SimpleFold.go
U+0061 'a'
U+0041 'A'
U+006B 'k'
U+212A 'K'
U+004B 'K'
U+0031 '1'
*/
