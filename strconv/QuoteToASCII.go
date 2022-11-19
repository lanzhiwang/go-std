package main

/*
func QuoteToASCII(s string) string
*/

import (
	"fmt"
	"strconv"
)

func main() {
	// This string literal contains a tab character.
	s := strconv.QuoteToASCII(`"Fran & Freddie's Diner	☺"`)
	fmt.Println(s)

}

/*
$ go run QuoteToASCII.go
"\"Fran & Freddie's Diner\t\u263a\""
*/
